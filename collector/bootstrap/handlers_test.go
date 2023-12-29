package main

import (
	"testing"

	"github.com/gilsuk/trulit-price-daily/collectors/bootstrap/mocks"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (suite *HandlerTestSuite) HandlerShouldCallService() {
	// given
	collector := mocks.NewCollector(suite.T())
	handler := handlerConstructor{collector}.GetHandler()

	// when, then
	collector.EXPECT().Collect().Return().Once()
	handler()
}
