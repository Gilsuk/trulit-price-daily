package handler_test

import (
	"testing"

	"github.com/gilsuk/trulit-price-daily/collector/presenter/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) TestIfHandlerCallService() {
	// // given
	// collector := mocks.NewCollector(suite.T())
	// handler := handler.New(collector)

	// // when, then
	// collector.EXPECT().Collect().Return().Once()
	// handler()
}

func (suite *HandlerTestSuite) TestIfHandlerFactoryReturnHandler() {
	// given
	handlerInstance := handler.Factory.New()

	// when, then
	assert.NotNil(suite.T(), handlerInstance)
	assert.Implements(suite.T(), (*handler.Handler)(nil), handlerInstance)
}
