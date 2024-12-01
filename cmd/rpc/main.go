package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/google/uuid"
	"net/rpc"
	"time"
)

func main() {
	// This is a placeholder for the main function.
	// Connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:3001")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	// Define the arguments for the Add method
	s3EventRecord := events.S3EventRecord{
		EventVersion:      "Test",
		EventSource:       "Cirrus",
		AWSRegion:         "ca-central-1",
		EventTime:         time.Now(),
		EventName:         "cirrus-test-event",
		PrincipalID:       events.S3UserIdentity{},
		RequestParameters: events.S3RequestParameters{},
		ResponseElements:  nil,
		S3:                events.S3Entity{},
	}

	s3Payload := events.S3Event{
		Records: []events.S3EventRecord{s3EventRecord},
	}

	requestPayload, err := json.Marshal(s3Payload)
	if err != nil {
		return
	}

	invokeRequest := messages.InvokeRequest{
		Payload:               requestPayload,
		RequestId:             uuid.NewString(),
		XAmznTraceId:          "",
		Deadline:              messages.InvokeRequest_Timestamp{},
		InvokedFunctionArn:    "arn:aws:lambda:ca-central-1:123456789012:function:my-function",
		CognitoIdentityId:     "1234",
		CognitoIdentityPoolId: "1234",
	}

	var invokeResponse messages.InvokeResponse

	// Call the Add method on the server
	err = client.Call("Function.Invoke", invokeRequest, &invokeResponse)
	if err != nil {
		fmt.Println("Error Invoking Lambda:", err)
		return
	}

	jsonResponse, _ := json.Marshal(invokeResponse)

	// Print the result
	fmt.Printf("result %s", jsonResponse)
}
