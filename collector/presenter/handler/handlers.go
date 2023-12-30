package handler

import "github.com/gilsuk/trulit-price-daily/collector/worker/collector"

type handler func()

func New(c collector.Collector) handler {
	return func() {
		c.Collect()
	}
}
