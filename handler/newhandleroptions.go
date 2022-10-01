// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

type NewHandlerOption func(*newHandlerOptionImpl)

type NewHandlerOptions interface {
	CliOnly() bool
	HasCliOnly() bool
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

type newHandlerOptionImpl struct {
	cliOnly     bool
	has_cliOnly bool
}

func (n *newHandlerOptionImpl) CliOnly() bool    { return n.cliOnly }
func (n *newHandlerOptionImpl) HasCliOnly() bool { return n.has_cliOnly }

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
