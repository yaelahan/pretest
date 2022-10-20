package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"pretest-indihomesmart/exceptions"
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
		ErrorHandler: exceptions.ErrorHandler,
	})
	app.Use(recover.New())

	// setup validator
	validate := validator.New()
	customValidator := internalValidator.New(validate, db)

	// register routes
	routes.NewRouter(app, db, customValidator)

	// start routes
	err = app.Listen(":1337")
	exceptions.PanicIfNeeded(err)
}
