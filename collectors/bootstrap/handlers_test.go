package main

import (
	"testing"

	bootstrap "github.com/gilsuk/trulit-price-daily/collectors/bootstrap/mocks"
)

func TestHandlerCallService(t *testing.T) {
	// given
	collector := bootstrap.NewCollector(t)
	handler := handlerConstructor{collector}.GetHandler()

	// when, then
	collector.EXPECT().Collect().Return().Once()
	handler()
}
