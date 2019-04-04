package models

import (
	"github.com/go-pg/pg"
)

// DBClient struct
type DBClient struct {
	*pg.DB
}
