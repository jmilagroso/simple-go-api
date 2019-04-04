package models

import (
	"github.com/go-pg/pg"
	"github.com/gomodule/redigo/redis"
)

// DBClient struct
type DBClient struct {
	*pg.DB
	redis.Conn
}
