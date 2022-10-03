package handler

import (
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spudtrooper/goutil/or"
)

type CtorFn func() interface{}

func NewHandlerFromStruct(name string, ctor CtorFn, optss ...NewHandlerOption) Handler {
	opts := MakeNewHandlerOptions(optss...)
	fn := fnFromStruct(ctor)
	metadata := metadataFromStruct(ctor)
	method := or.String(opts.Method(), "GET")
	return &handler{
		name:     name,
		fn:       fn,
		cliOnly:  opts.CliOnly(),
		metadata: metadata,
		method:   method,
	}
}

func metadataFromStruct(ctor CtorFn) HandlerMetadata {
	return HandlerMetadata{
		Params: paramsFromStruct(ctor),
	}
}

func toSnakeCase(s string) string {
	return strcase.ToSnake(s)
	// return strcase.ToLowerCamel(s)
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
	default:
		return HandlerMetadataParamTypeUnknown
	}
	return HandlerMetadataParamTypeUnknown
}

func paramsFromStruct(ctor CtorFn) []HandlerMetadataParam {
	var params []HandlerMetadataParam
	o := ctor()
	typ := reflect.TypeOf(o)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if !f.IsExported() {
			continue
		}
		params = append(params, HandlerMetadataParam{
			Name: toSnakeCase(f.Name),
			Type: typeFromKind(f.Type.Kind()),
		})
	}
	return params

}

func findJsonName(f reflect.StructField) string {
	if jt, ok := f.Tag.Lookup("json"); ok {
		return strings.Split(jt, ",")[0]
	}
	return toSnakeCase(f.Name)
}

func fnFromStruct(ctor CtorFn) HandlerFn {
	return func(ctx EvalContext) (interface{}, error) {
		o := ctor()
		typ := reflect.TypeOf(o)

		// https://stackoverflow.com/questions/63421976/panic-reflect-call-of-reflect-value-fieldbyname-on-interface-value
		v := reflect.ValueOf(&o).Elem()
		tmp := reflect.New(v.Elem().Type()).Elem()
		tmp.Set(v.Elem())

		for i := 0; i < tmp.NumField(); i++ {
			if f := typ.Field(i); !f.IsExported() {
				continue
			}
			f := tmp.FieldByName(typ.Field(i).Name)
			if !f.CanSet() || !f.IsValid() {
				continue
			}
			nameInCtx := findJsonName(typ.Field(i))
			switch f.Kind() {
			case reflect.String:
				f.SetString(ctx.String(nameInCtx))
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
			default:
				panic("unkown type")
			}
		}
		v.Set(tmp)

		handle := reflect.ValueOf(&o).Elem().Elem().MethodByName("Handle")
		if !handle.IsValid() {
			panic("invalid Handle method")
		}
		vals := handle.Call([]reflect.Value{})
		if len(vals) != 2 {
			panic("expecting 2 values")
		}
		var res interface{}
		var err error
		if !vals[0].IsNil() {
			res = vals[0].Interface()
		}
		if !vals[1].IsNil() {
			err = vals[1].Interface().(error)
		}
		return res, err
	}
}
