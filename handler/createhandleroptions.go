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

type createHandlerOptionImpl struct {
	indexTitle     string
	has_indexTitle bool
	prefix         string
	has_prefix     bool
	indexName      string
	has_indexName  bool
	editName       string
	has_editName   bool
}

func (c *createHandlerOptionImpl) IndexTitle() string  { return c.indexTitle }
func (c *createHandlerOptionImpl) HasIndexTitle() bool { return c.has_indexTitle }
func (c *createHandlerOptionImpl) Prefix() string      { return c.prefix }
func (c *createHandlerOptionImpl) HasPrefix() bool     { return c.has_prefix }
func (c *createHandlerOptionImpl) IndexName() string   { return c.indexName }
func (c *createHandlerOptionImpl) HasIndexName() bool  { return c.has_indexName }
func (c *createHandlerOptionImpl) EditName() string    { return c.editName }
func (c *createHandlerOptionImpl) HasEditName() bool   { return c.has_editName }

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
