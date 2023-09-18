package main

import (
	database "fiberr/Database"
	"fiberr/Router"
	util "fiberr/Util"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Post("/biolinks", router.CreateBioLink)
	app.Get("/biolinks", router.ReadBioLink)
	app.Delete("/biolinks/:id", router.DeleteBioLink)
	app.Put("/biolinks/:id", router.UpdateBioLink)
	database.Connect()
	log.Fatal(app.Listen("127.0.0.1:" + util.GetConfig("server", "port")))

}
