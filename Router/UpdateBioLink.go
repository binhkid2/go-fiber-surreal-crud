package router

import (
	database "fiberr/Database"

	"github.com/gofiber/fiber/v2"
)

func UpdateBioLink(c *fiber.Ctx) error {
	id := c.Params("id")
	title := c.FormValue("title")
	link := c.FormValue("link")
	_, err := database.QueryUpdate(id, title, link)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Can not Update",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Updated  1 biolink",
	})
}
