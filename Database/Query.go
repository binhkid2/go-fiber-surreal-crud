package database

import (
	model "fiberr/Model"
	util "fiberr/Util"
	"github.com/gofiber/fiber/v2"
)

func QueryCreate(title string, link string) (interface{}, error) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}
	data, err := db.Query("INSERT INTO biolinks (title, link) VALUES ($title, $link);", model.BioLink{
		Title: title,
		Link:  link,
	})
	if err != nil {
		panic(err)
	}
	util.Log("Database", "Biolink created")

	return data, nil
}

var (
	biolinks = map[string]model.BioLink{}
)

func QueryRead(c *fiber.Ctx) (interface{}, error) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}
	//Select all query
	data, err := db.Select("biolinks")
	if err != nil {
		panic(err)
	}
	util.Log("Database", " all Biolink ")
	c.Status(200).JSON(&fiber.Map{
		"biolinks": data,
	})
	return data, nil
}
func QueryDelete(id string) (interface{}, error) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}
	//Delete query
	data, err := db.Delete(id)
	if err != nil {
		panic(err)
	}
	util.Log("Database", " Delete 1 biolink ")
	return data, nil
}
func QueryUpdate(id string, title string, link string) (interface{}, error) {
	var err error
	// access database use namespace and database
	if _, err = db.Use("test", "test"); err != nil {
		panic(err)
	}
	//Update query

	if _, err = db.Update(id, map[string]interface{}{
		"title": title,
		"link":  link}); err != nil {
		panic(err)
	}

	util.Log("Database", " Update 1 biolink ")
	return db, nil
}
