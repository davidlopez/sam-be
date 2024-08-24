package root

import "github.com/gofiber/fiber/v2"

func Root() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}
}
