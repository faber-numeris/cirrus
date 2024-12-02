package lambdas

import (
	"errors"
)

var (
	ErrFunctionNotFound   = errors.New("function not found")
	ErrServerConnection   = errors.New("error connecting to server")
	ErrMarshallingPayload = errors.New("error marshalling payload")
)
