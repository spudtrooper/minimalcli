// Package handler generates a CLI and frontend API server to access a third party API.
package handler

import (
	"github.com/spudtrooper/goutil/or"
)

type HandlerFn func(ctx EvalContext) (interface{}, error)

type Handler interface{}

type handler struct {
	name     string
	fn       HandlerFn
	cliOnly  bool
	webOnly  bool
	isStatic bool
	method   string
	metadata HandlerMetadata
	renderer Renderer
}

type HandlerMetadataParamType string

const (
	HandlerMetadataParamTypeString   HandlerMetadataParamType = "string"
	HandlerMetadataParamTypeInt      HandlerMetadataParamType = "int"
	HandlerMetadataParamTypeBool     HandlerMetadataParamType = "bool"
	HandlerMetadataParamTypeDuration HandlerMetadataParamType = "duration"
	HandlerMetadataParamTypeTime     HandlerMetadataParamType = "time"
	HandlerMetadataParamTypeFloat32  HandlerMetadataParamType = "float32"
	HandlerMetadataParamTypeFloat64  HandlerMetadataParamType = "float64"
	HandlerMetadataParamTypeUnknown  HandlerMetadataParamType = "unknown"
)

type HandlerMetadataParam struct {
	Name     string
	Type     HandlerMetadataParamType
	Required bool
	Default  string
}

type HandlerMetadata struct {
	Params []HandlerMetadataParam
}

func (m HandlerMetadata) Empty() bool { return len(m.Params) == 0 }

type RendererConfig struct {
	// Whether the payload returned does not contain full HTML and should
	// be inserted into HTML before being shown.
	IsFragment bool
}

type Renderer func(any) ([]byte, RendererConfig, error)

//go:generate genopts --function NewHandler cliOnly metadata:HandlerMetadata method:string extraRequiredFields:[]string renderer:Renderer webOnly rendererConfig:RendererConfig

func NewHandlerFromHandlerFn(name string, fn HandlerFn, optss ...NewHandlerOption) Handler {
	opts := MakeNewHandlerOptions(optss...)
	return &handler{
		name:     name,
		fn:       fn,
		cliOnly:  opts.CliOnly(),
		webOnly:  opts.WebOnly(),
		metadata: opts.Metadata(),
		method:   or.String(opts.Method(), "GET"),
		renderer: opts.Renderer(),
	}
}
