package main

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/json/hackers", func(c *fiber.Ctx) error {
		data, err := give_data()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString(err.Error())
		}
		return c.JSON(data)
	})
	app.Listen(":3000")
}

var ctx = context.Background()

func give_data() ([]redis.Z, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := initDB(rdb)
	if err != nil {
		return nil, err
	}
	ret := rdb.ZRangeWithScores(ctx, "hackers", 0, -1)
	if ret.Err() != nil {
		return nil, err
	}
	return ret.Val(), nil
}

func initDB(c *redis.Client) error {
	return c.ZAdd(ctx, "hackers",
		&redis.Z{1953, "Richard Stallman"},
		&redis.Z{1940, "Alan Kay"},
		&redis.Z{1965, "Yukihiro Matsumoto"},
		&redis.Z{1916, "Claude Shannon"},
		&redis.Z{1969, "Linus Torvalds"},
		&redis.Z{1912, "Alan Turing"},
	).Err()
}
