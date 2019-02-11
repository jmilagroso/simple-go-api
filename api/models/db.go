package models

import (
	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
)

// DBClient struct
type DBClient struct {
	*pg.DB
	*redis.Client
}
