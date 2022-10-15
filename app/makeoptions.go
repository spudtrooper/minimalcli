// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package app

type MakeOption func(*makeOptionImpl)

type MakeOptions interface {
	PrintBanner() bool
	HasPrintBanner() bool
}

func MakePrintBanner(printBanner bool) MakeOption {
	return func(opts *makeOptionImpl) {
		opts.has_printBanner = true
		opts.printBanner = printBanner
	}
}
func MakePrintBannerFlag(printBanner *bool) MakeOption {
	return func(opts *makeOptionImpl) {
		if printBanner == nil {
			return
		}
		opts.has_printBanner = true
		opts.printBanner = *printBanner
	}
}

type makeOptionImpl struct {
	printBanner     bool
	has_printBanner bool
}

func (m *makeOptionImpl) PrintBanner() bool    { return m.printBanner }
func (m *makeOptionImpl) HasPrintBanner() bool { return m.has_printBanner }

func makeMakeOptionImpl(opts ...MakeOption) *makeOptionImpl {
	res := &makeOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeMakeOptions(opts ...MakeOption) MakeOptions {
	return makeMakeOptionImpl(opts...)
}
