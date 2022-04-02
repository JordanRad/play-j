package main

import (
	"fmt"
	"net/http"

	account "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/account"
	dbaccount "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbaccount"
	dbplaylist "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbplaylist"
	playlist "github.com/JordanRad/play-j/backend/cmd/accountsd/internal/playlist"
	accountsvc "github.com/JordanRad/play-j/backend/internal/gen/account"
	accountsrv "github.com/JordanRad/play-j/backend/internal/gen/http/account/server"
	playlistsrv "github.com/JordanRad/play-j/backend/internal/gen/http/playlist/server"
	playlistsvc "github.com/JordanRad/play-j/backend/internal/gen/playlist"
	goahttp "goa.design/goa/v3/http"

	"github.com/JordanRad/play-j/backend/internal/middleware"

	"database/sql"
	"time"

	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://playj:playj1307@localhost:5433/playj-accounts-db")
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(50)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	if err := db.Ping(); err != nil {
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
	playlistServer.Use(middleware.InjectJWTInContext())
	playlistsrv.Mount(mux, playlistServer)

	fmt.Print("Account service has just started...\n")
	http.ListenAndServe("localhost:8091", mux)
}
