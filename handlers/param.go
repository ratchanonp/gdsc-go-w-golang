package handlers

import "github.com/gofiber/fiber/v2"

func WildCard(c *fiber.Ctx) error {
	return c.SendString("this is a wildcard route : " + c.Params("*"))
}

func Resturant(c *fiber.Ctx) error {
	return c.SendString("this is a resturant route : " + c.Params("id") + " " + c.Params("food_id"))
}
