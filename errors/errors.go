package errors

import (
	"fmt"

	e "github.com/rockspoon/rs.cor.common-model/errors"
)

const (
	errorPrefix                 = 4723
	invalidMiddleware           = 1
	databaseError               = 2
	validationError             = 3
	decodeBody                  = 4
	invalidMicroserviceResponse = 5
)

// DecodeBody returns decode body error by calling the middleware.
func DecodeBody(err error) error {
	return e.NewError(errorPrefix, decodeBody, fmt.Errorf("could not parse request body: %v", err))
}

// ValidateStruct returns validate structure error.
func ValidateStruct(err error) error {
	return e.NewError(errorPrefix, validationError, fmt.Errorf("validation error: %v", err))
}

// DatabaseError returns database error
func DatabaseError(cause error) error {
	return e.NewError(errorPrefix, databaseError, cause)
}

// InvalidMiddlewareContext invalid unmarshal of middleware context
func InvalidMiddlewareContext() error {
	return e.NewError(errorPrefix, invalidMiddleware, fmt.Errorf("was impossible to unmarshal the middleware context"))
}

// InvalidMicroserviceResponse returns an error when a microservice responds with error.
func InvalidMicroserviceResponse(msName, status string) error {
	return e.NewError(errorPrefix, invalidMicroserviceResponse, fmt.Errorf("microservice "+msName+" responded with "+status))
}
