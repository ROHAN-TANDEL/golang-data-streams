package redis

import (
  "github.com/redis/go-redis/v9"
)

func NewClient(cfg struct {
	Addr     string
	Password string
	DB       int
}) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
