package handler

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/gilsuk/trulit-price-daily/collector/worker/collector"
)

type Handler func(*events.SQSEvent) error

func New(c collector.Collector) Handler {
	return func(s *events.SQSEvent) error {
		c.Collect()
		return nil
	}
}
