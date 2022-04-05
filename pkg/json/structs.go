package json

import "github.com/gofiber/fiber/v2"

func Error(err string) fiber.Map {
	return fiber.Map{
		"error": err,
	}
}
