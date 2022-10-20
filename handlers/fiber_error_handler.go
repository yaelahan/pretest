package handlers

import (
	"github.com/gofiber/fiber/v2"
	"pretest-indihomesmart/internal/validator"
	"pretest-indihomesmart/utils"
)

func FiberErrorHandler(ctx *fiber.Ctx, err error) error {
	var statusCode int
	var errors interface{}

	switch err.(type) {
	case validator.Error:
		statusCode = 400
		errors = err.(validator.Error).Errors

	case *fiber.Error:
		statusCode = err.(*fiber.Error).Code

	default:
		statusCode = 500
	}

	data := utils.ErrorResponse(err.Error(), errors)
	return ctx.Status(statusCode).JSON(data)
}
