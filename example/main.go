package main

import (
	"github.com/gofiber/fiber/v2"
	fiberprometheus "github.com/hepsiburada/fiber-prometheus"
)

func main() {
	app := fiber.New()

	p8sMiddleware := fiberprometheus.NewMiddleware("fiber", "http", "/metrics")
	p8sMiddleware.Register(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Go World!")
	})

	app.Get("/testurl", func(c *fiber.Ctx) error {
		return c.SendString("this is testurl.")
	})

	app.Get("/:name", func(c *fiber.Ctx) error {
		return c.SendString("Hello " + c.Params("name"))
	})

	app.Listen(":7000")
}
