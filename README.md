# Fiber Prometheus
Middleware for go fiber v1.x

## Installation
```bash
go get github.com/hepsiburada/fiber-prometheus
```

## Examples

```go
package main

import (
	"github.com/gofiber/fiber"
	"github.com/hepsiburada/fiber-prometheus"
)

func main() {
	app := fiber.New()

	p8sMiddleware := fiberprometheus.NewMiddleware("fiber","http","/metrics")
	p8sMiddleware.Register(app)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, Go World!")
    })
    
	app.Get("/testurl", func(c *fiber.Ctx) {
		c.Send("this is testurl.")
    })
    
	app.Listen(7000)
}
```