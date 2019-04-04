package routes

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	h "github.com/jmilagroso/api/helpers"
	m "github.com/jmilagroso/api/models"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/redis.v5"
)

// IndexDBClient db client(s) local type
type IndexDBClient m.DBClient

// GetIndex - Get index route
func (dbClient *IndexDBClient) GetIndex(w http.ResponseWriter, r *http.Request) {
	key := "index"
	val, err := dbClient.Get(key).Result()
	h.Error(err)

	var item m.Index

	if err == redis.Nil {
		item = m.Index{ServerTime: time.Now().String(), GoVersion: runtime.Version()}
		binary, err := msgpack.Marshal(&item)
		h.Error(err)
		dbClient.Set(key, binary, 60*time.Second)
	} else if err != nil {
		h.Error(err)
	} else {
		err = msgpack.Unmarshal([]byte(val), &item)
		h.Error(err)
	}

	h.Error(json.NewEncoder(w).Encode(item))
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

	if user.ID != "" {
		h.Error(json.NewEncoder(w).Encode(user))
	} else {
		h.Error(json.NewEncoder(w).Encode(m.ErrorResponse{Message: "ID record not found.", Status: 200}))
	}

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

		h.Error(json.NewEncoder(w).Encode(newUser))
	} else {
		h.Error(json.NewEncoder(w).Encode(m.ErrorResponse{Message: "Username/email already exists", Status: 200}))
	}

}
