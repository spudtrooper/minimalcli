// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

type AddSectionOption func(*addSectionOptionImpl)

type AddSectionOptions interface {
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
	Key() string
	HasKey() bool
	SerializedSourceLocations() []byte
	HasSerializedSourceLocations() bool
}

func AddSectionIndexName(indexName string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_indexName = true
		opts.indexName = indexName
	}
}
func AddSectionIndexNameFlag(indexName *string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if indexName == nil {
			return
		}
		opts.has_indexName = true
		opts.indexName = *indexName
	}
}

func AddSectionEditName(editName string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_editName = true
		opts.editName = editName
	}
}
func AddSectionEditNameFlag(editName *string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if editName == nil {
			return
		}
		opts.has_editName = true
		opts.editName = *editName
	}
}

func AddSectionFooterHTML(footerHTML string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}
}
func AddSectionFooterHTMLFlag(footerHTML *string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}
}

func AddSectionSourceLinks(sourceLinks bool) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_sourceLinks = true
		opts.sourceLinks = sourceLinks
	}
}
func AddSectionSourceLinksFlag(sourceLinks *bool) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if sourceLinks == nil {
			return
		}
		opts.has_sourceLinks = true
		opts.sourceLinks = *sourceLinks
	}
}

func AddSectionHandlersFiles(handlersFiles []string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_handlersFiles = true
		opts.handlersFiles = handlersFiles
	}
}
func AddSectionHandlersFilesFlag(handlersFiles *[]string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if handlersFiles == nil {
			return
		}
		opts.has_handlersFiles = true
		opts.handlersFiles = *handlersFiles
	}
}

func AddSectionHandlersFilesRoot(handlersFilesRoot string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_handlersFilesRoot = true
		opts.handlersFilesRoot = handlersFilesRoot
	}
}
func AddSectionHandlersFilesRootFlag(handlersFilesRoot *string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if handlersFilesRoot == nil {
			return
		}
		opts.has_handlersFilesRoot = true
		opts.handlersFilesRoot = *handlersFilesRoot
	}
}

func AddSectionSourceLinkURIRoot(sourceLinkURIRoot string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = sourceLinkURIRoot
	}
}
func AddSectionSourceLinkURIRootFlag(sourceLinkURIRoot *string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if sourceLinkURIRoot == nil {
			return
		}
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = *sourceLinkURIRoot
	}
}

func AddSectionFormatHTML(formatHTML bool) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}
}
func AddSectionFormatHTMLFlag(formatHTML *bool) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}
}

func AddSectionKey(key string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_key = true
		opts.key = key
	}
}
func AddSectionKeyFlag(key *string) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if key == nil {
			return
		}
		opts.has_key = true
		opts.key = *key
	}
}

func AddSectionSerializedSourceLocations(serializedSourceLocations []byte) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		opts.has_serializedSourceLocations = true
		opts.serializedSourceLocations = serializedSourceLocations
	}
}
func AddSectionSerializedSourceLocationsFlag(serializedSourceLocations *[]byte) AddSectionOption {
	return func(opts *addSectionOptionImpl) {
		if serializedSourceLocations == nil {
			return
		}
		opts.has_serializedSourceLocations = true
		opts.serializedSourceLocations = *serializedSourceLocations
	}
}

type addSectionOptionImpl struct {
	indexName                     string
	has_indexName                 bool
	editName                      string
	has_editName                  bool
	footerHTML                    string
	has_footerHTML                bool
	sourceLinks                   bool
	has_sourceLinks               bool
	handlersFiles                 []string
	has_handlersFiles             bool
	handlersFilesRoot             string
	has_handlersFilesRoot         bool
	sourceLinkURIRoot             string
	has_sourceLinkURIRoot         bool
	formatHTML                    bool
	has_formatHTML                bool
	key                           string
	has_key                       bool
	serializedSourceLocations     []byte
	has_serializedSourceLocations bool
}

func (a *addSectionOptionImpl) IndexName() string                 { return a.indexName }
func (a *addSectionOptionImpl) HasIndexName() bool                { return a.has_indexName }
func (a *addSectionOptionImpl) EditName() string                  { return a.editName }
func (a *addSectionOptionImpl) HasEditName() bool                 { return a.has_editName }
func (a *addSectionOptionImpl) FooterHTML() string                { return a.footerHTML }
func (a *addSectionOptionImpl) HasFooterHTML() bool               { return a.has_footerHTML }
func (a *addSectionOptionImpl) SourceLinks() bool                 { return a.sourceLinks }
func (a *addSectionOptionImpl) HasSourceLinks() bool              { return a.has_sourceLinks }
func (a *addSectionOptionImpl) HandlersFiles() []string           { return a.handlersFiles }
func (a *addSectionOptionImpl) HasHandlersFiles() bool            { return a.has_handlersFiles }
func (a *addSectionOptionImpl) HandlersFilesRoot() string         { return a.handlersFilesRoot }
func (a *addSectionOptionImpl) HasHandlersFilesRoot() bool        { return a.has_handlersFilesRoot }
func (a *addSectionOptionImpl) SourceLinkURIRoot() string         { return a.sourceLinkURIRoot }
func (a *addSectionOptionImpl) HasSourceLinkURIRoot() bool        { return a.has_sourceLinkURIRoot }
func (a *addSectionOptionImpl) FormatHTML() bool                  { return a.formatHTML }
func (a *addSectionOptionImpl) HasFormatHTML() bool               { return a.has_formatHTML }
func (a *addSectionOptionImpl) Key() string                       { return a.key }
func (a *addSectionOptionImpl) HasKey() bool                      { return a.has_key }
func (a *addSectionOptionImpl) SerializedSourceLocations() []byte { return a.serializedSourceLocations }
func (a *addSectionOptionImpl) HasSerializedSourceLocations() bool {
	return a.has_serializedSourceLocations
}

func makeAddSectionOptionImpl(opts ...AddSectionOption) *addSectionOptionImpl {
	res := &addSectionOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAddSectionOptions(opts ...AddSectionOption) AddSectionOptions {
	return makeAddSectionOptionImpl(opts...)
}
