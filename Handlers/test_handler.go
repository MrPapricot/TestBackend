package Handlers

import "github.com/gofiber/fiber/v2"

func Test(context *fiber.Ctx) error {
	return context.Status(200).JSON(fiber.Map{"message": "Всё окей"})
}
