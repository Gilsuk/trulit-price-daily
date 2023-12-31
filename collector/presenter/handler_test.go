package presenter_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/gilsuk/trulit-price-daily/collector/presenter"
	"github.com/gilsuk/trulit-price-daily/collector/worker"
	"github.com/gilsuk/trulit-price-daily/collector/worker/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) TestIfHandlerCallService() {
	// given
	workerMock := mocks.NewWorker(suite.T())
	handler := presenter.New(workerMock)
	request := worker.Request{
		Date: "2023-12-30",
	}
	input := marshal([]worker.Request{request})

	// when, then
	workerMock.EXPECT().Do(mock.Anything).Return().Once()
	handler(&input)
}

func (suite *HandlerTestSuite) TestIfHandlerFactoryReturnHandler() {
	// given
	handlerInstance := presenter.NewAWSLambdaHandler()

	// when, then
	assert.NotNil(suite.T(), handlerInstance)
	assert.IsType(suite.T(), (*presenter.AWSLambdaHandler)(nil), &handlerInstance)
}

func (suite *HandlerTestSuite) TestHandlerCanConvertInput() {
	// given
	workerMock := mocks.NewWorker(suite.T())
	handler := presenter.New(workerMock)
	request := worker.Request{
		Date: "2023-12-30",
	}
	input := marshal([]worker.Request{request})

	// when, then
	workerMock.EXPECT().Do(request).Return().Once()
	handler(&input)
}

func (suite *HandlerTestSuite) TestHandlerReturnErrorWhenWrongInputReceived() {
	cases := []struct {
		input  events.SQSEvent
		expect error
	}{{
		input:  marshal([]worker.Request{{Date: "2023-12-30"}}),
		expect: nil,
	}, {
		input:  marshal([]worker.Request{{Date: "2023-13-60"}}),
		expect: presenter.ErrInvalidDateFormat,
	}, {
		input:  marshal([]worker.Request{{Date: ""}}),
		expect: presenter.ErrInvalidDateFormat,
	}, {
		input:  events.SQSEvent{Records: []events.SQSMessage{{Body: ""}}},
		expect: presenter.ErrInvalidJsonSyntax,
	}}

	for _, tc := range cases {
		// given
		handler := presenter.NewAWSLambdaHandler()
		// when
		err := handler(&tc.input)
		// then
		assert.ErrorIs(suite.T(), err, tc.expect)
	}
}

// TODO. when does SQSEvent contains multiple SQSMessage

func marshal(rs []worker.Request) events.SQSEvent {
	res := make([]events.SQSMessage, len(rs))
	for i, v := range rs {
		body, _ := json.Marshal(v)
		res[i].Body = string(body)
	}
	return events.SQSEvent{
		Records: res,
	}
}
