package main

import (
	"fmt"
	"net/http"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbplayer"
	player "github.com/JordanRad/play-j/backend/cmd/playerd/internal/player"
	"github.com/JordanRad/play-j/backend/internal/cloudstorage"
	playersrv "github.com/JordanRad/play-j/backend/internal/playerservice/gen/http/player/server"
	playersvc "github.com/JordanRad/play-j/backend/internal/playerservice/gen/player"
	goahttp "goa.design/goa/v3/http"

	"github.com/JordanRad/play-j/backend/internal/middleware"

	"database/sql"
	"time"

	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	config, err := configFromEnv()

	if err != nil {
		log.Fatalf("Config file cannot be read: %v", err)
	}

	dbConnectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DBName)
	fmt.Println(dbConnectionString)
	var db *sql.DB
	db, err = sql.Open("pgx", dbConnectionString)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Maximum Idle Connections
	db.SetMaxIdleConns(config.Postgres.MaxIdleConns)
	// Maximum Open Connections
	db.SetMaxOpenConns(config.Postgres.MaxOpenConns)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	if err := db.Ping(); err != nil {
		log.Fatalf("Unable to reach database: %v", err)
	}

	fmt.Print("Connected to database \n")

	dbPlayerStore := &dbplayer.Store{
		DB: db,
	}

	musicStorage := cloudstorage.NewCloudStorage("")

	playerService := player.NewService(dbPlayerStore, musicStorage)

	var accountEndpoints *playersvc.Endpoints = playersvc.NewEndpoints(playerService)

	// Provide the transport specific request decoder and response encoder.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	var playerServer *playersrv.Server = playersrv.New(accountEndpoints, mux, dec, enc, nil, nil)
	playersrv.Mount(mux, playerServer)

	playerServer.Use(middleware.AuthenticateRequest())

	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	fmt.Printf("Account service has just started on %s ...\n", address)
	http.ListenAndServe(address, mux)
}
