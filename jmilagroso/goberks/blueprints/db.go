// redisclient.go
// Redis client struct
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package blueprints

import (
	"github.com/go-pg/pg"
	redis "gopkg.in/redis.v5"
)

type DBClient struct {
	*redis.Client
	*pg.DB
}
