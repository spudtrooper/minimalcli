package handler

import (
	"context"
	"log"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
)

type ctorFn func() any
type handlerFn func(ctx context.Context, ip any) (any, error)

func NewHandler(name string, hf handlerFn, p any, optss ...NewHandlerOption) Handler {
	var pCtor ctorFn = func() any {
		res := p
		return res
	}
	opts := MakeNewHandlerOptions(optss...)
	fields := exportedFields(p)
	metadata := metadataFromStruct(fields)
	fn := fnFromStructAndParams(hf, pCtor, fields)
	cliOnly := opts.CliOnly()
	method := or.String(opts.Method(), "GET")
	return &handler{
		name:     name,
		fn:       fn,
		cliOnly:  cliOnly,
		metadata: metadata,
		method:   method,
	}
}

func metadataFromStruct(fs []reflect.StructField) HandlerMetadata {
	return HandlerMetadata{
		Params: paramsFromStruct(fs),
	}
}

func typeFromKind(k reflect.Kind) HandlerMetadataParamType {
	switch k {
	case reflect.String:
		return HandlerMetadataParamTypeString
	case reflect.Int:
		return HandlerMetadataParamTypeInt
	case reflect.Bool:
		return HandlerMetadataParamTypeBool
	case reflect.Float32:
		return HandlerMetadataParamTypeFloat32
	case reflect.Float64:
		return HandlerMetadataParamTypeFloat64
	case reflect.Struct:
		if k.String() == "time.Duration" {
			return HandlerMetadataParamTypeDuration
		}
		if k.String() == "time.Time" {
			return HandlerMetadataParamTypeTime
		}
	default:
		return HandlerMetadataParamTypeUnknown
	}
	return HandlerMetadataParamTypeUnknown
}

func exportedFields(o any) []reflect.StructField {
	var fs []reflect.StructField
	typ := reflect.TypeOf(o)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if f.IsExported() {
			fs = append(fs, f)
		}
	}
	return fs
}

func paramsFromStruct(fs []reflect.StructField) []HandlerMetadataParam {
	var params []HandlerMetadataParam
	for _, f := range fs {
		name, required, def := findFieldMetadata(f)
		params = append(params, HandlerMetadataParam{
			Name:     name,
			Type:     typeFromKind(f.Type.Kind()),
			Required: required,
			Default:  def,
		})
	}
	return params
}

func findFieldMetadata(f reflect.StructField) (name string, required bool, def string) {
	if t, ok := f.Tag.Lookup("required"); ok {
		val := strings.Split(t, ",")[0]
		required = strings.EqualFold(val, "true") || strings.EqualFold(val, "1")
	}
	if t, ok := f.Tag.Lookup("json"); ok {
		name = strings.Split(t, ",")[0]
	} else {
		name = strcase.ToSnake(f.Name)
	}
	if t, ok := f.Tag.Lookup("default"); ok {
		def = strings.Split(t, ",")[0]
	}
	return
}

func setValuesOnParams(ctx EvalContext, pCtor ctorFn, fs []reflect.StructField) (interface{}, bool, error) {
	params := pCtor()

	// https://stackoverflow.com/questions/63421976/panic-reflect-call-of-reflect-value-fieldbyname-on-interface-value
	v := reflect.ValueOf(&params).Elem()
	tmp := reflect.New(v.Elem().Type()).Elem()
	tmp.Set(v.Elem())

	shouldHandle := true
	for _, sf := range fs {
		f := tmp.FieldByName(sf.Name)
		if !f.CanSet() {
			log.Printf("ERROR: !CanSet: %+v", f)
			continue
		}
		if !f.IsValid() {
			log.Printf("ERROR: !IsValid: %+v", f)
			continue
		}
		nameInCtx, required, _ := findFieldMetadata(sf)
		switch f.Kind() {
		case reflect.String:
			if required {
				v, ok := ctx.MustString(nameInCtx)
				if !ok {
					shouldHandle = false
					break
				}
				f.SetString(v)
			} else {
				f.SetString(ctx.String(nameInCtx))
			}
		case reflect.Int:
			if required {
				v, ok := ctx.MustInt(nameInCtx)
				if !ok {
					shouldHandle = false
					break
				}
				f.SetInt(int64(v))
			} else {
				f.SetInt(int64(ctx.Int(nameInCtx)))
			}
		case reflect.Bool:
			f.SetBool(ctx.Bool(nameInCtx))
		case reflect.Float32:
			f.SetFloat(float64(ctx.Float32(nameInCtx)))
		case reflect.Float64:
			f.SetFloat(ctx.Float64(nameInCtx))
		case reflect.Struct:
			if f.Type().String() == "time.Duration" {
				f.Set(reflect.ValueOf(ctx.Duration(nameInCtx)))
			}
			if f.Type().String() == "time.Time" {
				t, err := ctx.Time(nameInCtx)
				if err != nil {
					return nil, false, errors.Errorf("evaluating time for %q: %v", nameInCtx, err)
				}
				f.Set(reflect.ValueOf(t))
			}
		default:
			return nil, false, errors.Errorf("unkown type")
		}
	}
	v.Set(tmp)

	return params, shouldHandle, nil
}

func fnFromStructAndParams(hf handlerFn, pCtor ctorFn, fs []reflect.StructField) HandlerFn {
	return func(ctx EvalContext) (interface{}, error) {
		params, shouldHandle, err := setValuesOnParams(ctx, pCtor, fs)
		if err != nil {
			return nil, err
		}

		if !shouldHandle {
			return nil, nil
		}

		res, err := hf(ctx.Context(), params)

		return res, err
	}
}
