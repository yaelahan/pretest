package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"pretest-indihomesmart/services"
	"pretest-indihomesmart/utils"
	"strings"
)

func NewAuthMiddleware(jwtService services.JwtService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var token *jwt.Token

		authorization := c.GetReqHeaders()["Authorization"]
		tokenString := strings.Split(authorization, "Bearer ")

		if len(tokenString) < 2 {
			err = fmt.Errorf("invalid authorization headers")
		} else {
			token, err = jwtService.Verify(tokenString[1])
			if err == nil && token.Valid {
				// Store user information from token into context.
				c.Locals("user", token)
				return c.Next()
			}
		}

		resp := utils.ErrorResponse(err.Error(), nil)
		return c.Status(401).JSON(resp)
	}
}
