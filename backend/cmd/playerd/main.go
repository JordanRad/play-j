package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	player "github.com/JordanRad/play-j/backend/cmd/playerd/internal/player"
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

	configFile, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("Config file cannot be read: %v", err)
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)

	var configuration map[string]interface{}

	json.Unmarshal([]byte(byteValue), &configuration)

	dbConnectionString := fmt.Sprintf("postgres://%v:%v@%v/%v", configuration["postgresql_user"], configuration["postgresql_password"], configuration["postgresql_host"], configuration["postgresql_db_name"])

	db, err := sql.Open("pgx", dbConnectionString)
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

	// dbAccountStore := &dbaccount.Store{
	// 	DB: db,
	// }

	// dbPlaylistStore := &dbplaylist.Store{
	// 	DB: db,
	// }

	playerService := player.NewService(nil, nil)
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
	playerServer.Use(middleware.InjectJWTInContext())

	fmt.Print("Player service has just started...\n")
	http.ListenAndServe("localhost:8092", mux)
}
