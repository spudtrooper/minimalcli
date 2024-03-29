package handler

import (
	"context"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/sets"
)

type ctorFn func() any
type handlerFn func(ctx context.Context, ip any) (any, error)

func NewHandler(name string, hf handlerFn, paramsPrototype any, optss ...NewHandlerOption) Handler {
	var pCtor ctorFn = func() any {
		res := paramsPrototype
		return res
	}
	opts := MakeNewHandlerOptions(optss...)
	fields := exportedFields(paramsPrototype)
	metadata := metadataFromFields(fields, opts.ExtraRequiredFields())
	fn := fnFromStructAndParams(hf, pCtor, fields)
	cliOnly := opts.CliOnly()
	method := or.String(opts.Method(), "GET")
	return &handler{
		name:     name,
		fn:       fn,
		cliOnly:  cliOnly,
		metadata: metadata,
		method:   method,
		renderer: opts.Renderer(),
	}
}

func NewStaticHandler(name string, htmlBytes []byte, paramsPrototype any, optss ...NewHandlerOption) Handler {
	opts := MakeNewHandlerOptions(optss...)
	config := opts.RendererConfig()
	renderer := func(any) ([]byte, RendererConfig, error) {
		return htmlBytes, config, nil
	}
	method := or.String(opts.Method(), "GET")
	res := &handler{
		name:     name,
		webOnly:  true,
		isStatic: true,
		method:   method,
		renderer: renderer,
	}
	if paramsPrototype != nil {
		fields := exportedFields(paramsPrototype)
		res.metadata = metadataFromFields(fields, opts.ExtraRequiredFields())
	}
	return res
}

func metadataFromFields(fs []reflect.StructField, requiredFields []string) HandlerMetadata {
	return HandlerMetadata{
		Params: paramsFromStruct(fs, requiredFields),
	}
}

func typeFromKind(f reflect.StructField) HandlerMetadataParamType {
	switch k := f.Type.Kind(); k {
	case reflect.String:
		return HandlerMetadataParamTypeString
	case reflect.Int:
		return HandlerMetadataParamTypeInt
	case reflect.Int64:
		// Assume this is a time.Duration.
		return HandlerMetadataParamTypeDuration
	case reflect.Bool:
		return HandlerMetadataParamTypeBool
	case reflect.Float32:
		return HandlerMetadataParamTypeFloat32
	case reflect.Float64:
		return HandlerMetadataParamTypeFloat64
	case reflect.Struct:
		if f.Type.String() == "time.Time" {
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

func paramsFromStruct(fs []reflect.StructField, requiredFields []string) []HandlerMetadataParam {
	var params []HandlerMetadataParam
	reqFields := sets.String(requiredFields)
	for _, f := range fs {
		name, required, def := findFieldMetadata(f)
		if _, ok := reqFields[name]; ok {
			required = true
		}
		params = append(params, HandlerMetadataParam{
			Name:     name,
			Type:     typeFromKind(f),
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
			var s string
			if required {
				v, ok := ctx.MustString(nameInCtx)
				if !ok {
					shouldHandle = false
					break
				}
				s = v
			} else {
				s = ctx.String(nameInCtx)
			}
			if *debug {
				log.Printf("setting string %s to %q", sf.Name, s)
			}
			f.SetString(s)
		case reflect.Int:
			var i int
			if required {
				v, ok := ctx.MustInt(nameInCtx)
				if !ok {
					shouldHandle = false
					break
				}
				i = v
			} else {
				i = ctx.Int(nameInCtx)
			}
			if *debug {
				log.Printf("setting int %s to %d", sf.Name, i)
			}
			f.SetInt(int64(i))
		case reflect.Int64:
			// First try this as a time.Duration.
			if v := ctx.Duration(nameInCtx); v != 0 {
				if *debug {
					log.Printf("setting duration %s to %v", sf.Name, v)
				}
				// TODO: Allow required time.Duration
				f.SetInt(int64(v))
				break
			}
			// Fall back to int64.
			var i int
			if required {
				v, ok := ctx.MustInt(nameInCtx)
				if !ok {
					shouldHandle = false
					break
				}
				i = v
			} else {
				i = ctx.Int(nameInCtx)
			}
			if *debug {
				log.Printf("setting int64 %s to %d", sf.Name, i)
			}
			f.SetInt(int64(i))
		case reflect.Bool:
			v := ctx.Bool(nameInCtx)
			if *debug {
				log.Printf("setting int64 %s to %t", sf.Name, v)
			}
			f.SetBool(v)
		case reflect.Float32:
			v := ctx.Float32(nameInCtx)
			if *debug {
				log.Printf("setting float32 %s to %f", sf.Name, v)
			}
			f.SetFloat(float64(v))
		case reflect.Float64:
			v := ctx.Float64(nameInCtx)
			if *debug {
				log.Printf("setting float64 %s to %f", sf.Name, v)
			}
			f.SetFloat(v)
		case reflect.Struct:
			if f.Type().String() == "time.Time" {
				t, err := ctx.Time(nameInCtx)
				if err != nil {
					return nil, false, errors.Errorf("evaluating time for %q: %v", nameInCtx, err)
				}
				v := reflect.ValueOf(t)
				if *debug {
					log.Printf("setting time %s to %v", sf.Name, v)
				}
				f.Set(v)
			}
		default:
			return nil, false, errors.Errorf(
				"unkown type for nameInCtx:%s f:%+v kind:%s", nameInCtx, f, f.Kind())
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
