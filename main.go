package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"pretest-indihomesmart/exceptions"
	"pretest-indihomesmart/handlers"
	"pretest-indihomesmart/internal/database"
	internalValidator "pretest-indihomesmart/internal/validator"
	"pretest-indihomesmart/routes"
)

var err error

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// setup fiber
	db := database.New()
	app := fiber.New(fiber.Config{
		ErrorHandler: handlers.FiberErrorHandler,
	})
	app.Use(recover.New())

	// setup validator
	validate := validator.New()
	customValidator := internalValidator.New(validate, db)

	// register routes
	routes.NewRouter(app, db, customValidator)

	// start routes
	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	err = app.Listen(port)
	exceptions.PanicIfNeeded(err)
}
