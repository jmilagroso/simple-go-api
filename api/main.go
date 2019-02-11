// main.go

package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmilagroso/api/models"

	h "github.com/jmilagroso/api/helpers"

	"github.com/gorilla/mux"
	"github.com/jmilagroso/api/routes"

	"github.com/go-pg/pg"

	m "github.com/jmilagroso/api/middlewares"
)

var pgsqlDB *pg.DB

var dbClient models.DBClient

func main() {
	// --- Postgresql Server Connection --- //
	options, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	h.Error(err)

	options.TLSConfig.InsecureSkipVerify = true
	pgsqlDB = pg.Connect(options)
	defer pgsqlDB.Close()
	// --- Postgresql Server Connection --- //

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*30,
		"the duration for which the server gracefully "+
			"wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	// Default
	r.HandleFunc("/", routes.GetIndex).Methods("GET")

	// Users endpoints.
	index := routes.IndexDBClient{DB: pgsqlDB}
	r.HandleFunc("/user", index.NewUser).Methods("POST")
	r.HandleFunc("/users", index.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", index.GetUser).Methods("GET")
	r.HandleFunc("/users/{page:[0-9]+}/{per_page:[0-9]+}", index.GetUsersPaginated).Methods("GET")

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
