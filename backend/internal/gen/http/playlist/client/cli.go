// Code generated by goa v3.5.5, DO NOT EDIT.
//
// playlist HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	playlist "github.com/JordanRad/play-j/backend/internal/gen/playlist"
)

// BuildGetAccountPlaylistCollectionPayload builds the payload for the playlist
// getAccountPlaylistCollection endpoint from CLI flags.
func BuildGetAccountPlaylistCollectionPayload(playlistGetAccountPlaylistCollectionBody string, playlistGetAccountPlaylistCollectionAuth string) (*playlist.GetAccountPlaylistCollectionPayload, error) {
	var err error
	var body GetAccountPlaylistCollectionRequestBody
	{
		err = json.Unmarshal([]byte(playlistGetAccountPlaylistCollectionBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"accountID\": 11324780232132348642\n   }'")
		}
	}
	var auth *string
	{
		if playlistGetAccountPlaylistCollectionAuth != "" {
			auth = &playlistGetAccountPlaylistCollectionAuth
		}
	}
	v := &playlist.GetAccountPlaylistCollectionPayload{
		AccountID: body.AccountID,
	}
	v.Auth = auth

	return v, nil
}

// BuildCreateAccountPlaylistPayload builds the payload for the playlist
// createAccountPlaylist endpoint from CLI flags.
func BuildCreateAccountPlaylistPayload(playlistCreateAccountPlaylistBody string, playlistCreateAccountPlaylistAuth string) (*playlist.CreateAccountPlaylistPayload, error) {
	var err error
	var body CreateAccountPlaylistRequestBody
	{
		err = json.Unmarshal([]byte(playlistCreateAccountPlaylistBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Et enim consectetur.\"\n   }'")
		}
	}
	var auth *string
	{
		if playlistCreateAccountPlaylistAuth != "" {
			auth = &playlistCreateAccountPlaylistAuth
		}
	}
	v := &playlist.CreateAccountPlaylistPayload{
		Name: body.Name,
	}
	v.Auth = auth

	return v, nil
}

// BuildDeleteAccountPlaylistPayload builds the payload for the playlist
// deleteAccountPlaylist endpoint from CLI flags.
func BuildDeleteAccountPlaylistPayload(playlistDeleteAccountPlaylistBody string, playlistDeleteAccountPlaylistAuth string) (*playlist.DeleteAccountPlaylistPayload, error) {
	var err error
	var body DeleteAccountPlaylistRequestBody
	{
		err = json.Unmarshal([]byte(playlistDeleteAccountPlaylistBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"playlistID\": 1690349288743571665\n   }'")
		}
	}
	var auth *string
	{
		if playlistDeleteAccountPlaylistAuth != "" {
			auth = &playlistDeleteAccountPlaylistAuth
		}
	}
	v := &playlist.DeleteAccountPlaylistPayload{
		PlaylistID: body.PlaylistID,
	}
	v.Auth = auth

	return v, nil
}

// BuildGetAccountPlaylistPayload builds the payload for the playlist
// getAccountPlaylist endpoint from CLI flags.
func BuildGetAccountPlaylistPayload(playlistGetAccountPlaylistPlaylistID string, playlistGetAccountPlaylistAuth string) (*playlist.GetAccountPlaylistPayload, error) {
	var err error
	var playlistID uint
	{
		var v uint64
		v, err = strconv.ParseUint(playlistGetAccountPlaylistPlaylistID, 10, 64)
		playlistID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for playlistID, must be UINT")
		}
	}
	var auth *string
	{
		if playlistGetAccountPlaylistAuth != "" {
			auth = &playlistGetAccountPlaylistAuth
		}
	}
	v := &playlist.GetAccountPlaylistPayload{}
	v.PlaylistID = &playlistID
	v.Auth = auth

	return v, nil
}

// BuildAddTrackToAccountPlaylistPayload builds the payload for the playlist
// addTrackToAccountPlaylist endpoint from CLI flags.
func BuildAddTrackToAccountPlaylistPayload(playlistAddTrackToAccountPlaylistPlaylistID string, playlistAddTrackToAccountPlaylistTrackID string, playlistAddTrackToAccountPlaylistAuth string) (*playlist.AddTrackToAccountPlaylistPayload, error) {
	var err error
	var playlistID uint
	{
		var v uint64
		v, err = strconv.ParseUint(playlistAddTrackToAccountPlaylistPlaylistID, 10, 64)
		playlistID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for playlistID, must be UINT")
		}
	}
	var trackID uint
	{
		var v uint64
		v, err = strconv.ParseUint(playlistAddTrackToAccountPlaylistTrackID, 10, 64)
		trackID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for trackID, must be UINT")
		}
	}
	var auth *string
	{
		if playlistAddTrackToAccountPlaylistAuth != "" {
			auth = &playlistAddTrackToAccountPlaylistAuth
		}
	}
	v := &playlist.AddTrackToAccountPlaylistPayload{}
	v.PlaylistID = &playlistID
	v.TrackID = &trackID
	v.Auth = auth

	return v, nil
}

// BuildRemoveTrackFromAccountPlaylistPayload builds the payload for the
// playlist removeTrackFromAccountPlaylist endpoint from CLI flags.
func BuildRemoveTrackFromAccountPlaylistPayload(playlistRemoveTrackFromAccountPlaylistPlaylistID string, playlistRemoveTrackFromAccountPlaylistTrackID string, playlistRemoveTrackFromAccountPlaylistAuth string) (*playlist.RemoveTrackFromAccountPlaylistPayload, error) {
	var err error
	var playlistID uint
	{
		var v uint64
		v, err = strconv.ParseUint(playlistRemoveTrackFromAccountPlaylistPlaylistID, 10, 64)
		playlistID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for playlistID, must be UINT")
		}
	}
	var trackID uint
	{
		var v uint64
		v, err = strconv.ParseUint(playlistRemoveTrackFromAccountPlaylistTrackID, 10, 64)
		trackID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for trackID, must be UINT")
		}
	}
	var auth *string
	{
		if playlistRemoveTrackFromAccountPlaylistAuth != "" {
			auth = &playlistRemoveTrackFromAccountPlaylistAuth
		}
	}
	v := &playlist.RemoveTrackFromAccountPlaylistPayload{}
	v.PlaylistID = &playlistID
	v.TrackID = &trackID
	v.Auth = auth

	return v, nil
}
