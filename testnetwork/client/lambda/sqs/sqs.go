package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func main() {
	//err := os.Setenv("_LAMBDA_SERVER_PORT", "3001")
	//if err != nil {
	//	panic(err)
	//}
	//err = os.Setenv("AWS_LAMBDA_RUNTIME_API", "localhost")
	//if err != nil {
	//	panic(err)
	//}

	lambda.Start(Handle)
}

func Handle(_ context.Context, event events.SQSEvent) (string, error) {
	for _, message := range event.Records {
		log.Printf("Message ID: %s, Body: %s", message.MessageId, message.Body)
	}
	return "Processed SQS Event", nil
}
