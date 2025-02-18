package store

import "github.com/redis/go-redis/v9"

type RedisStore struct {
}

func Connect() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
