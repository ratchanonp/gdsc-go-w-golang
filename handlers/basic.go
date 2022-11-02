package handlers

import "github.com/gofiber/fiber/v2"

func Get(c *fiber.Ctx) error {
	return c.SendString("This is a GET request")
}

func Post(c *fiber.Ctx) error {
	return c.SendString("This is a POST request")
}

func Put(c *fiber.Ctx) error {
	return c.SendString("This is a PUT request")
}

func Patch(c *fiber.Ctx) error {
	return c.SendString("This is a PATCH request")
}

func Delete(c *fiber.Ctx) error {
	return c.SendString("This is a DELETE request")
}
