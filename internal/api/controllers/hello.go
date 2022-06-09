package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Hello(cb *fiber.Ctx) error {

	return cb.SendString("Hello, World!")
}
