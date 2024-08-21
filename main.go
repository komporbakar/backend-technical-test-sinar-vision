package main

import (
	"backend_technical_test/database"
	"backend_technical_test/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.Database()
	database.DBMigrate()

	//Router
	router.NewRouter(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
