package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type DB struct {
	client *redis.Client
}

type Hacker struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var ctx = context.Background()

func (rdb *DB) LoadData(data *[]Hacker) error {
	ret := rdb.client.ZRangeWithScores(ctx, "hackers", 0, -1)
	if ret.Err() != nil {
		return ret.Err()
	}
	*data = make([]Hacker, len(ret.Val()))
	for id, v := range ret.Val() {
		(*data)[id].Name = v.Member.(string)
		(*data)[id].Score = (int)(v.Score)
	}
	return nil
}

func NewData() []Hacker {
	return []Hacker{}
}

func NewDB() (*DB, error) {
	rdb := connect()
	err := initDB(rdb)
	return &DB{rdb}, err
}

func connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
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
