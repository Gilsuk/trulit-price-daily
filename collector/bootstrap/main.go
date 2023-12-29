package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gilsuk/trulit-price-daily/collector/presenter/handlers"
)

func main() {
	lambda.Start(handlers.HandlerConstructor{})
}
