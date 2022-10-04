package handler

type handlerBuilder struct {
	handlers []Handler
}

type HandlerBuilder interface {
	NewHandler(name string, fn HandlerFn, optss ...NewHandlerOption)
	NewHandlerFromParams(name string, hf handlerFn, pCtor ctorFn, optss ...NewHandlerOption)
	Build() []Handler
}

func NewHandlerBuilder() HandlerBuilder {
	return &handlerBuilder{}
}

func (h *handlerBuilder) NewHandlerFromParams(name string, hf handlerFn, pCtor ctorFn, optss ...NewHandlerOption) {
	h.handlers = append(h.handlers, NewHandlerFromParams(name, hf, pCtor, optss...))
}

func (h *handlerBuilder) NewHandler(name string, fn HandlerFn, optss ...NewHandlerOption) {
	h.handlers = append(h.handlers, NewHandler(name, fn, optss...))
}

func (h *handlerBuilder) Build() []Handler {
	return h.handlers
}
