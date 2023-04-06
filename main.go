package main

import (
	initializers "goTest/Initializers"
	router "goTest/Router"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func runServer() error {
	err := initializers.LoadEnvVariables()
	if err != nil {
		return err
	}

	err = initializers.Connect()
	if err != nil {
		panic(err)
	}

	defer initializers.CloseDB()

	port := os.Getenv("PORT")

	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())

	router.AddNutritionGroup(app)
	router.AddAuthGroup(app)

	log.Fatal(app.Listen(":" + port))

	return nil
}
func main() {
	err := runServer()
	if err != nil {
		panic(err)
	}
}
