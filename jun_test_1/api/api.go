package api

import (
	"github.com/St0rmPetrel/Golang/jun_test_1/db"
	"github.com/gofiber/fiber/v2"
)

var rdb *db.DB

func Up(dataBase *db.DB) {
	rdb = dataBase
	app := fiber.New()
	app.Get("/json/hackers", returnHackers)
	app.Listen(":8010")
}

func returnHackers(c *fiber.Ctx) error {
	data, err := rdb.TakeDataWithCache()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.JSON(&data)
}
