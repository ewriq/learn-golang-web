package Router

import (
	"test2/Handler"

	"github.com/gofiber/fiber/v2"
)

func Initalize(app *fiber.App){

    app.Get("/", Handler.Home) //Home import edilir router / olarak belirlenir 
	
}