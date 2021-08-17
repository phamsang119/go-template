package redis

import (
	"context"
	"fmt"
	"food-app/config"
	"github.com/go-redis/redis/v8"
)

type RedisRepo struct {
	Auth   AuthRepo
	Client *redis.Client
}

func NewRedisDB(ctx context.Context) (*RedisRepo, error) {
	redisHost := config.Env().GetRedisHost()
	redisPort := config.Env().GetRedisPort()
	redisPassword := config.Env().GetRedisPassword()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisHost, redisPort),
		Password: redisPassword,
		DB:       0,
	})
	return &RedisRepo{
		Auth:   NewAuth(redisClient, ctx),
		Client: redisClient,
	}, nil
}
