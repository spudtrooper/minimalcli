// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

type NewHandlerOption func(*newHandlerOptionImpl)

type NewHandlerOptions interface {
	CliOnly() bool
	HasCliOnly() bool
	Metadata() HandlerMetadata
	HasMetadata() bool
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

type newHandlerOptionImpl struct {
	cliOnly      bool
	has_cliOnly  bool
	metadata     HandlerMetadata
	has_metadata bool
}

func (n *newHandlerOptionImpl) CliOnly() bool             { return n.cliOnly }
func (n *newHandlerOptionImpl) HasCliOnly() bool          { return n.has_cliOnly }
func (n *newHandlerOptionImpl) Metadata() HandlerMetadata { return n.metadata }
func (n *newHandlerOptionImpl) HasMetadata() bool         { return n.has_metadata }

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
