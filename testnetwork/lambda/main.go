package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"log/slog"
	"os"
)

const (
	EnvVarEventType = "EVENT_TYPE"
	EventTypeS3     = "S3"
	EventTypeSQS    = "SQS"
)

func main() {
	eventType := os.Getenv(EnvVarEventType)
	switch eventType {
	case EventTypeS3:
		slog.Info("Starting S3 Lambda")
		lambda.Start(HandleS3)
	case EventTypeSQS:
		slog.Info("Starting SQS Lambda")
		lambda.Start(HandleSQS)
	}
}

func HandleS3(_ context.Context, event events.S3Event) (string, error) {
	for _, record := range event.Records {
		s3 := record.S3
		log.Printf("Bucket: %s, Key: %s", s3.Bucket.Name, s3.Object.Key)
	}
	return "Processed S3 Event", nil
}

func HandleSQS(_ context.Context, event events.SQSEvent) (string, error) {
	for _, message := range event.Records {
		log.Printf("Message ID: %s, Body: %s", message.MessageId, message.Body)
	}
	return "Processed SQS Event", nil
}
