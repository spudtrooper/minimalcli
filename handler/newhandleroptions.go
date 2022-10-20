// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

import "fmt"

type NewHandlerOption struct {
	f func(*newHandlerOptionImpl)
	s string
}

func (o NewHandlerOption) String() string { return o.s }

type NewHandlerOptions interface {
	CliOnly() bool
	HasCliOnly() bool
	ExtraRequiredFields() []string
	HasExtraRequiredFields() bool
	Metadata() HandlerMetadata
	HasMetadata() bool
	Method() string
	HasMethod() bool
	Renderer() Renderer
	HasRenderer() bool
	RendererConfig() RendererConfig
	HasRendererConfig() bool
	WebOnly() bool
	HasWebOnly() bool
}

func NewHandlerCliOnly(cliOnly bool) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		opts.has_cliOnly = true
		opts.cliOnly = cliOnly
	}, fmt.Sprintf("handler.NewHandlerCliOnly(bool %+v)}", cliOnly)}
}
func NewHandlerCliOnlyFlag(cliOnly *bool) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		if cliOnly == nil {
			return
		}
		opts.has_cliOnly = true
		opts.cliOnly = *cliOnly
	}, fmt.Sprintf("handler.NewHandlerCliOnly(bool %+v)}", cliOnly)}
}

func NewHandlerExtraRequiredFields(extraRequiredFields []string) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		opts.has_extraRequiredFields = true
		opts.extraRequiredFields = extraRequiredFields
	}, fmt.Sprintf("handler.NewHandlerExtraRequiredFields([]string %+v)}", extraRequiredFields)}
}
func NewHandlerExtraRequiredFieldsFlag(extraRequiredFields *[]string) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		if extraRequiredFields == nil {
			return
		}
		opts.has_extraRequiredFields = true
		opts.extraRequiredFields = *extraRequiredFields
	}, fmt.Sprintf("handler.NewHandlerExtraRequiredFields([]string %+v)}", extraRequiredFields)}
}

func NewHandlerMetadata(metadata HandlerMetadata) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		opts.has_metadata = true
		opts.metadata = metadata
	}, fmt.Sprintf("handler.NewHandlerMetadata(HandlerMetadata %+v)}", metadata)}
}
func NewHandlerMetadataFlag(metadata *HandlerMetadata) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		if metadata == nil {
			return
		}
		opts.has_metadata = true
		opts.metadata = *metadata
	}, fmt.Sprintf("handler.NewHandlerMetadata(HandlerMetadata %+v)}", metadata)}
}

func NewHandlerMethod(method string) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		opts.has_method = true
		opts.method = method
	}, fmt.Sprintf("handler.NewHandlerMethod(string %+v)}", method)}
}
func NewHandlerMethodFlag(method *string) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		if method == nil {
			return
		}
		opts.has_method = true
		opts.method = *method
	}, fmt.Sprintf("handler.NewHandlerMethod(string %+v)}", method)}
}

func NewHandlerRenderer(renderer Renderer) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		opts.has_renderer = true
		opts.renderer = renderer
	}, fmt.Sprintf("handler.NewHandlerRenderer(Renderer %+v)}", renderer)}
}
func NewHandlerRendererFlag(renderer *Renderer) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		if renderer == nil {
			return
		}
		opts.has_renderer = true
		opts.renderer = *renderer
	}, fmt.Sprintf("handler.NewHandlerRenderer(Renderer %+v)}", renderer)}
}

func NewHandlerRendererConfig(rendererConfig RendererConfig) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		opts.has_rendererConfig = true
		opts.rendererConfig = rendererConfig
	}, fmt.Sprintf("handler.NewHandlerRendererConfig(RendererConfig %+v)}", rendererConfig)}
}
func NewHandlerRendererConfigFlag(rendererConfig *RendererConfig) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		if rendererConfig == nil {
			return
		}
		opts.has_rendererConfig = true
		opts.rendererConfig = *rendererConfig
	}, fmt.Sprintf("handler.NewHandlerRendererConfig(RendererConfig %+v)}", rendererConfig)}
}

func NewHandlerWebOnly(webOnly bool) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		opts.has_webOnly = true
		opts.webOnly = webOnly
	}, fmt.Sprintf("handler.NewHandlerWebOnly(bool %+v)}", webOnly)}
}
func NewHandlerWebOnlyFlag(webOnly *bool) NewHandlerOption {
	return NewHandlerOption{func(opts *newHandlerOptionImpl) {
		if webOnly == nil {
			return
		}
		opts.has_webOnly = true
		opts.webOnly = *webOnly
	}, fmt.Sprintf("handler.NewHandlerWebOnly(bool %+v)}", webOnly)}
}

type newHandlerOptionImpl struct {
	cliOnly                 bool
	has_cliOnly             bool
	metadata                HandlerMetadata
	has_metadata            bool
	method                  string
	has_method              bool
	extraRequiredFields     []string
	has_extraRequiredFields bool
	renderer                Renderer
	has_renderer            bool
	webOnly                 bool
	has_webOnly             bool
	rendererConfig          RendererConfig
	has_rendererConfig      bool
}

func (n *newHandlerOptionImpl) CliOnly() bool                  { return n.cliOnly }
func (n *newHandlerOptionImpl) HasCliOnly() bool               { return n.has_cliOnly }
func (n *newHandlerOptionImpl) ExtraRequiredFields() []string  { return n.extraRequiredFields }
func (n *newHandlerOptionImpl) HasExtraRequiredFields() bool   { return n.has_extraRequiredFields }
func (n *newHandlerOptionImpl) Metadata() HandlerMetadata      { return n.metadata }
func (n *newHandlerOptionImpl) HasMetadata() bool              { return n.has_metadata }
func (n *newHandlerOptionImpl) Method() string                 { return n.method }
func (n *newHandlerOptionImpl) HasMethod() bool                { return n.has_method }
func (n *newHandlerOptionImpl) Renderer() Renderer             { return n.renderer }
func (n *newHandlerOptionImpl) HasRenderer() bool              { return n.has_renderer }
func (n *newHandlerOptionImpl) RendererConfig() RendererConfig { return n.rendererConfig }
func (n *newHandlerOptionImpl) HasRendererConfig() bool        { return n.has_rendererConfig }
func (n *newHandlerOptionImpl) WebOnly() bool                  { return n.webOnly }
func (n *newHandlerOptionImpl) HasWebOnly() bool               { return n.has_webOnly }

func makeNewHandlerOptionImpl(opts ...NewHandlerOption) *newHandlerOptionImpl {
	res := &newHandlerOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeNewHandlerOptions(opts ...NewHandlerOption) NewHandlerOptions {
	return makeNewHandlerOptionImpl(opts...)
}
