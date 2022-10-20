package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"pretest-indihomesmart/controllers/auth"
	"pretest-indihomesmart/internal/validator"
	"pretest-indihomesmart/middlewares"
	"pretest-indihomesmart/services"
)

func NewRouter(app *fiber.App, db *gorm.DB, validator *validator.Validator) {
	userService := services.NewUserService(db)
	jwtService := services.NewJwtService()

	registerController := auth.NewRegisterController(userService, validator)
	loginController := auth.NewLoginController(userService, jwtService, validator)
	profileController := auth.NewProfileController(userService, validator)

	api := app.Group("/api")
	api.Post("/register", registerController.Register)
	api.Post("/login", loginController.Login)

	authenticated := api.Use(middlewares.NewAuthMiddleware(jwtService))
	authenticated.Get("/logout", loginController.Logout)
	authenticated.Get("/profile", profileController.GetProfile)
}
