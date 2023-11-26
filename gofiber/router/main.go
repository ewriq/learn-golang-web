package main

import (
	Routers "test2/Router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New() // Fiber uygulaması oluşturur

	Routers.Initalize(app) //Yönendiriciye oluşturulan uygulama verilir

	app.Listen(":8080") //Sunucu 8080 portunda başlatılır 
}
