package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbpayments"
	"github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/db/dbsubscriptions"
	payment "github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/payment"
	"github.com/JordanRad/play-j/backend/internal/middleware"
	paymentsrv "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/http/payment/server"
	paymentsvc "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/payment"

	subscription "github.com/JordanRad/play-j/backend/cmd/paymentsd/internal/subscription"
	subscriptionsrv "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/http/subscription/server"
	subscriptionsvc "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/subscription"
	_ "github.com/jackc/pgx/v4/stdlib"

	goahttp "goa.design/goa/v3/http"
)

func connectDB(config *config) *sql.DB {
	dbConnectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DBName)
	fmt.Println(dbConnectionString)
	var db *sql.DB
	db, err := sql.Open("pgx", dbConnectionString)
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
	return db
}

func main() {
	config, err := configFromEnv()

	if err != nil {
		log.Fatalf("Config file cannot be read: %v", err)
	}

	db := connectDB(config)

	paymentStore := &dbpayments.Store{
		DB: db,
	}

	subscriptionStore := &dbsubscriptions.Store{
		DB: db,
	}

	paymentService := payment.NewService(paymentStore, subscriptionStore)
	var paymentEndpoints *paymentsvc.Endpoints = paymentsvc.NewEndpoints(paymentService)

	subscriptionService := subscription.NewService(subscriptionStore)
	var subscriptionEndpoints *subscriptionsvc.Endpoints = subscriptionsvc.NewEndpoints(subscriptionService)

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

	var paymentServer *paymentsrv.Server = paymentsrv.New(paymentEndpoints, mux, dec, enc, nil, nil)
	paymentServer.Use(middleware.AuthenticateRequest())
	paymentsrv.Mount(mux, paymentServer)

	var subscriptionServer *subscriptionsrv.Server = subscriptionsrv.New(subscriptionEndpoints, mux, dec, enc, nil, nil)
	subscriptionServer.Use(middleware.AuthenticateRequest())
	subscriptionsrv.Mount(mux, subscriptionServer)

	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	fmt.Printf("Payment service has just started on %s ...\n", address)
	http.ListenAndServe(address, mux)
}
