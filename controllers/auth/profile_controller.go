package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"pretest-indihomesmart/internal/validator"
	"pretest-indihomesmart/services"
	"pretest-indihomesmart/utils"
)

type ProfileController struct {
	userService services.UserService
	validator   *validator.Validator
}

func NewProfileController(userService services.UserService, validator *validator.Validator) ProfileController {
	return ProfileController{
		userService: userService,
		validator:   validator,
	}
}

func (c *ProfileController) GetProfile(ctx *fiber.Ctx) error {
	token := ctx.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	user := c.userService.FindByEmail(claims["email"].(string))

	resp := utils.SuccessResponse("", fiber.Map{
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
	})
	return ctx.JSON(resp)
}
