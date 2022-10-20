package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"pretest-indihomesmart/utils"
)

// ErrorHandler fiber
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var statusCode int
	var errors interface{}

	switch err.(type) {
	case ValidationError:
		statusCode = 400
		errors = err.(ValidationError).Errors

	case *fiber.Error:
		statusCode = err.(*fiber.Error).Code

	default:
		statusCode = 500
	}

	return ctx.Status(statusCode).
		JSON(utils.ErrorResponse(err.Error(), errors))
}
