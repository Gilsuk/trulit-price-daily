package presenter

import "github.com/gilsuk/trulit-price-daily/collector/worker/collector"

func NewAWSLambdaHandler() AWSLambdaHandler {
	return New(collector.InfoCollector{})
}
