// Code generated by goa v3.5.5, DO NOT EDIT.
//
// account HTTP client CLI support package
//
// Command:
// $ goa gen git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/design
// -o ./internal/

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	account "git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/gen/account"
)

// BuildRegisterPayload builds the payload for the account register endpoint
// from CLI flags.
func BuildRegisterPayload(accountRegisterBody string) (*account.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(accountRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"confirmedPassword\": \"Libero ut magnam.\",\n      \"email\": \"Dolorum reprehenderit repellendus praesentium sed neque incidunt.\",\n      \"password\": \"Rerum distinctio minus facilis consectetur facere.\"\n   }'")
		}
	}
	v := &account.RegisterPayload{
		Email:             body.Email,
		Password:          body.Password,
		ConfirmedPassword: body.ConfirmedPassword,
	}

	return v, nil
}

// BuildLoginPayload builds the payload for the account login endpoint from CLI
// flags.
func BuildLoginPayload(accountLoginBody string) (*account.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(accountLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Nemo provident eos quis ut ut ipsum.\",\n      \"password\": \"Molestiae voluptas dolorum et.\"\n   }'")
		}
	}
	v := &account.LoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}

// BuildGetUserPlaylistsPayload builds the payload for the account
// getUserPlaylists endpoint from CLI flags.
func BuildGetUserPlaylistsPayload(accountGetUserPlaylistsAccountID string, accountGetUserPlaylistsAuth string) (*account.GetUserPlaylistsPayload, error) {
	var err error
	var accountID uint
	{
		var v uint64
		v, err = strconv.ParseUint(accountGetUserPlaylistsAccountID, 10, 64)
		accountID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for accountID, must be UINT")
		}
	}
	var auth *string
	{
		if accountGetUserPlaylistsAuth != "" {
			auth = &accountGetUserPlaylistsAuth
		}
	}
	v := &account.GetUserPlaylistsPayload{}
	v.AccountID = &accountID
	v.Auth = auth

	return v, nil
}

// BuildCreateUserPlaylistPayload builds the payload for the account
// createUserPlaylist endpoint from CLI flags.
func BuildCreateUserPlaylistPayload(accountCreateUserPlaylistBody string, accountCreateUserPlaylistAccountID string, accountCreateUserPlaylistAuth string) (*account.CreateUserPlaylistPayload, error) {
	var err error
	var body CreateUserPlaylistRequestBody
	{
		err = json.Unmarshal([]byte(accountCreateUserPlaylistBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Quasi sint.\"\n   }'")
		}
	}
	var accountID uint
	{
		var v uint64
		v, err = strconv.ParseUint(accountCreateUserPlaylistAccountID, 10, 64)
		accountID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for accountID, must be UINT")
		}
	}
	var auth *string
	{
		if accountCreateUserPlaylistAuth != "" {
			auth = &accountCreateUserPlaylistAuth
		}
	}
	v := &account.CreateUserPlaylistPayload{
		Name: body.Name,
	}
	v.AccountID = &accountID
	v.Auth = auth

	return v, nil
}

// BuildDeleteUserPlaylistPayload builds the payload for the account
// deleteUserPlaylist endpoint from CLI flags.
func BuildDeleteUserPlaylistPayload(accountDeleteUserPlaylistAccountID string, accountDeleteUserPlaylistPlaylistID string, accountDeleteUserPlaylistAuth string) (*account.DeleteUserPlaylistPayload, error) {
	var err error
	var accountID uint
	{
		var v uint64
		v, err = strconv.ParseUint(accountDeleteUserPlaylistAccountID, 10, 64)
		accountID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for accountID, must be UINT")
		}
	}
	var playlistID uint
	{
		var v uint64
		v, err = strconv.ParseUint(accountDeleteUserPlaylistPlaylistID, 10, 64)
		playlistID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for playlistID, must be UINT")
		}
	}
	var auth *string
	{
		if accountDeleteUserPlaylistAuth != "" {
			auth = &accountDeleteUserPlaylistAuth
		}
	}
	v := &account.DeleteUserPlaylistPayload{}
	v.AccountID = &accountID
	v.PlaylistID = &playlistID
	v.Auth = auth

	return v, nil
}
