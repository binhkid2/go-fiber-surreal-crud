package router

import (
	database "fiberr/Database"
	model "fiberr/Model"

	"github.com/gofiber/fiber/v2"
)

func CreateBioLink(c *fiber.Ctx) error {

	body := new(model.BioLink)

	if err := c.BodyParser(body); err != nil {
		return err
	}
	_, err := database.QueryCreate(body.Title, body.Link)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "Biolink not created",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Biolink  created",
	})
}
