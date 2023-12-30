package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gilsuk/trulit-price-daily/collector/presenter/handler"
)

func main() {
	lambda.Start(handler.Factory.New())
}
