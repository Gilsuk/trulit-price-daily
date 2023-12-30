package handler

import "github.com/gilsuk/trulit-price-daily/collector/worker/collector"

type Handler interface {
	Handle()
}

func New(c collector.Collector) Handler {
	return infoHandler{
		collector: c,
	}
}

type infoHandler struct {
	collector collector.Collector
}

func (h infoHandler) Handle() {
	h.collector.Collect()
}
