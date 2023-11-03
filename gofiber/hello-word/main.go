package main 

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)


func main()  {
	app := fiber.New() //Fiber uygulaması oluştur 

	app.Get("/", func(c *fiber.Ctx) error { // / yolunda get istegi oluşturur 
		fmt.Println("Hello word") // hello word yazar 
		return c.JSON(fiber.Map{ // dönütü json olarak ayarlar  ve hello:word olarak dönderir
          "hello":"word",
		})
	})
	
   app.Listen(":8080") //:8080 portunda uygulamayı başlatır
}