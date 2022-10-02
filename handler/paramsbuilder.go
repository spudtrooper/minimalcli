package handler

// Type paramsBuilder accumulates and builds a HandlerMetadataParam.
type paramsBuilder struct {
	params []HandlerMetadataParam
}

// BuildOption builds and returns the HandlerMetadata.
func (p *paramsBuilder) Build() HandlerMetadata {
	return HandlerMetadata{
		Params: p.params,
	}
}

// BuildOption builds and returns the NewHandlerMetadata option.
func (p *paramsBuilder) BuildOption() NewHandlerOption {
	return NewHandlerMetadata(p.Build())
}

// Params creates a builder for HandlerMetadataParams.
func Params() *paramsBuilder {
	return &paramsBuilder{}
}

// String creates a new optional `string` param
func (p *paramsBuilder) String(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeString,
	})
	return p
}

// String creates a new required `string` param
func (p *paramsBuilder) RequiredString(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name:     name,
		Type:     HandlerMetadataParamTypeString,
		Required: true,
	})
	return p
}

// String creates a new optional `bool` param
func (p *paramsBuilder) Bool(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeBool,
	})
	return p
}

// String creates a new optional `int` param
func (p *paramsBuilder) Int(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeInt,
	})
	return p
}

// String creates a new optional `time.Duration` param
func (p *paramsBuilder) Duration(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeDuration,
	})
	return p
}
