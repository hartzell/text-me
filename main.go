package main

// credentials file, see https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
// AWS_REGION=us-west-2

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/tj/go-sms"
)

func main() {

	// rustle up some command line arguments
	var message string
	flag.StringVar(&message, "message", "empty message", "the message to send")
	var number string
	flag.StringVar(&number, "number", "+15108820056", "recipient phone number")
	var maxPrice float64
	flag.Float64Var(&maxPrice, "maxprice", 0.5, "recipient phone number")

	flag.Parse()

	sms.DefaultMaxPrice = maxPrice
	sms.DefaultType = sms.Transactional

	err := Send(message, number)
	if err != nil {
		fmt.Print("Error was: ", err)
		os.Exit(1)
	}
}

// Send `message` to `number` using defaults.
func Send(message, number string) error {
	service := sns.New(session.New(aws.NewConfig().WithRegion("us-west-2")))
	sms := sms.SMS{Service: service}
	return sms.Send(message, number)
}
