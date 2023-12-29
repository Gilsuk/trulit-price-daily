package handlers

import "github.com/gilsuk/trulit-price-daily/collector/worker/collectors"

type Handler func()

type HandlerConstructor struct {
	Collector collectors.Collector
}

func (c HandlerConstructor) GetHandler() Handler {
	return func() {
		c.Collector.Collect()
	}
}
