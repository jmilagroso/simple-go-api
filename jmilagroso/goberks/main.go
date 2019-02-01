// main.go
// Main application
// Jay Milagroso <jmilagroso@quadx.xyz> / Jan 24 2019

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"quadx.xyz/jmilagroso/goberks/blueprints"

	h "quadx.xyz/jmilagroso/goberks/helpers"

	"github.com/gorilla/mux"
	redis "gopkg.in/redis.v5"
	"quadx.xyz/jmilagroso/goberks/routes"

	"github.com/go-pg/pg"

	m "quadx.xyz/jmilagroso/goberks/middlewares"
)

var redisClient *redis.Client
var pgsqlDB *pg.DB

var dbClient blueprints.DBClient

func main() {
	parsedURL, err := url.Parse(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	pgOptions := &pg.Options{
		User:     parsedURL.User.Username(),
		Database: parsedURL.Path[1:],
		Addr:     parsedURL.Host,
	}

	if password, ok := parsedURL.User.Password(); ok {
		pgOptions.Password = password
	}

	pgsqlDB := pg.Connect(pgOptions)

	defer pgsqlDB.Close()
	// --- Postgresql Server Connection --- //

	// --- Tile38 Server Connection --- //
	// redisClient = redis.NewClient(&redis.Options{
	// 	Addr: h.GetEnvValue("TILE38_ADDRESS"),
	// })

	// // Set json output logging
	// h.Error(redisClient.Process(redis.NewStringCmd("OUTPUT", "json")))

	// defer redisClient.Close()
	// --- Tile38 Server Connection --- //

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully "+
			"wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	// @TODO Create Pub/Sub version of current impl.

	// // --- Rider Endpoint --- //
	// rider := routes.RiderDBClient{DB: pgsqlDB, Client: redisClient}
	// // Sends rider's coordinates
	// r.HandleFunc("/rider/coordinates", rider.PublishRiderCoordinates).Methods("POST")
	// // Gets rider's coordinates
	// r.HandleFunc("/rider/coordinates/{id}/{channel}", rider.GetRiderCoordinates).Methods("GET")
	// // --- Rider Endpoint --- //

	// // --- Destination Endpoint --- //
	// destination := routes.DestinationDBClient{DB: pgsqlDB, Client: redisClient}
	// // Sends destination's coordinates
	// r.HandleFunc("/destination/coordinates", destination.PublishDestinationCoordinates).Methods("POST")
	// // Gets destination's coordinates
	// r.HandleFunc("/destination/coordinates/{id}/{channel}", destination.GetDestinationCoordinates).Methods("GET")
	// // --- Destination Endpoint --- //

	// // Webhook notifier
	// r.HandleFunc("/notify", routes.Notify).Methods("POST")

	// Webhook notifier
	r.HandleFunc("/", routes.GetIndex).Methods("GET")

	index := routes.IndexDBClient{DB: pgsqlDB, Client: redisClient}
	r.HandleFunc("/users", index.GetUsers).Methods("GET")

	r.Use(m.JSON)

	srv := &http.Server{
		Addr: ":" + os.Getenv("PORT"),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	h.Error(srv.Shutdown(ctx))
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
