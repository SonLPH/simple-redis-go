package db

import (
	"context"
	"simple-redis-go/config"

	"github.com/redis/go-redis/v9"
)

func RedisConnection(conf config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: conf.RedisURL,
	})

	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	return rdb, err
}
