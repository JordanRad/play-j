// Code generated by goa v3.5.5, DO NOT EDIT.
//
// HTTP request path constructors for the player service.
//
// Command:
// $ goa gen github.com/JordanRad/play-j/backend/internal/design/player-service
// -o ./internal/playerservice

package client

import (
	"fmt"
)

// GetTrackByIDPlayerPath returns the URL path to the player service getTrackByID HTTP endpoint.
func GetTrackByIDPlayerPath(trackID uint) string {
	return fmt.Sprintf("/api/v1/player-service/tracks/%v", trackID)
}