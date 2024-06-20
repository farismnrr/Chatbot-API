package database

import (
	"capstone-project/helper"
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Client *redis.Client
	Context context.Context
}

func NewRedisConnection() (*Redis, error) {
	opt, err := redis.ParseURL(helper.GetEnv("REDIS_URL"))
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opt)
	return &Redis{Client: client, Context: context.Background()}, nil
}

func (r *Redis) Close() error {
	return r.Client.Close()
}

func (r *Redis) Reset() error {
	return r.Client.FlushDB(context.Background()).Err()
}
