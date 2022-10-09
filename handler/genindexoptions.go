// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

type GenIndexOption func(*genIndexOptionImpl)

type GenIndexOptions interface {
	Title() string
	HasTitle() bool
	FooterHTML() string
	HasFooterHTML() bool
	FormatHTML() bool
	HasFormatHTML() bool
	Route() string
	HasRoute() bool
}

func GenIndexTitle(title string) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		opts.has_title = true
		opts.title = title
	}
}
func GenIndexTitleFlag(title *string) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		if title == nil {
			return
		}
		opts.has_title = true
		opts.title = *title
	}
}

func GenIndexFooterHTML(footerHTML string) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}
}
func GenIndexFooterHTMLFlag(footerHTML *string) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}
}

func GenIndexFormatHTML(formatHTML bool) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}
}
func GenIndexFormatHTMLFlag(formatHTML *bool) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}
}

func GenIndexRoute(route string) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		opts.has_route = true
		opts.route = route
	}
}
func GenIndexRouteFlag(route *string) GenIndexOption {
	return func(opts *genIndexOptionImpl) {
		if route == nil {
			return
		}
		opts.has_route = true
		opts.route = *route
	}
}

type genIndexOptionImpl struct {
	title          string
	has_title      bool
	footerHTML     string
	has_footerHTML bool
	formatHTML     bool
	has_formatHTML bool
	route          string
	has_route      bool
}

func (g *genIndexOptionImpl) Title() string       { return g.title }
func (g *genIndexOptionImpl) HasTitle() bool      { return g.has_title }
func (g *genIndexOptionImpl) FooterHTML() string  { return g.footerHTML }
func (g *genIndexOptionImpl) HasFooterHTML() bool { return g.has_footerHTML }
func (g *genIndexOptionImpl) FormatHTML() bool    { return g.formatHTML }
func (g *genIndexOptionImpl) HasFormatHTML() bool { return g.has_formatHTML }
func (g *genIndexOptionImpl) Route() string       { return g.route }
func (g *genIndexOptionImpl) HasRoute() bool      { return g.has_route }

func makeGenIndexOptionImpl(opts ...GenIndexOption) *genIndexOptionImpl {
	res := &genIndexOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeGenIndexOptions(opts ...GenIndexOption) GenIndexOptions {
	return makeGenIndexOptionImpl(opts...)
}
