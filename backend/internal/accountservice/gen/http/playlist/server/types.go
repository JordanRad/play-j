// Code generated by goa v3.7.0, DO NOT EDIT.
//
// playlist HTTP server types
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package server

import (
	playlist "github.com/JordanRad/play-j/backend/internal/accountservice/gen/playlist"
)

// CreateAccountPlaylistRequestBody is the type of the "playlist" service
// "createAccountPlaylist" endpoint HTTP request body.
type CreateAccountPlaylistRequestBody struct {
	// Playlist name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// RenameAccountPlaylistRequestBody is the type of the "playlist" service
// "renameAccountPlaylist" endpoint HTTP request body.
type RenameAccountPlaylistRequestBody struct {
	// New playlist name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// GetAccountPlaylistCollectionResponseBody is the type of the "playlist"
// service "getAccountPlaylistCollection" endpoint HTTP response body.
type GetAccountPlaylistCollectionResponseBody struct {
	// Number of resources
	Total int32 `form:"total" json:"total" xml:"total"`
	// Operation completion status
	Resources []*AccountPlaylistResponseResponseBody `form:"resources" json:"resources" xml:"resources"`
}

// CreateAccountPlaylistResponseBody is the type of the "playlist" service
// "createAccountPlaylist" endpoint HTTP response body.
type CreateAccountPlaylistResponseBody struct {
	// Operation completion status
	Message string `form:"message" json:"message" xml:"message"`
}

// RenameAccountPlaylistResponseBody is the type of the "playlist" service
// "renameAccountPlaylist" endpoint HTTP response body.
type RenameAccountPlaylistResponseBody struct {
	// Operation completion status
	Message string `form:"message" json:"message" xml:"message"`
}

// DeleteAccountPlaylistResponseBody is the type of the "playlist" service
// "deleteAccountPlaylist" endpoint HTTP response body.
type DeleteAccountPlaylistResponseBody struct {
	// Operation completion status
	Message string `form:"message" json:"message" xml:"message"`
}

// GetAccountPlaylistResponseBody is the type of the "playlist" service
// "getAccountPlaylist" endpoint HTTP response body.
type GetAccountPlaylistResponseBody struct {
	// Playlist id
	ID int32 `form:"id" json:"id" xml:"id"`
	// Playlist name
	Name string `form:"name" json:"name" xml:"name"`
	// Array of TrackIDs
	TrackIDs []int32 `form:"trackIDs" json:"trackIDs" xml:"trackIDs"`
}

// AddTrackToAccountPlaylistResponseBody is the type of the "playlist" service
// "addTrackToAccountPlaylist" endpoint HTTP response body.
type AddTrackToAccountPlaylistResponseBody struct {
	// Operation completion status
	Message string `form:"message" json:"message" xml:"message"`
}

// RemoveTrackFromAccountPlaylistResponseBody is the type of the "playlist"
// service "removeTrackFromAccountPlaylist" endpoint HTTP response body.
type RemoveTrackFromAccountPlaylistResponseBody struct {
	// Operation completion status
	Message string `form:"message" json:"message" xml:"message"`
}

// AccountPlaylistResponseResponseBody is used to define fields on response
// body types.
type AccountPlaylistResponseResponseBody struct {
	// Playlist id
	ID int32 `form:"id" json:"id" xml:"id"`
	// Playlist name
	Name string `form:"name" json:"name" xml:"name"`
	// Array of TrackIDs
	TrackIDs []int32 `form:"trackIDs" json:"trackIDs" xml:"trackIDs"`
}

// NewGetAccountPlaylistCollectionResponseBody builds the HTTP response body
// from the result of the "getAccountPlaylistCollection" endpoint of the
// "playlist" service.
func NewGetAccountPlaylistCollectionResponseBody(res *playlist.AccountPlaylistCollectionResponse) *GetAccountPlaylistCollectionResponseBody {
	body := &GetAccountPlaylistCollectionResponseBody{
		Total: res.Total,
	}
	if res.Resources != nil {
		body.Resources = make([]*AccountPlaylistResponseResponseBody, len(res.Resources))
		for i, val := range res.Resources {
			body.Resources[i] = marshalPlaylistAccountPlaylistResponseToAccountPlaylistResponseResponseBody(val)
		}
	}
	return body
}

// NewCreateAccountPlaylistResponseBody builds the HTTP response body from the
// result of the "createAccountPlaylist" endpoint of the "playlist" service.
func NewCreateAccountPlaylistResponseBody(res *playlist.PlaylistModificationResponse) *CreateAccountPlaylistResponseBody {
	body := &CreateAccountPlaylistResponseBody{
		Message: res.Message,
	}
	return body
}

// NewRenameAccountPlaylistResponseBody builds the HTTP response body from the
// result of the "renameAccountPlaylist" endpoint of the "playlist" service.
func NewRenameAccountPlaylistResponseBody(res *playlist.PlaylistModificationResponse) *RenameAccountPlaylistResponseBody {
	body := &RenameAccountPlaylistResponseBody{
		Message: res.Message,
	}
	return body
}

// NewDeleteAccountPlaylistResponseBody builds the HTTP response body from the
// result of the "deleteAccountPlaylist" endpoint of the "playlist" service.
func NewDeleteAccountPlaylistResponseBody(res *playlist.PlaylistModificationResponse) *DeleteAccountPlaylistResponseBody {
	body := &DeleteAccountPlaylistResponseBody{
		Message: res.Message,
	}
	return body
}

// NewGetAccountPlaylistResponseBody builds the HTTP response body from the
// result of the "getAccountPlaylist" endpoint of the "playlist" service.
func NewGetAccountPlaylistResponseBody(res *playlist.AccountPlaylistResponse) *GetAccountPlaylistResponseBody {
	body := &GetAccountPlaylistResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	if res.TrackIDs != nil {
		body.TrackIDs = make([]int32, len(res.TrackIDs))
		for i, val := range res.TrackIDs {
			body.TrackIDs[i] = val
		}
	}
	return body
}

// NewAddTrackToAccountPlaylistResponseBody builds the HTTP response body from
// the result of the "addTrackToAccountPlaylist" endpoint of the "playlist"
// service.
func NewAddTrackToAccountPlaylistResponseBody(res *playlist.PlaylistModificationResponse) *AddTrackToAccountPlaylistResponseBody {
	body := &AddTrackToAccountPlaylistResponseBody{
		Message: res.Message,
	}
	return body
}

// NewRemoveTrackFromAccountPlaylistResponseBody builds the HTTP response body
// from the result of the "removeTrackFromAccountPlaylist" endpoint of the
// "playlist" service.
func NewRemoveTrackFromAccountPlaylistResponseBody(res *playlist.PlaylistModificationResponse) *RemoveTrackFromAccountPlaylistResponseBody {
	body := &RemoveTrackFromAccountPlaylistResponseBody{
		Message: res.Message,
	}
	return body
}

// NewGetAccountPlaylistCollectionPayload builds a playlist service
// getAccountPlaylistCollection endpoint payload.
func NewGetAccountPlaylistCollectionPayload(auth *string) *playlist.GetAccountPlaylistCollectionPayload {
	v := &playlist.GetAccountPlaylistCollectionPayload{}
	v.Auth = auth

	return v
}

// NewCreateAccountPlaylistPayload builds a playlist service
// createAccountPlaylist endpoint payload.
func NewCreateAccountPlaylistPayload(body *CreateAccountPlaylistRequestBody, auth *string) *playlist.CreateAccountPlaylistPayload {
	v := &playlist.CreateAccountPlaylistPayload{
		Name: body.Name,
	}
	v.Auth = auth

	return v
}

// NewRenameAccountPlaylistPayload builds a playlist service
// renameAccountPlaylist endpoint payload.
func NewRenameAccountPlaylistPayload(body *RenameAccountPlaylistRequestBody, playlistID uint, auth *string) *playlist.RenameAccountPlaylistPayload {
	v := &playlist.RenameAccountPlaylistPayload{
		Name: body.Name,
	}
	v.PlaylistID = &playlistID
	v.Auth = auth

	return v
}

// NewDeleteAccountPlaylistPayload builds a playlist service
// deleteAccountPlaylist endpoint payload.
func NewDeleteAccountPlaylistPayload(playlistID uint, auth *string) *playlist.DeleteAccountPlaylistPayload {
	v := &playlist.DeleteAccountPlaylistPayload{}
	v.PlaylistID = &playlistID
	v.Auth = auth

	return v
}

// NewGetAccountPlaylistPayload builds a playlist service getAccountPlaylist
// endpoint payload.
func NewGetAccountPlaylistPayload(playlistID uint, auth *string) *playlist.GetAccountPlaylistPayload {
	v := &playlist.GetAccountPlaylistPayload{}
	v.PlaylistID = &playlistID
	v.Auth = auth

	return v
}

// NewAddTrackToAccountPlaylistPayload builds a playlist service
// addTrackToAccountPlaylist endpoint payload.
func NewAddTrackToAccountPlaylistPayload(playlistID uint, trackID uint, auth *string) *playlist.AddTrackToAccountPlaylistPayload {
	v := &playlist.AddTrackToAccountPlaylistPayload{}
	v.PlaylistID = &playlistID
	v.TrackID = &trackID
	v.Auth = auth

	return v
}

// NewRemoveTrackFromAccountPlaylistPayload builds a playlist service
// removeTrackFromAccountPlaylist endpoint payload.
func NewRemoveTrackFromAccountPlaylistPayload(playlistID uint, trackID uint, auth *string) *playlist.RemoveTrackFromAccountPlaylistPayload {
	v := &playlist.RemoveTrackFromAccountPlaylistPayload{}
	v.PlaylistID = &playlistID
	v.TrackID = &trackID
	v.Auth = auth

	return v
}