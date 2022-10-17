// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package app

import "fmt"

type MakeOption struct {
	f func(*makeOptionImpl)
	s string
}

func (o MakeOption) String() string { return o.s }

type MakeOptions interface {
	PrintBanner() bool
	HasPrintBanner() bool
}

func MakePrintBanner(printBanner bool) MakeOption {
	return MakeOption{func(opts *makeOptionImpl) {
		opts.has_printBanner = true
		opts.printBanner = printBanner
	}, fmt.Sprintf("app.MakePrintBanner(bool %+v)}", printBanner)}
}
func MakePrintBannerFlag(printBanner *bool) MakeOption {
	return MakeOption{func(opts *makeOptionImpl) {
		if printBanner == nil {
			return
		}
		opts.has_printBanner = true
		opts.printBanner = *printBanner
	}, fmt.Sprintf("app.MakePrintBanner(bool %+v)}", printBanner)}
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
		opt.f(res)
	}
	return res
}

func MakeMakeOptions(opts ...MakeOption) MakeOptions {
	return makeMakeOptionImpl(opts...)
}
