package main

import (
	"merchant-portal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.MerchantPortalRoute(app)

	app.Listen(":3000")
}
