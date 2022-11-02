package handlers

import "github.com/gofiber/fiber/v2"

func Download(c *fiber.Ctx) error {
	return c.Download("public/download.pdf", "gofiber-wiki.pdf")
}
