// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

// ExampleResponse defines model for ExampleResponse.
type ExampleResponse struct {
	MessageId *string                 `json:"message_id,omitempty"`
	Payload   *map[string]interface{} `json:"payload,omitempty"`
}