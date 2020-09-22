package main

import (
	"github.com/gofiber/fiber"
	fiberprometheus "github.com/hepsiburada/fiber-prometheus"
)

func main() {
	app := fiber.New()

	p8sMiddleware := fiberprometheus.NewMiddleware("fiber", "http", "/metrics")
	p8sMiddleware.Register(app)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, Go World!")
	})

	app.Get("/testurl", func(c *fiber.Ctx) {
		c.Send("this is testurl.")
	})

	app.Get("/:name", func(c *fiber.Ctx) {
		c.Send("Hello " + c.Params("name"))
	})

	app.Listen(7000)
}
