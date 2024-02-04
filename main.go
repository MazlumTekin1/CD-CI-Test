package main

import (
	"github.com/gofiber/fiber/v2"
)

// func handler(c *fiber.Ctx) {
// 	result := c.Params("foo")

// 	buffer := make([]byte, len(result))
// 	copy(buffer, result)
// 	resultCopy := string(buffer)
// 	// c.SendString(resultCopy)

// }

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// app.Get("/:foo", handler)

	app.Listen(":3000")
}
