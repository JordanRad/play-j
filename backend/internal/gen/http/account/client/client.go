// Code generated by goa v3.5.5, DO NOT EDIT.
//
// account client HTTP transport
//
// Command:
// $ goa gen git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/design
// -o ./internal/

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the account service endpoint HTTP clients.
type Client struct {
	// Register Doer is the HTTP client used to make requests to the register
	// endpoint.
	RegisterDoer goahttp.Doer

	// Login Doer is the HTTP client used to make requests to the login endpoint.
	LoginDoer goahttp.Doer

	// GetUserPlaylists Doer is the HTTP client used to make requests to the
	// getUserPlaylists endpoint.
	GetUserPlaylistsDoer goahttp.Doer

	// CreateUserPlaylist Doer is the HTTP client used to make requests to the
	// createUserPlaylist endpoint.
	CreateUserPlaylistDoer goahttp.Doer

	// DeleteUserPlaylist Doer is the HTTP client used to make requests to the
	// deleteUserPlaylist endpoint.
	DeleteUserPlaylistDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the account service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RegisterDoer:           doer,
		LoginDoer:              doer,
		GetUserPlaylistsDoer:   doer,
		CreateUserPlaylistDoer: doer,
		DeleteUserPlaylistDoer: doer,
		RestoreResponseBody:    restoreBody,
		scheme:                 scheme,
		host:                   host,
		decoder:                dec,
		encoder:                enc,
	}
}

// Register returns an endpoint that makes HTTP requests to the account service
// register server.
func (c *Client) Register() goa.Endpoint {
	var (
		encodeRequest  = EncodeRegisterRequest(c.encoder)
		decodeResponse = DecodeRegisterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRegisterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RegisterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "register", err)
		}
		return decodeResponse(resp)
	}
}

// Login returns an endpoint that makes HTTP requests to the account service
// login server.
func (c *Client) Login() goa.Endpoint {
	var (
		encodeRequest  = EncodeLoginRequest(c.encoder)
		decodeResponse = DecodeLoginResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildLoginRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.LoginDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "login", err)
		}
		return decodeResponse(resp)
	}
}

// GetUserPlaylists returns an endpoint that makes HTTP requests to the account
// service getUserPlaylists server.
func (c *Client) GetUserPlaylists() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetUserPlaylistsRequest(c.encoder)
		decodeResponse = DecodeGetUserPlaylistsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetUserPlaylistsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetUserPlaylistsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "getUserPlaylists", err)
		}
		return decodeResponse(resp)
	}
}

// CreateUserPlaylist returns an endpoint that makes HTTP requests to the
// account service createUserPlaylist server.
func (c *Client) CreateUserPlaylist() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateUserPlaylistRequest(c.encoder)
		decodeResponse = DecodeCreateUserPlaylistResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateUserPlaylistRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateUserPlaylistDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "createUserPlaylist", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteUserPlaylist returns an endpoint that makes HTTP requests to the
// account service deleteUserPlaylist server.
func (c *Client) DeleteUserPlaylist() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteUserPlaylistRequest(c.encoder)
		decodeResponse = DecodeDeleteUserPlaylistResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteUserPlaylistRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteUserPlaylistDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "deleteUserPlaylist", err)
		}
		return decodeResponse(resp)
	}
}
