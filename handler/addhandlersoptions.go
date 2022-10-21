// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package handler

import "fmt"

type AddHandlersOption struct {
	f func(*addHandlersOptionImpl)
	s string
}

func (o AddHandlersOption) String() string { return o.s }

type AddHandlersOptions interface {
	EditName() string
	HasEditName() bool
	FooterHTML() string
	HasFooterHTML() bool
	FormatHTML() bool
	HasFormatHTML() bool
	HandlersFiles() []string
	HasHandlersFiles() bool
	HandlersFilesRoot() string
	HasHandlersFilesRoot() bool
	IndexName() string
	HasIndexName() bool
	IndexTitle() string
	HasIndexTitle() bool
	Key() string
	HasKey() bool
	Prefix() string
	HasPrefix() bool
	SerializedSourceLocations() []byte
	HasSerializedSourceLocations() bool
	SourceLinkURIRoot() string
	HasSourceLinkURIRoot() bool
	SourceLinks() bool
	HasSourceLinks() bool
	ToAddSectionOptions() []AddSectionOption
}

func AddHandlersEditName(editName string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_editName = true
		opts.editName = editName
	}, fmt.Sprintf("handler.AddHandlersEditName(string %+v)}", editName)}
}
func AddHandlersEditNameFlag(editName *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if editName == nil {
			return
		}
		opts.has_editName = true
		opts.editName = *editName
	}, fmt.Sprintf("handler.AddHandlersEditName(string %+v)}", editName)}
}

func AddHandlersFooterHTML(footerHTML string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_footerHTML = true
		opts.footerHTML = footerHTML
	}, fmt.Sprintf("handler.AddHandlersFooterHTML(string %+v)}", footerHTML)}
}
func AddHandlersFooterHTMLFlag(footerHTML *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if footerHTML == nil {
			return
		}
		opts.has_footerHTML = true
		opts.footerHTML = *footerHTML
	}, fmt.Sprintf("handler.AddHandlersFooterHTML(string %+v)}", footerHTML)}
}

func AddHandlersFormatHTML(formatHTML bool) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_formatHTML = true
		opts.formatHTML = formatHTML
	}, fmt.Sprintf("handler.AddHandlersFormatHTML(bool %+v)}", formatHTML)}
}
func AddHandlersFormatHTMLFlag(formatHTML *bool) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if formatHTML == nil {
			return
		}
		opts.has_formatHTML = true
		opts.formatHTML = *formatHTML
	}, fmt.Sprintf("handler.AddHandlersFormatHTML(bool %+v)}", formatHTML)}
}

func AddHandlersHandlersFiles(handlersFiles []string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_handlersFiles = true
		opts.handlersFiles = handlersFiles
	}, fmt.Sprintf("handler.AddHandlersHandlersFiles([]string %+v)}", handlersFiles)}
}
func AddHandlersHandlersFilesFlag(handlersFiles *[]string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if handlersFiles == nil {
			return
		}
		opts.has_handlersFiles = true
		opts.handlersFiles = *handlersFiles
	}, fmt.Sprintf("handler.AddHandlersHandlersFiles([]string %+v)}", handlersFiles)}
}

func AddHandlersHandlersFilesRoot(handlersFilesRoot string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_handlersFilesRoot = true
		opts.handlersFilesRoot = handlersFilesRoot
	}, fmt.Sprintf("handler.AddHandlersHandlersFilesRoot(string %+v)}", handlersFilesRoot)}
}
func AddHandlersHandlersFilesRootFlag(handlersFilesRoot *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if handlersFilesRoot == nil {
			return
		}
		opts.has_handlersFilesRoot = true
		opts.handlersFilesRoot = *handlersFilesRoot
	}, fmt.Sprintf("handler.AddHandlersHandlersFilesRoot(string %+v)}", handlersFilesRoot)}
}

func AddHandlersIndexName(indexName string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_indexName = true
		opts.indexName = indexName
	}, fmt.Sprintf("handler.AddHandlersIndexName(string %+v)}", indexName)}
}
func AddHandlersIndexNameFlag(indexName *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if indexName == nil {
			return
		}
		opts.has_indexName = true
		opts.indexName = *indexName
	}, fmt.Sprintf("handler.AddHandlersIndexName(string %+v)}", indexName)}
}

func AddHandlersIndexTitle(indexTitle string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_indexTitle = true
		opts.indexTitle = indexTitle
	}, fmt.Sprintf("handler.AddHandlersIndexTitle(string %+v)}", indexTitle)}
}
func AddHandlersIndexTitleFlag(indexTitle *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if indexTitle == nil {
			return
		}
		opts.has_indexTitle = true
		opts.indexTitle = *indexTitle
	}, fmt.Sprintf("handler.AddHandlersIndexTitle(string %+v)}", indexTitle)}
}

func AddHandlersKey(key string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_key = true
		opts.key = key
	}, fmt.Sprintf("handler.AddHandlersKey(string %+v)}", key)}
}
func AddHandlersKeyFlag(key *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if key == nil {
			return
		}
		opts.has_key = true
		opts.key = *key
	}, fmt.Sprintf("handler.AddHandlersKey(string %+v)}", key)}
}

func AddHandlersPrefix(prefix string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_prefix = true
		opts.prefix = prefix
	}, fmt.Sprintf("handler.AddHandlersPrefix(string %+v)}", prefix)}
}
func AddHandlersPrefixFlag(prefix *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if prefix == nil {
			return
		}
		opts.has_prefix = true
		opts.prefix = *prefix
	}, fmt.Sprintf("handler.AddHandlersPrefix(string %+v)}", prefix)}
}

func AddHandlersSerializedSourceLocations(serializedSourceLocations []byte) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_serializedSourceLocations = true
		opts.serializedSourceLocations = serializedSourceLocations
	}, fmt.Sprintf("handler.AddHandlersSerializedSourceLocations([]byte %+v)}", serializedSourceLocations)}
}
func AddHandlersSerializedSourceLocationsFlag(serializedSourceLocations *[]byte) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if serializedSourceLocations == nil {
			return
		}
		opts.has_serializedSourceLocations = true
		opts.serializedSourceLocations = *serializedSourceLocations
	}, fmt.Sprintf("handler.AddHandlersSerializedSourceLocations([]byte %+v)}", serializedSourceLocations)}
}

func AddHandlersSourceLinkURIRoot(sourceLinkURIRoot string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = sourceLinkURIRoot
	}, fmt.Sprintf("handler.AddHandlersSourceLinkURIRoot(string %+v)}", sourceLinkURIRoot)}
}
func AddHandlersSourceLinkURIRootFlag(sourceLinkURIRoot *string) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if sourceLinkURIRoot == nil {
			return
		}
		opts.has_sourceLinkURIRoot = true
		opts.sourceLinkURIRoot = *sourceLinkURIRoot
	}, fmt.Sprintf("handler.AddHandlersSourceLinkURIRoot(string %+v)}", sourceLinkURIRoot)}
}

func AddHandlersSourceLinks(sourceLinks bool) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		opts.has_sourceLinks = true
		opts.sourceLinks = sourceLinks
	}, fmt.Sprintf("handler.AddHandlersSourceLinks(bool %+v)}", sourceLinks)}
}
func AddHandlersSourceLinksFlag(sourceLinks *bool) AddHandlersOption {
	return AddHandlersOption{func(opts *addHandlersOptionImpl) {
		if sourceLinks == nil {
			return
		}
		opts.has_sourceLinks = true
		opts.sourceLinks = *sourceLinks
	}, fmt.Sprintf("handler.AddHandlersSourceLinks(bool %+v)}", sourceLinks)}
}

type addHandlersOptionImpl struct {
	editName                      string
	has_editName                  bool
	footerHTML                    string
	has_footerHTML                bool
	formatHTML                    bool
	has_formatHTML                bool
	handlersFiles                 []string
	has_handlersFiles             bool
	handlersFilesRoot             string
	has_handlersFilesRoot         bool
	indexName                     string
	has_indexName                 bool
	indexTitle                    string
	has_indexTitle                bool
	key                           string
	has_key                       bool
	prefix                        string
	has_prefix                    bool
	serializedSourceLocations     []byte
	has_serializedSourceLocations bool
	sourceLinkURIRoot             string
	has_sourceLinkURIRoot         bool
	sourceLinks                   bool
	has_sourceLinks               bool
}

func (a *addHandlersOptionImpl) EditName() string           { return a.editName }
func (a *addHandlersOptionImpl) HasEditName() bool          { return a.has_editName }
func (a *addHandlersOptionImpl) FooterHTML() string         { return a.footerHTML }
func (a *addHandlersOptionImpl) HasFooterHTML() bool        { return a.has_footerHTML }
func (a *addHandlersOptionImpl) FormatHTML() bool           { return a.formatHTML }
func (a *addHandlersOptionImpl) HasFormatHTML() bool        { return a.has_formatHTML }
func (a *addHandlersOptionImpl) HandlersFiles() []string    { return a.handlersFiles }
func (a *addHandlersOptionImpl) HasHandlersFiles() bool     { return a.has_handlersFiles }
func (a *addHandlersOptionImpl) HandlersFilesRoot() string  { return a.handlersFilesRoot }
func (a *addHandlersOptionImpl) HasHandlersFilesRoot() bool { return a.has_handlersFilesRoot }
func (a *addHandlersOptionImpl) IndexName() string          { return a.indexName }
func (a *addHandlersOptionImpl) HasIndexName() bool         { return a.has_indexName }
func (a *addHandlersOptionImpl) IndexTitle() string         { return a.indexTitle }
func (a *addHandlersOptionImpl) HasIndexTitle() bool        { return a.has_indexTitle }
func (a *addHandlersOptionImpl) Key() string                { return a.key }
func (a *addHandlersOptionImpl) HasKey() bool               { return a.has_key }
func (a *addHandlersOptionImpl) Prefix() string             { return a.prefix }
func (a *addHandlersOptionImpl) HasPrefix() bool            { return a.has_prefix }
func (a *addHandlersOptionImpl) SerializedSourceLocations() []byte {
	return a.serializedSourceLocations
}
func (a *addHandlersOptionImpl) HasSerializedSourceLocations() bool {
	return a.has_serializedSourceLocations
}
func (a *addHandlersOptionImpl) SourceLinkURIRoot() string  { return a.sourceLinkURIRoot }
func (a *addHandlersOptionImpl) HasSourceLinkURIRoot() bool { return a.has_sourceLinkURIRoot }
func (a *addHandlersOptionImpl) SourceLinks() bool          { return a.sourceLinks }
func (a *addHandlersOptionImpl) HasSourceLinks() bool       { return a.has_sourceLinks }

// ToAddSectionOptions converts AddHandlersOption to an array of AddSectionOption
func (o *addHandlersOptionImpl) ToAddSectionOptions() []AddSectionOption {
	return []AddSectionOption{
		AddSectionEditName(o.EditName()),
		AddSectionFooterHTML(o.FooterHTML()),
		AddSectionFormatHTML(o.FormatHTML()),
		AddSectionHandlersFiles(o.HandlersFiles()),
		AddSectionHandlersFilesRoot(o.HandlersFilesRoot()),
		AddSectionIndexName(o.IndexName()),
		AddSectionKey(o.Key()),
		AddSectionSerializedSourceLocations(o.SerializedSourceLocations()),
		AddSectionSourceLinkURIRoot(o.SourceLinkURIRoot()),
		AddSectionSourceLinks(o.SourceLinks()),
	}
}

func makeAddHandlersOptionImpl(opts ...AddHandlersOption) *addHandlersOptionImpl {
	res := &addHandlersOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAddHandlersOptions(opts ...AddHandlersOption) AddHandlersOptions {
	return makeAddHandlersOptionImpl(opts...)
}
