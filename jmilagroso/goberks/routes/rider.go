// rider.go
// Rider endpoints
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	redis "gopkg.in/redis.v5"
	"quadx.xyz/jmilagroso/goberks/blueprints"
	h "quadx.xyz/jmilagroso/goberks/helpers"
)

type Address struct {
	Id      int64
	PartyId int64
	Xcode   string
}

// RiderDBClient db client(s) local type
type RiderDBClient blueprints.DBClient

// RiderResult Result local type
type RiderResult = blueprints.Result

// RiderResponse Response local type
type RiderResponse = blueprints.Response

// PublishRiderCoordinates - Publish rider coordinates
func (dbClient *RiderDBClient) PublishRiderCoordinates(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	channel := r.FormValue("channel")

	latF, latFError := strconv.ParseFloat(r.FormValue("lat"), 64)
	lngF, lngFError := strconv.ParseFloat(r.FormValue("lng"), 64)

	h.Error(latFError)
	h.Error(lngFError)

	cmd := redis.NewStringCmd("SET", channel, id, "POINT", latF, lngF)
	h.Error(dbClient.Process(cmd))

	res, cmdError := cmd.Result()
	h.Error(cmdError)

	riderRes := RiderResult{Type: "Point", Coordinates: []float64{latF, lngF}}
	h.Error(json.Unmarshal([]byte(res), &riderRes))

	h.Error(json.NewEncoder(w).Encode(RiderResponse{Id: id, Channel: channel, Result: riderRes}))
}

// GetRiderCoordinates - Get rider latest coordinates
func (dbClient *RiderDBClient) GetRiderCoordinates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	channel := vars["channel"]

	cmd := redis.NewStringCmd("GET", channel, id)
	h.Error(dbClient.Process(cmd))

	res, cmdError := cmd.Result()
	h.Error(cmdError)

	riderRes := RiderResult{}
	h.Error(json.Unmarshal([]byte(res), &riderRes))

	h.Error(json.NewEncoder(w).Encode(RiderResponse{Id: id, Channel: channel, Result: riderRes}))
}
