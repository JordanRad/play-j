package main

import (
	"fmt"
	"net/http"

	account "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account"
	dbaccount "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbaccount"
	dbplaylist "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbplaylist"
	playlist "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/playlist"
	accountsvc "github.com/JordanRad/play-j/backend/internal/accountservice/gen/account"
	accountsrv "github.com/JordanRad/play-j/backend/internal/accountservice/gen/http/account/server"
	playlistsrv "github.com/JordanRad/play-j/backend/internal/accountservice/gen/http/playlist/server"
	playlistsvc "github.com/JordanRad/play-j/backend/internal/accountservice/gen/playlist"
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

	if err = db.Ping(); err != nil {
		log.Fatalf("Unable to reach database: %v", err)
	}

	fmt.Print("Connected to database \n")

	dbAccountStore := &dbaccount.Store{
		DB: db,
	}

	dbPlaylistStore := &dbplaylist.Store{
		DB: db,
	}

	accountService := account.NewService(dbAccountStore)
	var accountEndpoints *accountsvc.Endpoints = accountsvc.NewEndpoints(accountService)

	playlistService := playlist.NewService(dbPlaylistStore)
	var playlistEndpoints *playlistsvc.Endpoints = playlistsvc.NewEndpoints(playlistService)

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

	var accountServer *accountsrv.Server = accountsrv.New(accountEndpoints, mux, dec, enc, nil, nil)
	accountsrv.Mount(mux, accountServer)

	var playlistServer *playlistsrv.Server = playlistsrv.New(playlistEndpoints, mux, dec, enc, nil, nil)
	playlistServer.Use(middleware.AuthenticateRequest())
	playlistsrv.Mount(mux, playlistServer)

	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	fmt.Printf("Account service has just started on %s ...\n", address)
	http.ListenAndServe(address, mux)
}
