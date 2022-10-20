package utils

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func ErrorResponse(message string, errors interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Errors:  errors,
	}
}
