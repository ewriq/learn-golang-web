package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)


func main()  {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println("Hello word")
		return c.JSON(fiber.Map{
          "hello":"word",
		})
	})
	
   app.Listen(":8080")
}