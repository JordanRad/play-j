// Code generated by goa v3.5.5, DO NOT EDIT.
//
// HTTP request path constructors for the account service.
//
// Command:
// $ goa gen github.com/JordanRad/play-j/backend/internal/design -o ./internal/

package client

import (
	"fmt"
)

// RegisterAccountPath returns the URL path to the account service register HTTP endpoint.
func RegisterAccountPath() string {
	return "/api/v1/account/register"
}

// LoginAccountPath returns the URL path to the account service login HTTP endpoint.
func LoginAccountPath() string {
	return "/api/v1/account/login"
}

// GetAccountPlaylistCollectionAccountPath returns the URL path to the account service getAccountPlaylistCollection HTTP endpoint.
func GetAccountPlaylistCollectionAccountPath(accountID uint) string {
	return fmt.Sprintf("/api/v1/account/%v/playlists", accountID)
}

// CreateAccountPlaylistAccountPath returns the URL path to the account service createAccountPlaylist HTTP endpoint.
func CreateAccountPlaylistAccountPath(accountID uint) string {
	return fmt.Sprintf("/api/v1/account/%v/playlists", accountID)
}

// DeleteAccountPlaylistAccountPath returns the URL path to the account service deleteAccountPlaylist HTTP endpoint.
func DeleteAccountPlaylistAccountPath(accountID uint, playlistID uint) string {
	return fmt.Sprintf("/api/v1/account/%v/playlists/%v", accountID, playlistID)
}

// GetAccountPlaylistAccountPath returns the URL path to the account service getAccountPlaylist HTTP endpoint.
func GetAccountPlaylistAccountPath(accountID uint, playlistID uint) string {
	return fmt.Sprintf("/api/v1/account/%v/playlists/%v", accountID, playlistID)
}
