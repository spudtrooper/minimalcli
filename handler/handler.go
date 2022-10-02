// Package handler generates a CLI and frontend API server to access a third party API.
package handler

import "github.com/spudtrooper/goutil/or"

type HandlerFn func(ctx EvalContext) (interface{}, error)

type Handler interface{}

type handler struct {
	name     string
	fn       HandlerFn
	cliOnly  bool
	method   string
	metadata HandlerMetadata
}

type HandlerMetadataParamType string

const (
	HandlerMetadataParamTypeString   HandlerMetadataParamType = "string"
	HandlerMetadataParamTypeInt      HandlerMetadataParamType = "int"
	HandlerMetadataParamTypeBool     HandlerMetadataParamType = "bool"
	HandlerMetadataParamTypeDuration HandlerMetadataParamType = "duration"
	HandlerMetadataParamTypeFloat32  HandlerMetadataParamType = "float32"
)

type HandlerMetadataParam struct {
	Name     string
	Type     HandlerMetadataParamType
	Required bool
}

type HandlerMetadata struct {
	Params []HandlerMetadataParam
}

func (m HandlerMetadata) Empty() bool { return len(m.Params) == 0 }

//go:generate genopts --function NewHandler cliOnly metadata:HandlerMetadata method:string
func NewHandler(name string, fn HandlerFn, optss ...NewHandlerOption) Handler {
	opts := MakeNewHandlerOptions(optss...)
	return &handler{
		name:     name,
		fn:       fn,
		cliOnly:  opts.CliOnly(),
		metadata: opts.Metadata(),
		method:   or.String(opts.Method(), "GET"),
	}
}
