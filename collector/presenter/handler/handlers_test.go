package handler_test

import (
	"testing"

	"github.com/gilsuk/trulit-price-daily/collector/presenter/handler"
	"github.com/gilsuk/trulit-price-daily/collector/worker/collector/mocks"
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
	collector := mocks.NewCollector(suite.T())
	handler := handler.New(collector)

	// when, then
	collector.EXPECT().Collect().Return().Once()
	handler()
}
