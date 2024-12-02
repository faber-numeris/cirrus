package lambdas

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda/messages"
	"github.com/google/uuid"
	"net/rpc"
	"time"
)

// Lambda is a struct that holds the metadata of a lambda function
type Lambda struct {
	FunctionName string
	Address      string
	Port         string
}

// LambdaRegistrar is an interface that defines the methods to register and invoke lambda functions
type LambdaRegistrar interface {
	Add(functionName string, address string, port string)
	Invoke(functionName string, payload any)
}

// LambdaRegistry is a struct that holds the lambda functions registered in the system
type LambdaRegistry struct {
	// Lambdas is a map that holds the lambda functions registered in the system. The key is the function name and the value is the Lambda struct
	Lambdas map[string]Lambda
}

// Invoke is a method that invokes a lambda function
func (l LambdaRegistry) Invoke(functionName string, payload any) (any, error) {

	// GET the lambda function from the registry
	lambda, ok := l.Lambdas[functionName]
	if !ok {
		return nil, fmt.Errorf("%w %s", ErrFunctionNotFound, functionName)
	}

	lambdaAddress := fmt.Sprintf("%s:%s", lambda.Address, lambda.Port)
	client, err := rpc.Dial("tcp", lambdaAddress)
	if err != nil {
		return nil, fmt.Errorf(
			"%w for function %s : %s",
			ErrServerConnection,
			functionName,
			lambdaAddress,
		)
	}
	defer func(client *rpc.Client) { _ = client.Close() }(client)

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
		return nil, fmt.Errorf("%w : %s", ErrMarshallingPayload, err)
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
