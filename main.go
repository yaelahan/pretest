package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"pretest-indihomesmart/config"
	"pretest-indihomesmart/exceptions"
	"pretest-indihomesmart/routes"
	"pretest-indihomesmart/utils"
)

var err error

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// setup fiber
	db := config.NewDB()
	app := fiber.New(fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	})
	app.Use(recover.New())

	// setup validator
	validate := validator.New()
	customValidator := utils.NewCustomValidator(validate)

	// register routes
	routes.NewRouter(app, db, customValidator)

	// start routes
	err = app.Listen(":1337")
	exceptions.PanicIfNeeded(err)
}
