// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type GenIndexOption struct {
	f func(*genIndexOptionImpl)
	s string
}

func (o GenIndexOption) String() string { return o.s }

type GenIndexOptions interface {
	FooterHTML() string
	HasFooterHTML() bool
	FormatHTML() bool
	HasFormatHTML() bool
	Route() string
	HasRoute() bool
	Title() string
	HasTitle() bool
}

func GenIndexFooterHTML(footerHTML string) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}, fmt.Sprintf("handler.GenIndexFooterHTML(string %+v)}", footerHTML)}
}
func GenIndexFooterHTMLFlag(footerHTML *string) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}, fmt.Sprintf("handler.GenIndexFooterHTML(string %+v)}", footerHTML)}
}

func GenIndexFormatHTML(formatHTML bool) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}, fmt.Sprintf("handler.GenIndexFormatHTML(bool %+v)}", formatHTML)}
}
func GenIndexFormatHTMLFlag(formatHTML *bool) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}, fmt.Sprintf("handler.GenIndexFormatHTML(bool %+v)}", formatHTML)}
}

func GenIndexRoute(route string) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		opts.has_route = true
		opts.route = route
	}, fmt.Sprintf("handler.GenIndexRoute(string %+v)}", route)}
}
func GenIndexRouteFlag(route *string) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		if route == nil {
			return
		}
		opts.has_route = true
		opts.route = *route
	}, fmt.Sprintf("handler.GenIndexRoute(string %+v)}", route)}
}

func GenIndexTitle(title string) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		opts.has_title = true
		opts.title = title
	}, fmt.Sprintf("handler.GenIndexTitle(string %+v)}", title)}
}
func GenIndexTitleFlag(title *string) GenIndexOption {
	return GenIndexOption{func(opts *genIndexOptionImpl) {
		if title == nil {
			return
		}
		opts.has_title = true
		opts.title = *title
	}, fmt.Sprintf("handler.GenIndexTitle(string %+v)}", title)}
}

type genIndexOptionImpl struct {
	footerHTML     string
	has_footerHTML bool
	formatHTML     bool
	has_formatHTML bool
	route          string
	has_route      bool
	title          string
	has_title      bool
}

func (g *genIndexOptionImpl) FooterHTML() string  { return g.footerHTML }
func (g *genIndexOptionImpl) HasFooterHTML() bool { return g.has_footerHTML }
func (g *genIndexOptionImpl) FormatHTML() bool    { return g.formatHTML }
func (g *genIndexOptionImpl) HasFormatHTML() bool { return g.has_formatHTML }
func (g *genIndexOptionImpl) Route() string       { return or.String(g.route, "/") }
func (g *genIndexOptionImpl) HasRoute() bool      { return g.has_route }
func (g *genIndexOptionImpl) Title() string       { return or.String(g.title, "Index") }
func (g *genIndexOptionImpl) HasTitle() bool      { return g.has_title }

func makeGenIndexOptionImpl(opts ...GenIndexOption) *genIndexOptionImpl {
	res := &genIndexOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeGenIndexOptions(opts ...GenIndexOption) GenIndexOptions {
	return makeGenIndexOptionImpl(opts...)
}
