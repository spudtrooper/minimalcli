// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

type NewHandlerOption func(*newHandlerOptionImpl)

type NewHandlerOptions interface {
	CliOnly() bool
	HasCliOnly() bool
	Metadata() HandlerMetadata
	HasMetadata() bool
	Method() string
	HasMethod() bool
	ExtraRequiredFields() []string
	HasExtraRequiredFields() bool
	Renderer() Renderer
	HasRenderer() bool
}

func NewHandlerCliOnly(cliOnly bool) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		opts.has_cliOnly = true
		opts.cliOnly = cliOnly
	}
}
func NewHandlerCliOnlyFlag(cliOnly *bool) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		if cliOnly == nil {
			return
		}
		opts.has_cliOnly = true
		opts.cliOnly = *cliOnly
	}
}

func NewHandlerMetadata(metadata HandlerMetadata) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		opts.has_metadata = true
		opts.metadata = metadata
	}
}
func NewHandlerMetadataFlag(metadata *HandlerMetadata) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		if metadata == nil {
			return
		}
		opts.has_metadata = true
		opts.metadata = *metadata
	}
}

func NewHandlerMethod(method string) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		opts.has_method = true
		opts.method = method
	}
}
func NewHandlerMethodFlag(method *string) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		if method == nil {
			return
		}
		opts.has_method = true
		opts.method = *method
	}
}

func NewHandlerExtraRequiredFields(extraRequiredFields []string) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		opts.has_extraRequiredFields = true
		opts.extraRequiredFields = extraRequiredFields
	}
}
func NewHandlerExtraRequiredFieldsFlag(extraRequiredFields *[]string) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		if extraRequiredFields == nil {
			return
		}
		opts.has_extraRequiredFields = true
		opts.extraRequiredFields = *extraRequiredFields
	}
}

func NewHandlerRenderer(renderer Renderer) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		opts.has_renderer = true
		opts.renderer = renderer
	}
}
func NewHandlerRendererFlag(renderer *Renderer) NewHandlerOption {
	return func(opts *newHandlerOptionImpl) {
		if renderer == nil {
			return
		}
		opts.has_renderer = true
		opts.renderer = *renderer
	}
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
}

func (n *newHandlerOptionImpl) CliOnly() bool                 { return n.cliOnly }
func (n *newHandlerOptionImpl) HasCliOnly() bool              { return n.has_cliOnly }
func (n *newHandlerOptionImpl) Metadata() HandlerMetadata     { return n.metadata }
func (n *newHandlerOptionImpl) HasMetadata() bool             { return n.has_metadata }
func (n *newHandlerOptionImpl) Method() string                { return n.method }
func (n *newHandlerOptionImpl) HasMethod() bool               { return n.has_method }
func (n *newHandlerOptionImpl) ExtraRequiredFields() []string { return n.extraRequiredFields }
func (n *newHandlerOptionImpl) HasExtraRequiredFields() bool  { return n.has_extraRequiredFields }
func (n *newHandlerOptionImpl) Renderer() Renderer            { return n.renderer }
func (n *newHandlerOptionImpl) HasRenderer() bool             { return n.has_renderer }

func makeNewHandlerOptionImpl(opts ...NewHandlerOption) *newHandlerOptionImpl {
	res := &newHandlerOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeNewHandlerOptions(opts ...NewHandlerOption) NewHandlerOptions {
	return makeNewHandlerOptionImpl(opts...)
}
