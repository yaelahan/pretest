package handlers

import (
	"github.com/gofiber/fiber/v2"
	"pretest-indihomesmart/utils"
)

func JwtErrorHandler(ctx *fiber.Ctx, err error) error {
	resp := utils.ErrorResponse(err.Error(), nil)
	return ctx.Status(401).JSON(resp)
}
