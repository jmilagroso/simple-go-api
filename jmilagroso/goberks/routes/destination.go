// destination.go
// Destination endpoints
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	redis "gopkg.in/redis.v5"
	"quadx.xyz/jmilagroso/goberks/blueprints"
	h "quadx.xyz/jmilagroso/goberks/helpers"
)

// DestinationDBClient db client(s) local type
type DestinationDBClient blueprints.DBClient

// DestinationResult Result local type
type DestinationResult = blueprints.Result

// DestinationResponse Response local type
type DestinationResponse = blueprints.Response

// PublishDestinationCoordinates - Publish destination coordinates endpoint
func (dbClient *DestinationDBClient) PublishDestinationCoordinates(w http.ResponseWriter, r *http.Request) {
	// Object Unique Identifier
	id := r.FormValue("id")

	// Object Subscribed Channel
	channel := r.FormValue("channel")

	// Object x,y coordinates
	latF, latFError := strconv.ParseFloat(r.FormValue("lat"), 64)
	lngF, lngFError := strconv.ParseFloat(r.FormValue("lng"), 64)

	h.Error(latFError)
	h.Error(lngFError)

	// Set current Object x,y coordinates
	cmd := redis.NewStringCmd("SET", channel, id, "POINT", latF, lngF)
	h.Error(dbClient.Process(cmd))
	res, cmdError := cmd.Result()
	h.Error(cmdError)

	// Delete current Object hook if any.
	// Should run concurrently.
	go dbClient.Process(redis.NewStringCmd("DELHOOK", id))

	// Construct WebHook URL: http(s)://<base url>:<port>/notify
	var bufferURL bytes.Buffer
	bufferURL = h.Concat(bufferURL, (map[bool]string{true: "http://", false: "https://"})[len(r.URL.Scheme) == 0])
	bufferURL = h.Concat(bufferURL, r.Host)
	bufferURL = h.Concat(bufferURL, "/notify")

	// Set hook so that current object (x,y)
	// will detect any object (x,y) if [inside,outside,enter,exit,cross]
	// calls http endopoint that sends [email (via sendgrid), sms (via twilio)].
	// Should run concurrently.
	// Reference: https://tile38.com/commands/sethook/
	go dbClient.Process(redis.NewStringCmd(
		"SETHOOK",
		id,
		bufferURL.String(),
		"NEARBY",
		channel,
		"FENCE",
		"DETECT",
		"inside",
		"POINT",
		latF,
		lngF,
		300,
	))

	destRes := DestinationResult{Type: "Point", Coordinates: []float64{latF, lngF}}
	h.Error(json.Unmarshal([]byte(res), &destRes))

	h.Error(json.NewEncoder(w).Encode(DestinationResponse{Id: id, Channel: channel, Result: destRes}))
}

// GetDestinationCoordinates - Get destination coordinate endpoint.
func (dbClient *DestinationDBClient) GetDestinationCoordinates(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	channel := vars["channel"]

	cmd := redis.NewStringCmd("GET", channel, id)
	h.Error(dbClient.Process(cmd))
	res, cmdError := cmd.Result()
	h.Error(cmdError)

	destRes := DestinationResult{}
	h.Error(json.Unmarshal([]byte(res), &destRes))

	h.Error(json.NewEncoder(w).Encode(DestinationResponse{Id: id, Channel: channel, Result: destRes}))
}
