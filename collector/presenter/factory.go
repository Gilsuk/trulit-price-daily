package presenter

import "github.com/gilsuk/trulit-price-daily/collector/worker/collector"

var Factory interface {
	New() Handler
} = factory{}

type factory struct {
}

func (f factory) New() Handler {
	return New(collector.InfoCollector{})
}
