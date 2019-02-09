package routes

import (
	"encoding/json"
	"net/http"

	"runtime"
	"time"

	"github.com/gorilla/mux"
	h "github.com/jmilagroso/api/helpers"
	m "github.com/jmilagroso/api/models"
)

// IndexDBClient db client(s) local type
type IndexDBClient m.DBClient

// GetIndex - Get index route
func GetIndex(w http.ResponseWriter, r *http.Request) {
	h.Error(json.NewEncoder(w).Encode(m.Index{ServerTime: time.Now().String(), GoVersion: runtime.Version()}))
}

// GetUsers - Get users route
func (dbClient *IndexDBClient) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Collection container
	var rows []m.User

	// @TODO Inject Caching using Heroku Redis (Needs account upgrade)
	// @TODO Cache page_per_page = 1_10
	// @TODO Cache timeout 30 secs

	// Add resultset to `rows` variable via reference
	dbClient.Query(&rows, `SELECT id, username, email FROM users ORDER BY id DESC`)

	h.Error(json.NewEncoder(w).Encode(rows))

}

// GetUsersPaginated - Get index paginated route
func (dbClient *IndexDBClient) GetUsersPaginated(w http.ResponseWriter, r *http.Request) {

	// Get request variables
	vars := mux.Vars(r)

	page := h.StrToInt(vars["page"])
	perPage := h.StrToInt(vars["per_page"])
	offset := (page - 1) * perPage

	// @TODO Inject Caching using Heroku Redis (Needs account upgrade)
	// @TODO Cache page_per_page = 1_10
	// @TODO Cache timeout 30 secs

	// Output
	var rows []m.User

	dbClient.Query(&rows, `SELECT id, username, email FROM users ORDER BY id DESC LIMIT ? OFFSET ?`, perPage, offset)

	h.Error(json.NewEncoder(w).Encode(rows))
}
