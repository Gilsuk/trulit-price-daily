package main

import (
	"testing"

	"github.com/gilsuk/trulit-price-daily/collectors/bootstrap/mocks"
)

func TestHandlerCallService(t *testing.T) {
	// given
	collector := mocks.NewCollector(t)
	handler := handlerConstructor{collector}.GetHandler()

	// when, then
	collector.EXPECT().Collect().Return().Once()
	handler()
}
