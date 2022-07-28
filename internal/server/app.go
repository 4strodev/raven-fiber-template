package server

import "github.com/gofiber/fiber/v2"

var App *fiber.App

func init() {
    App = fiber.New()

    App.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("{{ .ProjectName }} works!!")
    })
}
