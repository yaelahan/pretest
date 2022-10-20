package exceptions

type ValidationError struct {
	Message string
	Errors  interface{}
}

func (err ValidationError) Error() string {
	return err.Message
}

func NewValidationError(errors interface{}) {
	panic(ValidationError{
		Message: "The given data was invalid.",
		Errors:  errors,
	})
}
