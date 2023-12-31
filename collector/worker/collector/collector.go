package collector

type Worker interface {
	Do(Request)
}

type InfoCollector struct {
}

func (c InfoCollector) Do(r Request) {

}
