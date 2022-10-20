package handler

type handlerBuilder struct {
	handlers []Handler
}

type HandlerBuilder interface {
	NewHandlerFromHandlerFn(name string, fn HandlerFn, optss ...NewHandlerOption)
	NewHandler(name string, hf handlerFn, paramsPrototype any, optss ...NewHandlerOption)
	NewStaticHandler(name string, res []byte, paramsPrototype any, optss ...NewHandlerOption)
	Build() []Handler
}

func NewHandlerBuilder() HandlerBuilder {
	return &handlerBuilder{}
}

func (h *handlerBuilder) NewHandler(name string, hf handlerFn, paramsPrototype any, optss ...NewHandlerOption) {
	h.handlers = append(h.handlers, NewHandler(name, hf, paramsPrototype, optss...))
}

func (h *handlerBuilder) NewHandlerFromHandlerFn(name string, fn HandlerFn, optss ...NewHandlerOption) {
	h.handlers = append(h.handlers, NewHandlerFromHandlerFn(name, fn, optss...))
}

func (h *handlerBuilder) NewStaticHandler(name string, htmlBytes []byte, paramsPrototype any, optss ...NewHandlerOption) {
	h.handlers = append(h.handlers, NewStaticHandler(name, htmlBytes, paramsPrototype, optss...))
}

func (h *handlerBuilder) Build() []Handler {
	return append([]Handler{}, h.handlers...)
}
