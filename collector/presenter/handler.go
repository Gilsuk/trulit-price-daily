package presenter

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gilsuk/trulit-price-daily/collector/worker"
)

type AWSLambdaHandler func(*events.SQSEvent) error

func New(w worker.Worker) AWSLambdaHandler {
	return func(s *events.SQSEvent) error {
		request := &worker.Request{}
		err := json.Unmarshal([]byte(s.Records[0].Body), request)
		if err != nil {
			return errors.Join(ErrInvalidJsonSyntax, err)
		}
		_, err = time.Parse("2006-01-02", request.Date)
		if err != nil {
			return errors.Join(ErrInvalidDateFormat, err)
		}
		w.Do(*request)
		return nil
	}
}
