// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

import (
	"github.com/spudtrooper/goutil/or"
)

type GenAllOption func(*genAllOptionImpl)

type GenAllOptions interface {
	Title() string
	HasTitle() bool
	FooterHTML() string
	HasFooterHTML() bool
	FormatHTML() bool
	HasFormatHTML() bool
	Route() string
	HasRoute() bool
}

func GenAllTitle(title string) GenAllOption {
	return func(opts *genAllOptionImpl) {
		opts.has_title = true
		opts.title = title
	}
}
func GenAllTitleFlag(title *string) GenAllOption {
	return func(opts *genAllOptionImpl) {
		if title == nil {
			return
		}
		opts.has_title = true
		opts.title = *title
	}
}

func GenAllFooterHTML(footerHTML string) GenAllOption {
	return func(opts *genAllOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}
}
func GenAllFooterHTMLFlag(footerHTML *string) GenAllOption {
	return func(opts *genAllOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}
}

func GenAllFormatHTML(formatHTML bool) GenAllOption {
	return func(opts *genAllOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}
}
func GenAllFormatHTMLFlag(formatHTML *bool) GenAllOption {
	return func(opts *genAllOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}
}

func GenAllRoute(route string) GenAllOption {
	return func(opts *genAllOptionImpl) {
		opts.has_route = true
		opts.route = route
	}
}
func GenAllRouteFlag(route *string) GenAllOption {
	return func(opts *genAllOptionImpl) {
		if route == nil {
			return
		}
		opts.has_route = true
		opts.route = *route
	}
}

type genAllOptionImpl struct {
	title          string
	has_title      bool
	footerHTML     string
	has_footerHTML bool
	formatHTML     bool
	has_formatHTML bool
	route          string
	has_route      bool
}

func (g *genAllOptionImpl) Title() string       { return or.String(g.title, "All") }
func (g *genAllOptionImpl) HasTitle() bool      { return g.has_title }
func (g *genAllOptionImpl) FooterHTML() string  { return g.footerHTML }
func (g *genAllOptionImpl) HasFooterHTML() bool { return g.has_footerHTML }
func (g *genAllOptionImpl) FormatHTML() bool    { return g.formatHTML }
func (g *genAllOptionImpl) HasFormatHTML() bool { return g.has_formatHTML }
func (g *genAllOptionImpl) Route() string       { return or.String(g.route, "/_all") }
func (g *genAllOptionImpl) HasRoute() bool      { return g.has_route }

func makeGenAllOptionImpl(opts ...GenAllOption) *genAllOptionImpl {
	res := &genAllOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGenAllOptions(opts ...GenAllOption) GenAllOptions {
	return makeGenAllOptionImpl(opts...)
}
