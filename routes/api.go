package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"os"
	"pretest-indihomesmart/controllers/auth"
	"pretest-indihomesmart/handlers"
	"pretest-indihomesmart/internal/validator"
	"pretest-indihomesmart/services"
)

func NewRouter(app *fiber.App, db *gorm.DB, validator *validator.Validator) {
	userService := services.NewUserService(db)
	registerController := auth.NewRegisterController(userService, validator)
	loginController := auth.NewLoginController(userService, validator)
	profileController := auth.NewProfileController(userService, validator)

	api := app.Group("/api")
	api.Post("/register", registerController.Register)
	api.Post("/login", loginController.Login)

	authenticated := api.Use(jwtware.New(jwtware.Config{
		SigningMethod: jwt.SigningMethodHS256.Name,
		SigningKey:    []byte(os.Getenv("APP_KEY")),
		ErrorHandler:  handlers.JwtErrorHandler,
	}))
	authenticated.Get("/profile", profileController.GetProfile)
}
