package routes

import (
	"encoding/json"
	"net/http"

	c "github.com/jmilagroso/api/core"
	h "github.com/jmilagroso/api/helpers"
	m "github.com/jmilagroso/api/models"
)

// AuthDBClient db client(s) local type
type AuthDBClient m.DBClient

var authBackend = c.InitJWTAuthenticationBackend()

// Auth authentication route
func (authDBClient *AuthDBClient) Auth(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := h.Hash256(r.FormValue("password"))

	var u m.User

	_, err := authDBClient.
		Query(&u, `SELECT id FROM users WHERE username = ? AND password = ?`, username, password)
	h.Error(err)

	if u.ID != "" {
		h.Error(json.NewEncoder(w).Encode(authBackend.GenerateToken(u.ID)))
	} else {
		h.Error(json.NewEncoder(w).Encode(m.ErrorResponse{Message: "Username/password is invalid", Status: 200}))
	}

}
