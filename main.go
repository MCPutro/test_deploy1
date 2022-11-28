package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}

	err := app.Listen(":" + PORT)
	if err != nil {
		fmt.Println(err)
	}
}
