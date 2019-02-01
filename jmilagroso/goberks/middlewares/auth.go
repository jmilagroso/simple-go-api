package middlewares

import (
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		w.Header().Set("Content-Type", "application/json")

		if token == "123" {
			// We found the token in our map
			log.Printf("Authenticated")
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {

			// Write an error and stop the handler chain
			http.Error(w, "{\"status\": \"Forbidden\"}", http.StatusForbidden)
		}
	})
}
