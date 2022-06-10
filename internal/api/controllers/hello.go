package controllers

import (
	"github.com/backsoul/doberman/pkg/services"
	"github.com/gofiber/fiber/v2"
)

func Hello(cb *fiber.Ctx) error {
	msg := cb.Query("message")
	services.SendMessage(msg)
	return cb.SendString(msg)
}
