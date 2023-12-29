package handlers_test

import (
	"testing"

	"github.com/gilsuk/trulit-price-daily/collector/presenter/handlers"
	"github.com/gilsuk/trulit-price-daily/collector/worker/collectors/mocks"
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
	handler := handlers.HandlerConstructor{collector}.GetHandler()

	// when, then
	collector.EXPECT().Collect().Return().Once()
	handler()
}
