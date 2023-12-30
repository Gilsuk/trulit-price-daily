package handler

var Factory Ifactory = handlerFactory{}

type Handler interface {
	Handle()
}

type Ifactory interface {
	New() Handler
}

type dummyHandler struct {
}

func (dummyHandler) Handle() {

}

type handlerFactory struct {
}

func (f handlerFactory) New() Handler {
	return dummyHandler{}
}
