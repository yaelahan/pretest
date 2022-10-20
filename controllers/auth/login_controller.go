package auth

import (
	"github.com/gofiber/fiber/v2"
	fiberUtils "github.com/gofiber/fiber/v2/utils"
	"pretest-indihomesmart/entities"
	"pretest-indihomesmart/exceptions"
	"pretest-indihomesmart/internal/validator"
	"pretest-indihomesmart/models"
	"pretest-indihomesmart/services"
	"pretest-indihomesmart/utils"
	"strings"
)

type LoginController struct {
	userService services.UserService
	jwtService  services.JwtService
	validator   *validator.Validator
}

func NewLoginController(
	userService services.UserService,
	jwtService services.JwtService,
	validator *validator.Validator,
) LoginController {
	return LoginController{
		userService: userService,
		jwtService:  jwtService,
		validator:   validator,
	}
}

func (c *LoginController) Login(ctx *fiber.Ctx) error {
	request := entities.LoginRequest{}
	err := ctx.BodyParser(&request)
	exceptions.PanicIfNeeded(err)

	if validationError := c.validator.Validate(request); validationError != nil {
		exceptions.NewValidationException(validationError)
	}

	user := c.userService.FindByEmail(request.Email)
	if user == (models.User{}) {
		exceptions.NewLoginException()
	}

	if err := utils.ComparePassword(user.Password, request.Password); err != nil {
		exceptions.NewLoginException()
	}

	payload := fiber.Map{
		"id":    fiberUtils.UUIDv4(),
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
	}
	token, err := c.jwtService.Sign(payload, 60)
	resp := entities.LoginResponse{Token: token}
	return ctx.JSON(utils.SuccessResponse("Success", resp))
}

func (c *LoginController) Logout(ctx *fiber.Ctx) error {
	authorization := ctx.GetReqHeaders()["Authorization"]
	tokenString := strings.Split(authorization, "Bearer ")[1]

	err := c.jwtService.Revoke(tokenString)
	exceptions.PanicIfNeeded(err)

	return ctx.JSON(utils.SuccessResponse("Success", nil))
}
