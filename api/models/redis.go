package models

import (
	"github.com/go-redis/redis"
)

// RedisClient struct
type RedisClient struct {
	*redis.Client
}
