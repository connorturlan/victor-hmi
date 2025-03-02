package main

import (
	"fuelpriceservice/cmd"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(cmd.Handler)
}

//{"CollectionMethod":{"S":"T"},"FuelId":{"N":"2"},"Price":{"N":"2799"},"SiteId":{"N":"61577372"},"TransactionDateUtc":{"S":"2023-10-27T05:11:11.663"}}
