package collector

type Collector interface {
	Collect(Request)
}
