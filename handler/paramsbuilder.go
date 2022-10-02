package handler

type paramsBuilder struct {
	params []HandlerMetadataParam
}

func (p *paramsBuilder) Build() HandlerMetadata {
	return HandlerMetadata{
		Params: p.params,
	}
}

func Params() *paramsBuilder { return &paramsBuilder{} }

func (p *paramsBuilder) String(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeString,
	})
	return p
}

func (p *paramsBuilder) RequiredString(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name:     name,
		Type:     HandlerMetadataParamTypeString,
		Required: true,
	})
	return p
}

func (p *paramsBuilder) Bool(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeBool,
	})
	return p
}

func (p *paramsBuilder) Int(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeInt,
	})
	return p
}

func (p *paramsBuilder) Duration(name string) *paramsBuilder {
	p.params = append(p.params, HandlerMetadataParam{
		Name: name,
		Type: HandlerMetadataParamTypeDuration,
	})
	return p
}
