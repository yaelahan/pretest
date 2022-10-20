package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pretest-indihomesmart/controllers/auth"
	"pretest-indihomesmart/services"
	"pretest-indihomesmart/utils"
)

func NewRouter(app *fiber.App, db *gorm.DB, validator *utils.CustomValidator) {
	userService := services.NewUserService(db)
	registerController := auth.NewRegisterController(userService, validator)

	app.Post("/register", registerController.Register)
}
