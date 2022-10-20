package exceptions

import (
	"pretest-indihomesmart/internal/validator"
)

func NewValidationException(errors interface{}) {
	panic(validator.Error{
		Message: "The given data was invalid.",
		Errors:  errors,
	})
}
