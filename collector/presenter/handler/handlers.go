package handler

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gilsuk/trulit-price-daily/collector/worker/collector"
)

type Handler func(*events.SQSEvent) error

func New(c collector.Collector) Handler {
	return func(s *events.SQSEvent) error {
		request := &collector.Request{}
		err := json.Unmarshal([]byte(s.Records[0].Body), request)
		if err != nil {
			return errors.Join(ErrInvalidJsonSyntax, err)
		}
		_, err = time.Parse("2006-01-02", request.Date)
		if err != nil {
			return errors.Join(ErrInvalidDateFormat, err)
		}
		c.Collect(*request)
		return nil
	}
}
