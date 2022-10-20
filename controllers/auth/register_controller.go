package auth

import (
	"github.com/gofiber/fiber/v2"
	"pretest-indihomesmart/entities"
	"pretest-indihomesmart/exceptions"
	"pretest-indihomesmart/internal/validator"
	"pretest-indihomesmart/services"
	"pretest-indihomesmart/utils"
)

type RegisterController struct {
	userService services.UserService
	validator   *validator.Validator
}

func NewRegisterController(userService services.UserService, validator *validator.Validator) RegisterController {
	return RegisterController{
		userService: userService,
		validator:   validator,
	}
}

func (c *RegisterController) Register(ctx *fiber.Ctx) error {
	request := entities.RegisterRequest{}
	err := ctx.BodyParser(&request)
	exceptions.PanicIfNeeded(err)

	if validationError := c.validator.Validate(request); validationError != nil {
		validator.NewValidationError(validationError)
	}

	user := c.userService.Create(request)

	return ctx.JSON(utils.SuccessResponse(user))
}
