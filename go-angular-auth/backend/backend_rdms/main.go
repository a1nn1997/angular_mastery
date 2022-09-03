package main

import (
	"github.com/a1nn1997/go-auth/database"
	"github.com/a1nn1997/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)

func main() {
	//db conection
	database.Connect()

	app := fiber.New()
	//cors ke lie v2
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	//set up routes
	routes.Setup(app)
	//port
	app.Listen(":8069")
}