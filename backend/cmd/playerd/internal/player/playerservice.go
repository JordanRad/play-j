package player

import (
	"net/http"
)

type SearchResult struct {
	TrackID uint
	AlbumID uint
	Artist  uint
	Result  string
}

type SearchResponse struct {
	Total         uint
	SearchResults []SearchResult
}

type PlayerService interface {
	StreamTrackByID(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}
