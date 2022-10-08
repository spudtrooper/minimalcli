// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

type CreateHandlerOption func(*createHandlerOptionImpl)

type CreateHandlerOptions interface {
	IndexTitle() string
	HasIndexTitle() bool
	Prefix() string
	HasPrefix() bool
	IndexName() string
	HasIndexName() bool
	EditName() string
	HasEditName() bool
	FooterHTML() string
	HasFooterHTML() bool
	SourceLinks() bool
	HasSourceLinks() bool
	HandlersFiles() []string
	HasHandlersFiles() bool
	SourceLinkURIRoot() string
	HasSourceLinkURIRoot() bool
	FormatHTML() bool
	HasFormatHTML() bool
}

func CreateHandlerIndexTitle(indexTitle string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_indexTitle = true
		opts.indexTitle = indexTitle
	}
}
func CreateHandlerIndexTitleFlag(indexTitle *string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if indexTitle == nil {
			return
		}
		opts.has_indexTitle = true
		opts.indexTitle = *indexTitle
	}
}

func CreateHandlerPrefix(prefix string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_prefix = true
		opts.prefix = prefix
	}
}
func CreateHandlerPrefixFlag(prefix *string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if prefix == nil {
			return
		}
		opts.has_prefix = true
		opts.prefix = *prefix
	}
}

func CreateHandlerIndexName(indexName string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_indexName = true
		opts.indexName = indexName
	}
}
func CreateHandlerIndexNameFlag(indexName *string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if indexName == nil {
			return
		}
		opts.has_indexName = true
		opts.indexName = *indexName
	}
}

func CreateHandlerEditName(editName string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_editName = true
		opts.editName = editName
	}
}
func CreateHandlerEditNameFlag(editName *string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if editName == nil {
			return
		}
		opts.has_editName = true
		opts.editName = *editName
	}
}

func CreateHandlerFooterHTML(footerHTML string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}
}
func CreateHandlerFooterHTMLFlag(footerHTML *string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}
}

func CreateHandlerSourceLinks(sourceLinks bool) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_sourceLinks = true
		opts.sourceLinks = sourceLinks
	}
}
func CreateHandlerSourceLinksFlag(sourceLinks *bool) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if sourceLinks == nil {
			return
		}
		opts.has_sourceLinks = true
		opts.sourceLinks = *sourceLinks
	}
}

func CreateHandlerHandlersFiles(handlersFiles []string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_handlersFiles = true
		opts.handlersFiles = handlersFiles
	}
}
func CreateHandlerHandlersFilesFlag(handlersFiles *[]string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if handlersFiles == nil {
			return
		}
		opts.has_handlersFiles = true
		opts.handlersFiles = *handlersFiles
	}
}

func CreateHandlerSourceLinkURIRoot(sourceLinkURIRoot string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = sourceLinkURIRoot
	}
}
func CreateHandlerSourceLinkURIRootFlag(sourceLinkURIRoot *string) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if sourceLinkURIRoot == nil {
			return
		}
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = *sourceLinkURIRoot
	}
}

func CreateHandlerFormatHTML(formatHTML bool) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}
}
func CreateHandlerFormatHTMLFlag(formatHTML *bool) CreateHandlerOption {
	return func(opts *createHandlerOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}
}

type createHandlerOptionImpl struct {
	indexTitle            string
	has_indexTitle        bool
	prefix                string
	has_prefix            bool
	indexName             string
	has_indexName         bool
	editName              string
	has_editName          bool
	footerHTML            string
	has_footerHTML        bool
	sourceLinks           bool
	has_sourceLinks       bool
	handlersFiles         []string
	has_handlersFiles     bool
	sourceLinkURIRoot     string
	has_sourceLinkURIRoot bool
	formatHTML            bool
	has_formatHTML        bool
}

func (c *createHandlerOptionImpl) IndexTitle() string         { return c.indexTitle }
func (c *createHandlerOptionImpl) HasIndexTitle() bool        { return c.has_indexTitle }
func (c *createHandlerOptionImpl) Prefix() string             { return c.prefix }
func (c *createHandlerOptionImpl) HasPrefix() bool            { return c.has_prefix }
func (c *createHandlerOptionImpl) IndexName() string          { return c.indexName }
func (c *createHandlerOptionImpl) HasIndexName() bool         { return c.has_indexName }
func (c *createHandlerOptionImpl) EditName() string           { return c.editName }
func (c *createHandlerOptionImpl) HasEditName() bool          { return c.has_editName }
func (c *createHandlerOptionImpl) FooterHTML() string         { return c.footerHTML }
func (c *createHandlerOptionImpl) HasFooterHTML() bool        { return c.has_footerHTML }
func (c *createHandlerOptionImpl) SourceLinks() bool          { return c.sourceLinks }
func (c *createHandlerOptionImpl) HasSourceLinks() bool       { return c.has_sourceLinks }
func (c *createHandlerOptionImpl) HandlersFiles() []string    { return c.handlersFiles }
func (c *createHandlerOptionImpl) HasHandlersFiles() bool     { return c.has_handlersFiles }
func (c *createHandlerOptionImpl) SourceLinkURIRoot() string  { return c.sourceLinkURIRoot }
func (c *createHandlerOptionImpl) HasSourceLinkURIRoot() bool { return c.has_sourceLinkURIRoot }
func (c *createHandlerOptionImpl) FormatHTML() bool           { return c.formatHTML }
func (c *createHandlerOptionImpl) HasFormatHTML() bool        { return c.has_formatHTML }

func makeCreateHandlerOptionImpl(opts ...CreateHandlerOption) *createHandlerOptionImpl {
	res := &createHandlerOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeCreateHandlerOptions(opts ...CreateHandlerOption) CreateHandlerOptions {
	return makeCreateHandlerOptionImpl(opts...)
}
