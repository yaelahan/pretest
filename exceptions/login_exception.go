package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"pretest-indihomesmart/internal/validator"
)

func NewLoginException() {
	panic(validator.Error{
		Message: "Login Failed",
		Errors: fiber.Map{
			"email": "These credentials do not match our records.",
		},
	})
}
