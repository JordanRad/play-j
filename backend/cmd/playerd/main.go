package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbplayer"
	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/player"
	"github.com/JordanRad/play-j/backend/internal/cloudstorage"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
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

	dbPlayerStore := &dbplayer.Store{
		DB: db,
	}

	musicStorage := cloudstorage.NewCloudStorage("")

	playerService := player.NewService(dbPlayerStore, musicStorage)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/player-service/audio/{trackID}", playerService.StreamTrackByID)
	router.HandleFunc("/api/v1/player-service/search", playerService.Search)
	http.Handle("/api/v1/player-service/", router)

	address := fmt.Sprintf("%s:%d", config.HTTP.Host, config.HTTP.Port)
	fmt.Printf("Player streaming service has just started on %s ...\n", address)
	http.ListenAndServe(address, router)
}
