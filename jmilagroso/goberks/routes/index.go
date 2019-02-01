// rider.go
// Rider endpoints
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package routes

import (
	"encoding/json"
	"net/http"

	"runtime"
	"time"

	"quadx.xyz/jmilagroso/goberks/blueprints"
	h "quadx.xyz/jmilagroso/goberks/helpers"
)

type Index struct {
	ServerTime string
	GoVersion  string
}

type User struct {
	Id           string
	Email        string
	RegisteredOn string
	PublicId     string
	Username     string
}

// IndexDBClient db client(s) local type
type IndexDBClient blueprints.DBClient

// IndexResult Result local type
type IndexResult = blueprints.Result

// IndexResponse Response local type
type IndexResponse = blueprints.Response

// GetIndex - Get index route
func GetIndex(w http.ResponseWriter, r *http.Request) {
	h.Error(json.NewEncoder(w).Encode(Index{ServerTime: time.Now().String(), GoVersion: runtime.Version()}))
}

// GetIndex - Get index route
func (dbClient *IndexDBClient) GetUsers(w http.ResponseWriter, r *http.Request) {

	var rows []User

	_, dbErr := dbClient.Query(&rows, `SELECT id, email, registered_on, public_id, username FROM user ORDER BY id DESC`)
	h.Error(dbErr)

	h.Error(json.NewEncoder(w).Encode(rows))
}
