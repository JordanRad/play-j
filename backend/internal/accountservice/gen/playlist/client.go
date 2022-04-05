// Code generated by goa v3.7.0, DO NOT EDIT.
//
// playlist client
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package playlist

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "playlist" service client.
type Client struct {
	GetAccountPlaylistCollectionEndpoint   goa.Endpoint
	CreateAccountPlaylistEndpoint          goa.Endpoint
	RenameAccountPlaylistEndpoint          goa.Endpoint
	DeleteAccountPlaylistEndpoint          goa.Endpoint
	GetAccountPlaylistEndpoint             goa.Endpoint
	AddTrackToAccountPlaylistEndpoint      goa.Endpoint
	RemoveTrackFromAccountPlaylistEndpoint goa.Endpoint
}

// NewClient initializes a "playlist" service client given the endpoints.
func NewClient(getAccountPlaylistCollection, createAccountPlaylist, renameAccountPlaylist, deleteAccountPlaylist, getAccountPlaylist, addTrackToAccountPlaylist, removeTrackFromAccountPlaylist goa.Endpoint) *Client {
	return &Client{
		GetAccountPlaylistCollectionEndpoint:   getAccountPlaylistCollection,
		CreateAccountPlaylistEndpoint:          createAccountPlaylist,
		RenameAccountPlaylistEndpoint:          renameAccountPlaylist,
		DeleteAccountPlaylistEndpoint:          deleteAccountPlaylist,
		GetAccountPlaylistEndpoint:             getAccountPlaylist,
		AddTrackToAccountPlaylistEndpoint:      addTrackToAccountPlaylist,
		RemoveTrackFromAccountPlaylistEndpoint: removeTrackFromAccountPlaylist,
	}
}

// GetAccountPlaylistCollection calls the "getAccountPlaylistCollection"
// endpoint of the "playlist" service.
func (c *Client) GetAccountPlaylistCollection(ctx context.Context, p *GetAccountPlaylistCollectionPayload) (res *AccountPlaylistCollectionResponse, err error) {
	var ires interface{}
	ires, err = c.GetAccountPlaylistCollectionEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*AccountPlaylistCollectionResponse), nil
}

// CreateAccountPlaylist calls the "createAccountPlaylist" endpoint of the
// "playlist" service.
func (c *Client) CreateAccountPlaylist(ctx context.Context, p *CreateAccountPlaylistPayload) (res *PlaylistModificationResponse, err error) {
	var ires interface{}
	ires, err = c.CreateAccountPlaylistEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PlaylistModificationResponse), nil
}

// RenameAccountPlaylist calls the "renameAccountPlaylist" endpoint of the
// "playlist" service.
func (c *Client) RenameAccountPlaylist(ctx context.Context, p *RenameAccountPlaylistPayload) (res *PlaylistModificationResponse, err error) {
	var ires interface{}
	ires, err = c.RenameAccountPlaylistEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PlaylistModificationResponse), nil
}

// DeleteAccountPlaylist calls the "deleteAccountPlaylist" endpoint of the
// "playlist" service.
func (c *Client) DeleteAccountPlaylist(ctx context.Context, p *DeleteAccountPlaylistPayload) (res *PlaylistModificationResponse, err error) {
	var ires interface{}
	ires, err = c.DeleteAccountPlaylistEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PlaylistModificationResponse), nil
}

// GetAccountPlaylist calls the "getAccountPlaylist" endpoint of the "playlist"
// service.
func (c *Client) GetAccountPlaylist(ctx context.Context, p *GetAccountPlaylistPayload) (res *AccountPlaylistResponse, err error) {
	var ires interface{}
	ires, err = c.GetAccountPlaylistEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*AccountPlaylistResponse), nil
}

// AddTrackToAccountPlaylist calls the "addTrackToAccountPlaylist" endpoint of
// the "playlist" service.
func (c *Client) AddTrackToAccountPlaylist(ctx context.Context, p *AddTrackToAccountPlaylistPayload) (res *PlaylistModificationResponse, err error) {
	var ires interface{}
	ires, err = c.AddTrackToAccountPlaylistEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PlaylistModificationResponse), nil
}

// RemoveTrackFromAccountPlaylist calls the "removeTrackFromAccountPlaylist"
// endpoint of the "playlist" service.
func (c *Client) RemoveTrackFromAccountPlaylist(ctx context.Context, p *RemoveTrackFromAccountPlaylistPayload) (res *PlaylistModificationResponse, err error) {
	var ires interface{}
	ires, err = c.RemoveTrackFromAccountPlaylistEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*PlaylistModificationResponse), nil
}
