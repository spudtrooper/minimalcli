// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

type AddHandlersOption func(*addHandlersOptionImpl)

type AddHandlersOptions interface {
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
	HandlersFilesRoot() string
	HasHandlersFilesRoot() bool
	SourceLinkURIRoot() string
	HasSourceLinkURIRoot() bool
	FormatHTML() bool
	HasFormatHTML() bool
	ToAddSectionOptions() []AddSectionOption
}

func AddHandlersIndexTitle(indexTitle string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_indexTitle = true
		opts.indexTitle = indexTitle
	}
}
func AddHandlersIndexTitleFlag(indexTitle *string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if indexTitle == nil {
			return
		}
		opts.has_indexTitle = true
		opts.indexTitle = *indexTitle
	}
}

func AddHandlersPrefix(prefix string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_prefix = true
		opts.prefix = prefix
	}
}
func AddHandlersPrefixFlag(prefix *string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if prefix == nil {
			return
		}
		opts.has_prefix = true
		opts.prefix = *prefix
	}
}

func AddHandlersIndexName(indexName string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_indexName = true
		opts.indexName = indexName
	}
}
func AddHandlersIndexNameFlag(indexName *string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if indexName == nil {
			return
		}
		opts.has_indexName = true
		opts.indexName = *indexName
	}
}

func AddHandlersEditName(editName string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_editName = true
		opts.editName = editName
	}
}
func AddHandlersEditNameFlag(editName *string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if editName == nil {
			return
		}
		opts.has_editName = true
		opts.editName = *editName
	}
}

func AddHandlersFooterHTML(footerHTML string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}
}
func AddHandlersFooterHTMLFlag(footerHTML *string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}
}

func AddHandlersSourceLinks(sourceLinks bool) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_sourceLinks = true
		opts.sourceLinks = sourceLinks
	}
}
func AddHandlersSourceLinksFlag(sourceLinks *bool) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if sourceLinks == nil {
			return
		}
		opts.has_sourceLinks = true
		opts.sourceLinks = *sourceLinks
	}
}

func AddHandlersHandlersFiles(handlersFiles []string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_handlersFiles = true
		opts.handlersFiles = handlersFiles
	}
}
func AddHandlersHandlersFilesFlag(handlersFiles *[]string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if handlersFiles == nil {
			return
		}
		opts.has_handlersFiles = true
		opts.handlersFiles = *handlersFiles
	}
}

func AddHandlersHandlersFilesRoot(handlersFilesRoot string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_handlersFilesRoot = true
		opts.handlersFilesRoot = handlersFilesRoot
	}
}
func AddHandlersHandlersFilesRootFlag(handlersFilesRoot *string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if handlersFilesRoot == nil {
			return
		}
		opts.has_handlersFilesRoot = true
		opts.handlersFilesRoot = *handlersFilesRoot
	}
}

func AddHandlersSourceLinkURIRoot(sourceLinkURIRoot string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = sourceLinkURIRoot
	}
}
func AddHandlersSourceLinkURIRootFlag(sourceLinkURIRoot *string) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if sourceLinkURIRoot == nil {
			return
		}
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = *sourceLinkURIRoot
	}
}

func AddHandlersFormatHTML(formatHTML bool) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}
}
func AddHandlersFormatHTMLFlag(formatHTML *bool) AddHandlersOption {
	return func(opts *addHandlersOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}
}

type addHandlersOptionImpl struct {
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
	handlersFilesRoot     string
	has_handlersFilesRoot bool
	sourceLinkURIRoot     string
	has_sourceLinkURIRoot bool
	formatHTML            bool
	has_formatHTML        bool
}

func (a *addHandlersOptionImpl) IndexTitle() string         { return a.indexTitle }
func (a *addHandlersOptionImpl) HasIndexTitle() bool        { return a.has_indexTitle }
func (a *addHandlersOptionImpl) Prefix() string             { return a.prefix }
func (a *addHandlersOptionImpl) HasPrefix() bool            { return a.has_prefix }
func (a *addHandlersOptionImpl) IndexName() string          { return a.indexName }
func (a *addHandlersOptionImpl) HasIndexName() bool         { return a.has_indexName }
func (a *addHandlersOptionImpl) EditName() string           { return a.editName }
func (a *addHandlersOptionImpl) HasEditName() bool          { return a.has_editName }
func (a *addHandlersOptionImpl) FooterHTML() string         { return a.footerHTML }
func (a *addHandlersOptionImpl) HasFooterHTML() bool        { return a.has_footerHTML }
func (a *addHandlersOptionImpl) SourceLinks() bool          { return a.sourceLinks }
func (a *addHandlersOptionImpl) HasSourceLinks() bool       { return a.has_sourceLinks }
func (a *addHandlersOptionImpl) HandlersFiles() []string    { return a.handlersFiles }
func (a *addHandlersOptionImpl) HasHandlersFiles() bool     { return a.has_handlersFiles }
func (a *addHandlersOptionImpl) HandlersFilesRoot() string  { return a.handlersFilesRoot }
func (a *addHandlersOptionImpl) HasHandlersFilesRoot() bool { return a.has_handlersFilesRoot }
func (a *addHandlersOptionImpl) SourceLinkURIRoot() string  { return a.sourceLinkURIRoot }
func (a *addHandlersOptionImpl) HasSourceLinkURIRoot() bool { return a.has_sourceLinkURIRoot }
func (a *addHandlersOptionImpl) FormatHTML() bool           { return a.formatHTML }
func (a *addHandlersOptionImpl) HasFormatHTML() bool        { return a.has_formatHTML }

// ToAddSectionOptions converts AddHandlersOption to an array of AddSectionOption
func (o *addHandlersOptionImpl) ToAddSectionOptions() []AddSectionOption {
	return []AddSectionOption{
		AddSectionIndexName(o.IndexName()),
		AddSectionEditName(o.EditName()),
		AddSectionFooterHTML(o.FooterHTML()),
		AddSectionSourceLinks(o.SourceLinks()),
		AddSectionHandlersFiles(o.HandlersFiles()),
		AddSectionHandlersFilesRoot(o.HandlersFilesRoot()),
		AddSectionSourceLinkURIRoot(o.SourceLinkURIRoot()),
		AddSectionFormatHTML(o.FormatHTML()),
	}
}

func makeAddHandlersOptionImpl(opts ...AddHandlersOption) *addHandlersOptionImpl {
	res := &addHandlersOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAddHandlersOptions(opts ...AddHandlersOption) AddHandlersOptions {
	return makeAddHandlersOptionImpl(opts...)
}
