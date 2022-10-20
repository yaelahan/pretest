package validator

type Error struct {
	Message string
	Errors  interface{}
}

func (err Error) Error() string {
	return err.Message
}
