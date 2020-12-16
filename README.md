# Fiber Prometheus
Middleware for go fiber v2.x

## Installation
```bash
go get github.com/hepsiburada/fiber-prometheus/v2
```

## Examples

```go
package main

import (
	"github.com/gofiber/fiber/v2"
	fiberprometheus "github.com/hepsiburada/fiber-prometheus/v2"
)

func main() {
	app := fiber.New()

	p8sMiddleware := fiberprometheus.NewMiddleware("fiber","http","/metrics")
	p8sMiddleware.Register(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Go World!")
	})

	app.Get("/testurl", func(c *fiber.Ctx) error {
		return c.SendString("this is testurl.")
	})

	app.Listen(":7000")
}
```