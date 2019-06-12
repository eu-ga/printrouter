package errors

import (
	"github.com/rockspoon/go-common/errors"
)

const (
	// errCodePrefix represents code prefix error.
	errCodePrefix = "device-ms"

	codeValidateStruct = "ValidateStructure"
)

func errCodeWithPrefix(code string) string {
	return errCodePrefix + code
}

func makeError(code, details string, cause error, statusSetter errors.ErrorStatusSetter) error {
	return statusSetter(errCodeWithPrefix(code), details, cause)
}

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
