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
	"github.com/go-redis/redis"

	m "github.com/jmilagroso/api/middlewares"
)

var pgsqlDB *pg.DB
var redisClient *redis.Client

var dbClient models.DBClient

func main() {
	// --- Postgresql Server Connection --- //
	options, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	h.Error(err)

	options.TLSConfig.InsecureSkipVerify = true
	pgsqlDB = pg.Connect(options)
	defer pgsqlDB.Close()
	// --- Postgresql Server Connection --- //

	// --- Redis Server Connection --- //
	redisClient := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		DB:   0, // use default DB
	})
	defer redisClient.Close()
	// --- Redis Server Connection --- //

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*30,
		"the duration for which the server gracefully "+
			"wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()

	r.Use(m.JSONMiddleware)

	//r.Handle("/", m.AuthMiddleware(http.HandlerFunc(routes.GetIndex))).Methods("GET")
	r.Handle("/", http.HandlerFunc(routes.GetIndex)).Methods("GET")

	// Users endpoints.
	users := routes.IndexDBClient{DB: pgsqlDB, Client: redisClient}
	r.HandleFunc("/user", users.NewUser).Methods("POST")

	r.Handle("/users", m.AuthMiddleware(http.HandlerFunc(users.GetUsers))).Methods("GET")

	r.Handle("/users/{id:[0-9]+}", m.AuthMiddleware(http.HandlerFunc(users.GetUser))).Methods("GET")

	r.Handle("/users/{page:[0-9]+}/{per_page:[0-9]+}",
		m.AuthMiddleware(
			http.HandlerFunc(
				users.GetUsersPaginated))).
		Methods("GET")

	auth := routes.AuthDBClient{DB: pgsqlDB, Client: redisClient}
	r.HandleFunc("/auth", auth.Auth).Methods("POST")

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
