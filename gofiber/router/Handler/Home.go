package Handler

import (
	"github.com/gofiber/fiber/v2"

)

func Home(c *fiber.Ctx) error {
	c.JSON(fiber.Map{ //Json olarak  { "hello":"word" } döndürülür
		"hello":"world",
	})
	return nil
}