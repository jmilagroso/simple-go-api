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

// GetUsers - Get users
func (dbClient *IndexDBClient) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Collection container
	var users []m.User

	// @TODO Inject Caching using Heroku Redis (Needs account upgrade)
	// @TODO Cache page_per_page = 1_10
	// @TODO Cache timeout 30 secs

	// Add resultset to `rows` variable via reference
	_, err := dbClient.Query(&users, `SELECT id, username, email FROM users ORDER BY id DESC`)
	h.Error(err)

	h.Error(json.NewEncoder(w).Encode(users))

}

// GetUsersPaginated - Get users with paginated result
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
	var users []m.User

	_, err := dbClient.Query(&users, `SELECT id, username, email FROM users ORDER BY id DESC LIMIT ? OFFSET ?`, perPage, offset)
	h.Error(err)

	h.Error(json.NewEncoder(w).Encode(users))
}

// GetUser - Get user by id
func (dbClient *IndexDBClient) GetUser(w http.ResponseWriter, r *http.Request) {

	// Get request variables
	vars := mux.Vars(r)

	id := h.StrToInt(vars["id"])

	// @TODO Inject Caching using Heroku Redis (Needs account upgrade)
	// @TODO Cache page_per_page = 1_10
	// @TODO Cache timeout 30 secs

	// Output
	var user m.User

	_, err := dbClient.Query(&user, `SELECT id, username, email FROM users WHERE id = ?`, id)
	h.Error(err)

	h.Error(json.NewEncoder(w).Encode(user))
}

// NewUser - New user
func (dbClient *IndexDBClient) NewUser(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := h.Hash256(r.FormValue("password"))

	// @TODO Inject Caching using Heroku Redis (Needs account upgrade)
	// @TODO Cache page_per_page = 1_10
	// @TODO Cache timeout 30 secs

	// Output
	var users []m.User

	exists, errExists := dbClient.Model(&users).Where("username = ?", username).WhereOr("email = ?", email).Exists()
	h.Error(errExists)

	newUser := m.User{Username: username, Email: email, Password: password}
	if !exists {

		errInsert := dbClient.Insert(&newUser)
		h.Error(errInsert)
	}

	h.Error(json.NewEncoder(w).Encode(newUser))
}
