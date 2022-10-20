// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type GenAllOption struct {
	f func(*genAllOptionImpl)
	s string
}

func (o GenAllOption) String() string { return o.s }

type GenAllOptions interface {
	FooterHTML() string
	HasFooterHTML() bool
	FormatHTML() bool
	HasFormatHTML() bool
	Route() string
	HasRoute() bool
	Title() string
	HasTitle() bool
}

func GenAllFooterHTML(footerHTML string) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}, fmt.Sprintf("handler.GenAllFooterHTML(string %+v)}", footerHTML)}
}
func GenAllFooterHTMLFlag(footerHTML *string) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}, fmt.Sprintf("handler.GenAllFooterHTML(string %+v)}", footerHTML)}
}

func GenAllFormatHTML(formatHTML bool) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}, fmt.Sprintf("handler.GenAllFormatHTML(bool %+v)}", formatHTML)}
}
func GenAllFormatHTMLFlag(formatHTML *bool) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}, fmt.Sprintf("handler.GenAllFormatHTML(bool %+v)}", formatHTML)}
}

func GenAllRoute(route string) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		opts.has_route = true
		opts.route = route
	}, fmt.Sprintf("handler.GenAllRoute(string %+v)}", route)}
}
func GenAllRouteFlag(route *string) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		if route == nil {
			return
		}
		opts.has_route = true
		opts.route = *route
	}, fmt.Sprintf("handler.GenAllRoute(string %+v)}", route)}
}

func GenAllTitle(title string) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		opts.has_title = true
		opts.title = title
	}, fmt.Sprintf("handler.GenAllTitle(string %+v)}", title)}
}
func GenAllTitleFlag(title *string) GenAllOption {
	return GenAllOption{func(opts *genAllOptionImpl) {
		if title == nil {
			return
		}
		opts.has_title = true
		opts.title = *title
	}, fmt.Sprintf("handler.GenAllTitle(string %+v)}", title)}
}

type genAllOptionImpl struct {
	footerHTML     string
	has_footerHTML bool
	formatHTML     bool
	has_formatHTML bool
	route          string
	has_route      bool
	title          string
	has_title      bool
}

func (g *genAllOptionImpl) FooterHTML() string  { return g.footerHTML }
func (g *genAllOptionImpl) HasFooterHTML() bool { return g.has_footerHTML }
func (g *genAllOptionImpl) FormatHTML() bool    { return g.formatHTML }
func (g *genAllOptionImpl) HasFormatHTML() bool { return g.has_formatHTML }
func (g *genAllOptionImpl) Route() string       { return or.String(g.route, "/_all") }
func (g *genAllOptionImpl) HasRoute() bool      { return g.has_route }
func (g *genAllOptionImpl) Title() string       { return or.String(g.title, "All") }
func (g *genAllOptionImpl) HasTitle() bool      { return g.has_title }

func makeGenAllOptionImpl(opts ...GenAllOption) *genAllOptionImpl {
	res := &genAllOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGenAllOptions(opts ...GenAllOption) GenAllOptions {
	return makeGenAllOptionImpl(opts...)
}
