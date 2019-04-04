package helpers

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	m "github.com/jmilagroso/api/models"
)

// CacheDBClient db client(s) local type
type CacheDBClient m.DBClient

// Cache
func (dbClient *CacheDBClient) Set(key string, ttl int64, structure interface{}) interface{} {
	serialized, err := json.Marshal(structure)
	Error(err)
	dbClient.Conn.Do("SET", key, string(serialized))
	dbClient.Conn.Do("EXPIRE", key, 60)

	var deserialized = structure
	val, err := redis.String(dbClient.Conn.Do("GET", key))

	err = json.Unmarshal([]byte(val), &deserialized)
	Error(err)

	return deserialized
}

// CacheExists
func (dbClient *CacheDBClient) CacheExists(key string) bool {

	val, err := redis.String(dbClient.Conn.Do("GET", key))
	Error(err)

	if val == "" {
		return false
	} else if err != nil {
		return false
	} else {
		return true
	}
}

// Get
func (dbClient *CacheDBClient) Get(key string, structure interface{}) interface{} {

	val, err := redis.String(dbClient.Conn.Do("GET", key))
	Error(err)

	var deserialized = structure

	if val == "" {
		deserialized = structure

	} else if err != nil {
		Error(err)
	} else {
		Error(err)
		err = json.Unmarshal([]byte(val), &deserialized)
		Error(err)
	}
	return deserialized
}
