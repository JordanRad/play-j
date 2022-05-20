package player

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
)

// You only need **one** of these per package!
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Store
type Store interface {
	GetTrackByID(ctx context.Context, id uint) (*dbmodels.Track, error)
	Search(ctx context.Context, searchInput string, limit uint) ([]*dbmodels.PlayerSearch, error)
}

//counterfeiter:generate . CloudStorage
type CloudStorage interface {
	ReadFileFromFolder(ctx context.Context, filePath string) ([]byte, error)
}

type Service struct {
	store        Store
	cloudStorage CloudStorage
}

type SearchResponse struct {
	Total         uint
	SearchResults []*dbmodels.PlayerSearch
}

func NewService(store Store, cloudStorage CloudStorage) *Service {
	return &Service{
		store:        store,
		cloudStorage: cloudStorage,
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
