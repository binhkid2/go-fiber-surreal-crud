package router

import (
	database "fiberr/Database"

	"github.com/gofiber/fiber/v2"
)

func DeleteBioLink(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.QueryDelete(id)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Can not delete",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Biolink  deleted",
	})
}
