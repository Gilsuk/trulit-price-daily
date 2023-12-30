package collector

type Collector interface {
	Collect(Request)
}

type InfoCollector struct {
}

func (c InfoCollector) Collect(r Request) {

}
