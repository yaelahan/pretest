package validator

type Error struct {
	Message string
	Errors  interface{}
}

func (err Error) Error() string {
	return err.Message
}

func NewValidationError(errors interface{}) {
	panic(Error{
		Message: "The given data was invalid.",
		Errors:  errors,
	})
}
