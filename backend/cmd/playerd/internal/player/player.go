package player

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
	"github.com/gorilla/mux"
)

type Store interface {
	GetTrackByID(ctx context.Context, id uint) (*dbmodels.Track, error)
	Search(ctx context.Context, searchInput string, limit uint) ([]*dbmodels.PlayerSearch, error)
}

type CloudStorage interface {
	ReadFileFromFolder(ctx context.Context, folderName string, fileName string) ([]byte, error)
}

type Service struct {
	store        Store
	cloudStorage CloudStorage
}

func NewService(store Store, cloudStorage CloudStorage) *Service {
	return &Service{
		store:        store,
		cloudStorage: cloudStorage,
	}
}

var _ PlayerService = (*Service)(nil)

func (s *Service) StreamTrackByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)

	trackID, err := strconv.Atoi(vars["trackID"])
	if err != nil {
		fmt.Printf("error getting the vars: %v", err)
		w.Write([]byte("Cannot find the file: "))
	}

	fmt.Println(trackID)

	track, err := s.store.GetTrackByID(ctx, uint(trackID))
	if err != nil {
		fmt.Printf("error getting the file location from the database: %v", err)
		w.Write([]byte("Cannot find the file: "))
	}

	fmt.Println(track)
	musicFile, err := s.cloudStorage.ReadFileFromFolder(ctx, "Metallica", "One.mp3")

	if err != nil {
		fmt.Printf("error getting the file: %v", err)
		w.Write([]byte("Cannot find the file: "))
	}

	fmt.Println(len(musicFile))

	flusher, ok := w.(http.Flusher)

	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	w.Header().Set("Content-Type", "audio/mpegurl")
	w.Header().Set("Transfer-encoding", "chunked")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	for i := 0; i < 5; i++ {
		start := (i * len(musicFile) / 5)
		end := ((i + 1) * len(musicFile)) / 5

		fmt.Printf("Chunk #%v sent...\n", i+1)
		w.Write(musicFile[start:end])
		flusher.Flush() // Trigger "chunked" encoding and send a chunk...
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Service) Search(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	limit := r.URL.Query().Get("limit")

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		fmt.Printf("error: %v\n", err)
		w.Write([]byte("limit query paramter has not been found"))
	}

	searchInput := r.URL.Query().Get("search_input")

	fmt.Printf("limit: %d, search_input: %s", limitInt, searchInput)

	databaseResults, err := s.store.Search(ctx, searchInput, uint(limitInt))

	if err != nil {
		fmt.Printf("error: %v\n", err)
		w.Write([]byte("db error"))
	}

	responseBody := &SearchResponse{
		Total:         1,
		SearchResults: databaseResults,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)
}
