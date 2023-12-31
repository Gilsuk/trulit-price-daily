package presenter

import "github.com/gilsuk/trulit-price-daily/collector/worker"

func NewAWSLambdaHandler() AWSLambdaHandler {
	return New(worker.InfoWorker{})
}
