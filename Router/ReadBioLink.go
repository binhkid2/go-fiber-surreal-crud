package router

import (
	database "fiberr/Database"

	"github.com/gofiber/fiber/v2"
)

func ReadBioLink(c *fiber.Ctx) error {

	rows, err := database.QueryRead(c)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "loi ",
		})
	}
	c.Status(200).JSON(&fiber.Map{
		"biolinks": rows,
	})
	return nil
}
