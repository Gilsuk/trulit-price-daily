package handler

var Factory interface {
	New() Handler
} = factory{}

type factory struct {
}

func (f factory) New() Handler {
	return infoHandler{}
}
