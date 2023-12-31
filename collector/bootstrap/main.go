package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gilsuk/trulit-price-daily/collector/presenter"
)

func main() {
	lambda.Start(presenter.NewAWSLambdaHandler())
}
