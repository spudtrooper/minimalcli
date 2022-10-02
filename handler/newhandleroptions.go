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

type newHandlerOptionImpl struct {
	cliOnly      bool
	has_cliOnly  bool
	metadata     HandlerMetadata
	has_metadata bool
	method       string
	has_method   bool
}

func (n *newHandlerOptionImpl) CliOnly() bool             { return n.cliOnly }
func (n *newHandlerOptionImpl) HasCliOnly() bool          { return n.has_cliOnly }
func (n *newHandlerOptionImpl) Metadata() HandlerMetadata { return n.metadata }
func (n *newHandlerOptionImpl) HasMetadata() bool         { return n.has_metadata }
func (n *newHandlerOptionImpl) Method() string            { return n.method }
func (n *newHandlerOptionImpl) HasMethod() bool           { return n.has_method }

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
