package routes

import (
	"encoding/json"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	h "github.com/jmilagroso/api/helpers"
	m "github.com/jmilagroso/api/models"
)

// IndexDBClient db client(s) local type
type IndexDBClient m.DBClient

// GetIndex - Get index route
func (dbClient *IndexDBClient) GetIndex(w http.ResponseWriter, r *http.Request) {

	cache := h.CacheDBClient{DB: dbClient.DB, Conn: dbClient.Conn}
	var cacheKey = "index"

	if cache.CacheExists(cacheKey) {
		h.Error(json.NewEncoder(w).Encode(cache.Get(cacheKey, m.Index{})))
	} else {
		h.Error(json.NewEncoder(w).Encode(cache.Set(cacheKey, 60,
			m.Index{
				ServerTime:   time.Now().String(),
				GoVersion:    runtime.Version(),
				CacheTimeout: 60})))
	}
}

// GetUsers - Get users
func (dbClient *IndexDBClient) GetUsers(w http.ResponseWriter, r *http.Request) {

	cache := h.CacheDBClient{DB: dbClient.DB, Conn: dbClient.Conn}
	cacheKey := "users"

	if cache.CacheExists(cacheKey) {
		h.Error(json.NewEncoder(w).Encode(cache.Get(cacheKey, []m.User{})))
	} else {
		// Collection container
		var users []m.User

		// Add resultset to `rows` variable via reference
		_, err := dbClient.Query(&users, `SELECT id, username, email FROM users ORDER BY id DESC`)
		h.Error(err)

		if users != nil {
			h.Error(json.NewEncoder(w).Encode(cache.Set(cacheKey, 60, users)))
		} else {
			h.Error(json.NewEncoder(w).Encode(m.ErrorResponse{Message: "No Record(s) Found.", Status: 200}))
		}
	}
}

// GetUsersPaginated - Get users with paginated result
func (dbClient *IndexDBClient) GetUsersPaginated(w http.ResponseWriter, r *http.Request) {

	// Get request variables
	vars := mux.Vars(r)

	page := h.StrToInt(vars["page"])
	perPage := h.StrToInt(vars["per_page"])
	offset := (page - 1) * perPage

	cache := h.CacheDBClient{DB: dbClient.DB, Conn: dbClient.Conn}
	cacheKey := strings.Join([]string{"page_per_page_", strconv.Itoa(perPage), strconv.Itoa(offset)}, "")

	if cache.CacheExists(cacheKey) {
		h.Error(json.NewEncoder(w).Encode(cache.Get(cacheKey, []m.User{})))
	} else {
		// Output
		var users []m.User

		_, err := dbClient.Query(
			&users,
			`SELECT id, username, email FROM users ORDER BY id DESC LIMIT ? OFFSET ?`,
			perPage,
			offset)
		h.Error(err)

		if users != nil {
			h.Error(json.NewEncoder(w).Encode(cache.Set(cacheKey, 60, users)))
		} else {
			h.Error(json.NewEncoder(w).Encode(m.ErrorResponse{Message: "No Record(s) Found.", Status: 200}))
		}
	}
}

// GetUser - Get user by id
func (dbClient *IndexDBClient) GetUser(w http.ResponseWriter, r *http.Request) {

	// Get request variables
	vars := mux.Vars(r)

	id := h.StrToInt(vars["id"])

	cache := h.CacheDBClient{DB: dbClient.DB, Conn: dbClient.Conn}
	cacheKey := strings.Join([]string{"user_", strconv.Itoa(id)}, "")

	if cache.CacheExists(cacheKey) {
		h.Error(json.NewEncoder(w).Encode(cache.Get(cacheKey, m.User{})))
	} else {
		// Output
		var user m.User

		_, err := dbClient.Query(&user, `SELECT id, username, email FROM users WHERE id = ?`, id)
		h.Error(err)

		if user != (m.User{}) {
			h.Error(json.NewEncoder(w).Encode(cache.Set(cacheKey, 60, user)))
		} else {
			h.Error(json.NewEncoder(w).Encode(m.ErrorResponse{Message: "No Record(s) Found.", Status: 200}))
		}
	}
}

// NewUser - New user
func (dbClient *IndexDBClient) NewUser(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	email := r.FormValue("email")
	password := h.Hash256(r.FormValue("password"))

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
