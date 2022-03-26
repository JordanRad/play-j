package main

import (
	"fmt"
	"net/http"

	user "git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/account"
	dbuser "git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/db/dbaccount"
	accountsvc "git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/gen/account"
	accountsrv "git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/gen/http/account/server"
	goahttp "goa.design/goa/v3/http"

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

	dbStore := &dbuser.Store{
		DB: db,
	}

	userService := user.NewService(dbStore)

	var userEndpoints *accountsvc.Endpoints = accountsvc.NewEndpoints(userService)

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

	var userServer *accountsrv.Server = accountsrv.New(userEndpoints, mux, dec, enc, nil, nil)

	accountsrv.Mount(mux, userServer)

	fmt.Println("Account service has just started...")
	http.ListenAndServe("localhost:8091", mux)
}