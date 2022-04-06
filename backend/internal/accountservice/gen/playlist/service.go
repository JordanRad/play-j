// Code generated by goa v3.7.0, DO NOT EDIT.
//
// playlist service
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package playlist

import (
	"context"
)

// Playlist operations
type Service interface {
	// GetAccountPlaylistCollection implements getAccountPlaylistCollection.
	GetAccountPlaylistCollection(context.Context, *GetAccountPlaylistCollectionPayload) (res *AccountPlaylistCollectionResponse, err error)
	// CreateAccountPlaylist implements createAccountPlaylist.
	CreateAccountPlaylist(context.Context, *CreateAccountPlaylistPayload) (res *PlaylistModificationResponse, err error)
	// RenameAccountPlaylist implements renameAccountPlaylist.
	RenameAccountPlaylist(context.Context, *RenameAccountPlaylistPayload) (res *PlaylistModificationResponse, err error)
	// DeleteAccountPlaylist implements deleteAccountPlaylist.
	DeleteAccountPlaylist(context.Context, *DeleteAccountPlaylistPayload) (res *PlaylistModificationResponse, err error)
	// GetAccountPlaylist implements getAccountPlaylist.
	GetAccountPlaylist(context.Context, *GetAccountPlaylistPayload) (res *AccountPlaylistResponse, err error)
	// AddTrackToAccountPlaylist implements addTrackToAccountPlaylist.
	AddTrackToAccountPlaylist(context.Context, *AddTrackToAccountPlaylistPayload) (res *PlaylistModificationResponse, err error)
	// RemoveTrackFromAccountPlaylist implements removeTrackFromAccountPlaylist.
	RemoveTrackFromAccountPlaylist(context.Context, *RemoveTrackFromAccountPlaylistPayload) (res *PlaylistModificationResponse, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "playlist"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [7]string{"getAccountPlaylistCollection", "createAccountPlaylist", "renameAccountPlaylist", "deleteAccountPlaylist", "getAccountPlaylist", "addTrackToAccountPlaylist", "removeTrackFromAccountPlaylist"}

// AccountPlaylistCollectionResponse is the result type of the playlist service
// getAccountPlaylistCollection method.
type AccountPlaylistCollectionResponse struct {
	// Number of resources
	Total int32
	// Operation completion status
	Resources []*AccountPlaylistResponse
}

// AccountPlaylistResponse is the result type of the playlist service
// getAccountPlaylist method.
type AccountPlaylistResponse struct {
	// Playlist id
	ID int32
	// Playlist name
	Name string
	// Array of TrackIDs
	TrackIDs []int32
}

// AddTrackToAccountPlaylistPayload is the payload type of the playlist service
// addTrackToAccountPlaylist method.
type AddTrackToAccountPlaylistPayload struct {
	// Playlist ID to modify
	PlaylistID *uint
	// Track ID to be added
	TrackID *uint
	// Authorization Header
	Auth *string
}

// CreateAccountPlaylistPayload is the payload type of the playlist service
// createAccountPlaylist method.
type CreateAccountPlaylistPayload struct {
	// Authorization Header
	Auth *string
	// Playlist name
	Name *string
}

// DeleteAccountPlaylistPayload is the payload type of the playlist service
// deleteAccountPlaylist method.
type DeleteAccountPlaylistPayload struct {
	Auth       *string
	PlaylistID *uint
}

// GetAccountPlaylistCollectionPayload is the payload type of the playlist
// service getAccountPlaylistCollection method.
type GetAccountPlaylistCollectionPayload struct {
	Auth *string
}

// GetAccountPlaylistPayload is the payload type of the playlist service
// getAccountPlaylist method.
type GetAccountPlaylistPayload struct {
	// Playlist ID
	PlaylistID *uint
	// Authorization Header
	Auth *string
}

// PlaylistModificationResponse is the result type of the playlist service
// createAccountPlaylist method.
type PlaylistModificationResponse struct {
	// Operation completion status
	Message string
}

// RemoveTrackFromAccountPlaylistPayload is the payload type of the playlist
// service removeTrackFromAccountPlaylist method.
type RemoveTrackFromAccountPlaylistPayload struct {
	// Playlist ID to modify
	PlaylistID *uint
	// Track ID to be deleted
	TrackID *uint
	// Authorization Header
	Auth *string
}

// RenameAccountPlaylistPayload is the payload type of the playlist service
// renameAccountPlaylist method.
type RenameAccountPlaylistPayload struct {
	// Authorization Header
	Auth *string
	// Playlist id to modify
	PlaylistID *uint
	// New playlist name
	Name *string
}
