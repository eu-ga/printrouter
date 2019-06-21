package errors

import (
	"github.com/rockspoon/go-common/errors"
)

const (
	codeValidateStruct = "ValidateStructure"
)

// DecodeBody returns decode body error
func DecodeBody(cause error) error {
	return errors.DecodeBody(cause)
}

// ValidateStruct returns validate structure error.
func ValidateStruct(err error) error {
	return errors.BadRequest(codeValidateStruct, "validation error", err)
}

// DatabaseError returns database error
func DatabaseError(details string, cause error) error {
	return errors.DatabaseError(details, cause)
}
