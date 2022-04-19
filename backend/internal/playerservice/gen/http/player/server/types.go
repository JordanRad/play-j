// Code generated by goa v3.7.0, DO NOT EDIT.
//
// player HTTP server types
//
// Command:
// $ goa gen github.com/JordanRad/play-j/backend/internal/design/player-service
// -o ./internal/playerservice

package server

import (
	player "github.com/JordanRad/play-j/backend/internal/playerservice/gen/player"
	playerviews "github.com/JordanRad/play-j/backend/internal/playerservice/gen/player/views"
)

// GetTrackByIDResponseBody is the type of the "player" service "getTrackByID"
// endpoint HTTP response body.
type GetTrackByIDResponseBody struct {
	File []byte `form:"file,omitempty" json:"file,omitempty" xml:"file,omitempty"`
}

// NewGetTrackByIDResponseBody builds the HTTP response body from the result of
// the "getTrackByID" endpoint of the "player" service.
func NewGetTrackByIDResponseBody(res *playerviews.MusicFileResponseView) *GetTrackByIDResponseBody {
	body := &GetTrackByIDResponseBody{
		File: res.File,
	}
	return body
}

// NewGetTrackByIDPayload builds a player service getTrackByID endpoint payload.
func NewGetTrackByIDPayload(trackID uint) *player.GetTrackByIDPayload {
	v := &player.GetTrackByIDPayload{}
	v.TrackID = trackID

	return v
}
