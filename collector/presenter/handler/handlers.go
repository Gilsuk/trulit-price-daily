package handler

var Factory = factory{}

type Handler interface {
	Handle()
}

type factory struct {
}

type dummyHandler struct {
}

func (dummyHandler) Handle() {

}

func (f factory) New() Handler {
	return dummyHandler{}
}
