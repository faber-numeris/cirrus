package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
)

func main() {
	err := os.Setenv("_LAMBDA_SERVER_PORT", "3001")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("AWS_LAMBDA_RUNTIME_API", "localhost")
	if err != nil {
		panic(err)
	}

	lambda.Start(Handle)
}

func Handle(_ context.Context, event events.S3Event) (string, error) {
	for _, record := range event.Records {
		s3 := record.S3
		log.Printf("Bucket: %s, Key: %s", s3.Bucket.Name, s3.Object.Key)
	}
	return "Processed S3 Event", nil
}
