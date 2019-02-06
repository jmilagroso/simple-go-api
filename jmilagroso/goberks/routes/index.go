// rider.go
// Rider endpoints
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package routes

import (
	"encoding/json"
	"net/http"

	"runtime"
	"time"

	"github.com/gorilla/mux"
	"quadx.xyz/jmilagroso/goberks/blueprints"
	h "quadx.xyz/jmilagroso/goberks/helpers"
)

type Index struct {
	ServerTime string `json:"server_time"`
	GoVersion  string `json:"server_time"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
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

// GetUsers - Get users route
func (dbClient *IndexDBClient) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Output
	var rows []User

	dbClient.Query(&rows, `SELECT id, username, email FROM users ORDER BY id DESC`)

	h.Error(json.NewEncoder(w).Encode(rows))
}

// GetUsersPaginated - Get index paginated route
func (dbClient *IndexDBClient) GetUsersPaginated(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	page := h.StrToInt(vars["page"])
	perPage := h.StrToInt(vars["per_page"])
	offset := (page - 1) * perPage

	// Output
	var rows []User

	dbClient.Query(&rows, `SELECT id, username, email FROM users ORDER BY id DESC LIMIT ? OFFSET ?`, perPage, offset)

	h.Error(json.NewEncoder(w).Encode(rows))
}
