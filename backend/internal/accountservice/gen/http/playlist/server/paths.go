// Code generated by goa v3.7.0, DO NOT EDIT.
//
// HTTP request path constructors for the playlist service.
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package server

import (
	"fmt"
)

// GetAccountPlaylistCollectionPlaylistPath returns the URL path to the playlist service getAccountPlaylistCollection HTTP endpoint.
func GetAccountPlaylistCollectionPlaylistPath() string {
	return "/api/v1/account-service/playlists"
}

// CreateAccountPlaylistPlaylistPath returns the URL path to the playlist service createAccountPlaylist HTTP endpoint.
func CreateAccountPlaylistPlaylistPath() string {
	return "/api/v1/account-service/playlists"
}

// RenameAccountPlaylistPlaylistPath returns the URL path to the playlist service renameAccountPlaylist HTTP endpoint.
func RenameAccountPlaylistPlaylistPath(playlistID uint) string {
	return fmt.Sprintf("/api/v1/account-service/playlists/%v", playlistID)
}

// DeleteAccountPlaylistPlaylistPath returns the URL path to the playlist service deleteAccountPlaylist HTTP endpoint.
func DeleteAccountPlaylistPlaylistPath(playlistID uint) string {
	return fmt.Sprintf("/api/v1/account-service/playlists/%v", playlistID)
}

// GetAccountPlaylistPlaylistPath returns the URL path to the playlist service getAccountPlaylist HTTP endpoint.
func GetAccountPlaylistPlaylistPath(playlistID uint) string {
	return fmt.Sprintf("/api/v1/account-service/playlists/%v", playlistID)
}

// AddTrackToAccountPlaylistPlaylistPath returns the URL path to the playlist service addTrackToAccountPlaylist HTTP endpoint.
func AddTrackToAccountPlaylistPlaylistPath(playlistID uint, trackID uint) string {
	return fmt.Sprintf("/api/v1/account-service/playlists/%v/tracks/%v", playlistID, trackID)
}

// RemoveTrackFromAccountPlaylistPlaylistPath returns the URL path to the playlist service removeTrackFromAccountPlaylist HTTP endpoint.
func RemoveTrackFromAccountPlaylistPlaylistPath(playlistID uint, trackID uint) string {
	return fmt.Sprintf("/api/v1/account-service/playlists/%v/tracks/%v", playlistID, trackID)
}
