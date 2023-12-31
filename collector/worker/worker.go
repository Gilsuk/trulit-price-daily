package worker

type Worker interface {
	Do(Request)
}

type InfoWorker struct {
}

func (w InfoWorker) Do(r Request) {

}
