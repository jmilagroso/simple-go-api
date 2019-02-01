// rider.go
// Rider endpoints
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package routes

import (
	"encoding/json"
	"net/http"

	"runtime"
	"time"

	h "quadx.xyz/jmilagroso/goberks/helpers"
)

type Index struct {
	ServerTime string
	GoVersion  string
}

// GetIndex - Get index route
func GetIndex(w http.ResponseWriter, r *http.Request) {
	h.Error(json.NewEncoder(w).Encode(Index{ServerTime: time.Now().String(), GoVersion: runtime.Version()}))
}
