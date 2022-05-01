package player

import (
	"net/http"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
)

type SearchResponse struct {
	Total         uint
	SearchResults []*dbmodels.PlayerSearch
}

type PlayerService interface {
	StreamTrackByID(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}
