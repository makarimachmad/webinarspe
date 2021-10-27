package main

import (
	"spewebinar/router"
	"spewebinar/database"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()

	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("berhasil terhubung fiber")
	})

	router.SetupRoutes(app)

	app.Listen(":3000")
}
