package handler

import (
	"log"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
)

type CtorFn func() any
type ReflectHandlerFn func(ip any) (any, error)

func NewHandlerFromStruct(name string, ctor CtorFn, optss ...NewHandlerOption) Handler {
	opts := MakeNewHandlerOptions(optss...)
	fields := exportedFields(ctor)
	metadata := metadataFromStruct(fields)
	fn := fnFromStruct(ctor, fields)
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

func NewHandlerFromParams(name string, hf ReflectHandlerFn, pCtor CtorFn, optss ...NewHandlerOption) Handler {
	opts := MakeNewHandlerOptions(optss...)
	fields := exportedFields(pCtor)
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

func exportedFields(ctor CtorFn) []reflect.StructField {
	var fs []reflect.StructField
	o := ctor()
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
		name, required := findFieldMetadata(f)
		params = append(params, HandlerMetadataParam{
			Name:     name,
			Type:     typeFromKind(f.Type.Kind()),
			Required: required,
		})
	}
	return params
}

func findFieldMetadata(f reflect.StructField) (name string, required bool) {
	if t, ok := f.Tag.Lookup("required"); ok {
		val := strings.Split(t, ",")[0]
		required = strings.EqualFold(val, "true") || strings.EqualFold(val, "1")
	}
	if t, ok := f.Tag.Lookup("json"); ok {
		name = strings.Split(t, ",")[0]
	} else {
		name = strcase.ToSnake(f.Name)
	}
	return
}

func setValuesOnParams(ctx EvalContext, pCtor CtorFn, fs []reflect.StructField) (interface{}, bool, error) {
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
		nameInCtx, required := findFieldMetadata(sf)
		switch f.Kind() {
		case reflect.String:
			if required {
				s, ok := ctx.MustString(nameInCtx)
				if !ok {
					shouldHandle = false
					break
				}
				f.SetString(s)
			} else {
				f.SetString(ctx.String(nameInCtx))
			}
		case reflect.Int:
			f.SetInt(int64(ctx.Int(nameInCtx)))
		case reflect.Bool:
			f.SetBool(ctx.Bool(nameInCtx))
		case reflect.Float32:
			f.SetFloat(float64(ctx.Float32(nameInCtx)))
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

func fnFromStructAndParams(hf ReflectHandlerFn, pCtor CtorFn, fs []reflect.StructField) HandlerFn {
	return func(ctx EvalContext) (interface{}, error) {
		params, shouldHandle, err := setValuesOnParams(ctx, pCtor, fs)
		if err != nil {
			return nil, err
		}

		if !shouldHandle {
			return nil, nil
		}

		res, err := hf(params)

		return res, err
	}
}

func fnFromStruct(ctor CtorFn, fs []reflect.StructField) HandlerFn {
	return func(ctx EvalContext) (interface{}, error) {
		o, shouldHandle, err := setValuesOnParams(ctx, ctor, fs)
		if err != nil {
			return nil, err
		}

		if !shouldHandle {
			return nil, nil
		}

		// Looking for Handle() (interface{}, error)

		handle := reflect.ValueOf(&o).Elem().Elem().MethodByName("Handle")
		if !handle.IsValid() {
			return nil, errors.Errorf("Handle method isn't valid")
		}

		vals := handle.Call([]reflect.Value{})
		if len(vals) != 2 {
			return nil, errors.Errorf(
				"expecting 2 return values from Handle and got %d", len(vals))
		}
		var res interface{}
		if !vals[0].IsNil() {
			res = vals[0].Interface()
		}
		if !vals[1].IsNil() {
			err = vals[1].Interface().(error)
		}

		return res, err
	}
}
