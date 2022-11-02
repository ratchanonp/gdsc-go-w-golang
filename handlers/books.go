package handlers

import (
	types "go-w-golang/basic/Types"

	"github.com/gofiber/fiber/v2"
)

var books = make([]types.Book, 0)

func GetBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}

	return c.SendStatus(404)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(types.Book)

	if err := c.BodyParser(book); err != nil {
		return err
	}

	book.ID = len(books) + 1
	books = append(books, *book)

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")

	for index, book := range books {
		if book.ID == id {
			books = append(books[:index], books[index+1:]...)
			return c.SendStatus(200)
		}
	}

	return c.SendStatus(404)
}
