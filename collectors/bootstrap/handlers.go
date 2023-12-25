package main

type handler func()

type handlerConstructor struct {
	collector Collector
}

func (c handlerConstructor) GetHandler() handler {
	return func() {
		c.collector.Collect()
	}
}
