package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	c "github.com/jmilagroso/api/core"
	h "github.com/jmilagroso/api/helpers"
)

// ErrorResponse error response
type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var authBackend = c.InitJWTAuthenticationBackend()

// AuthMiddleware auth middleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		paramToken := r.Header.Get("X-Token")

		if paramToken != "" {
			token, err := authBackend.ValidateToken(paramToken)

			if err == nil && token.Valid {
				next.ServeHTTP(w, r)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)

				if ve, ok := err.(*jwt.ValidationError); ok {
					if ve.Errors&jwt.ValidationErrorMalformed != 0 {
						h.Error(json.NewEncoder(w).Encode(ErrorResponse{Message: "Malformed", Status: 400}))
					} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						h.Error(json.NewEncoder(w).Encode(ErrorResponse{
							Message: "Token has expired or not active yet",
							Status:  400}))
					} else {
						h.Error(json.NewEncoder(w).Encode(ErrorResponse{Message: "Couldn't handle token", Status: 400}))
					}
				} else {
					h.Error(json.NewEncoder(w).Encode(ErrorResponse{Message: "Couldn't handle token", Status: 400}))
				}
			}
		} else {
			h.Error(json.NewEncoder(w).Encode(ErrorResponse{Message: "Forbidden", Status: 403}))
		}
	})
}
