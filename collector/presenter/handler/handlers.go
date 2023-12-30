package handler

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gilsuk/trulit-price-daily/collector/worker/collector"
)

type Handler func(*events.SQSEvent) error

func New(c collector.Collector) Handler {
	return func(s *events.SQSEvent) error {
		request := &collector.Request{}
		err := json.Unmarshal([]byte(s.Records[0].Body), request)
		if err != nil {
			return errors.Join(ErrInvalidDateFormat, err)
		}
		c.Collect(*request)
		return ErrInvalidDateFormat
	}
}
