package api

import (
	"github.com/St0rmPetrel/Golang/jun_test_1/db"
	"github.com/gofiber/fiber/v2"
)

var rdb *db.DB

func Up(dataBase *db.DB) {
	app := setupApp(dataBase)
	app.Listen(":8010")
}

func setupApp(dataBase *db.DB) *fiber.App {
	rdb = dataBase
	app := fiber.New()
	app.Get("/json/hackers", returnHackers)
	return app
}

func returnHackers(c *fiber.Ctx) error {
	data, err := rdb.TakeDataWithCache()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}
	return c.JSON(&data)
}
