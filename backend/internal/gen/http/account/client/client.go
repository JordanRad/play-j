// Code generated by goa v3.5.5, DO NOT EDIT.
//
// account client HTTP transport
//
// Command:
// $ goa gen github.com/JordanRad/play-j/backend/internal/design -o ./internal/

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

	// GetAccountPlaylistCollection Doer is the HTTP client used to make requests
	// to the getAccountPlaylistCollection endpoint.
	GetAccountPlaylistCollectionDoer goahttp.Doer

	// CreateAccountPlaylist Doer is the HTTP client used to make requests to the
	// createAccountPlaylist endpoint.
	CreateAccountPlaylistDoer goahttp.Doer

	// DeleteAccountPlaylist Doer is the HTTP client used to make requests to the
	// deleteAccountPlaylist endpoint.
	DeleteAccountPlaylistDoer goahttp.Doer

	// GetAccountPlaylist Doer is the HTTP client used to make requests to the
	// getAccountPlaylist endpoint.
	GetAccountPlaylistDoer goahttp.Doer

	// AddTrackToAccountPlaylist Doer is the HTTP client used to make requests to
	// the addTrackToAccountPlaylist endpoint.
	AddTrackToAccountPlaylistDoer goahttp.Doer

	// RemoveTrackFromAccountPlaylist Doer is the HTTP client used to make requests
	// to the removeTrackFromAccountPlaylist endpoint.
	RemoveTrackFromAccountPlaylistDoer goahttp.Doer

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
		RegisterDoer:                       doer,
		LoginDoer:                          doer,
		GetAccountPlaylistCollectionDoer:   doer,
		CreateAccountPlaylistDoer:          doer,
		DeleteAccountPlaylistDoer:          doer,
		GetAccountPlaylistDoer:             doer,
		AddTrackToAccountPlaylistDoer:      doer,
		RemoveTrackFromAccountPlaylistDoer: doer,
		RestoreResponseBody:                restoreBody,
		scheme:                             scheme,
		host:                               host,
		decoder:                            dec,
		encoder:                            enc,
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

// GetAccountPlaylistCollection returns an endpoint that makes HTTP requests to
// the account service getAccountPlaylistCollection server.
func (c *Client) GetAccountPlaylistCollection() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetAccountPlaylistCollectionRequest(c.encoder)
		decodeResponse = DecodeGetAccountPlaylistCollectionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetAccountPlaylistCollectionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAccountPlaylistCollectionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "getAccountPlaylistCollection", err)
		}
		return decodeResponse(resp)
	}
}

// CreateAccountPlaylist returns an endpoint that makes HTTP requests to the
// account service createAccountPlaylist server.
func (c *Client) CreateAccountPlaylist() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateAccountPlaylistRequest(c.encoder)
		decodeResponse = DecodeCreateAccountPlaylistResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateAccountPlaylistRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateAccountPlaylistDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "createAccountPlaylist", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteAccountPlaylist returns an endpoint that makes HTTP requests to the
// account service deleteAccountPlaylist server.
func (c *Client) DeleteAccountPlaylist() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteAccountPlaylistRequest(c.encoder)
		decodeResponse = DecodeDeleteAccountPlaylistResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteAccountPlaylistRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteAccountPlaylistDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "deleteAccountPlaylist", err)
		}
		return decodeResponse(resp)
	}
}

// GetAccountPlaylist returns an endpoint that makes HTTP requests to the
// account service getAccountPlaylist server.
func (c *Client) GetAccountPlaylist() goa.Endpoint {
	var (
		encodeRequest  = EncodeGetAccountPlaylistRequest(c.encoder)
		decodeResponse = DecodeGetAccountPlaylistResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetAccountPlaylistRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetAccountPlaylistDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "getAccountPlaylist", err)
		}
		return decodeResponse(resp)
	}
}

// AddTrackToAccountPlaylist returns an endpoint that makes HTTP requests to
// the account service addTrackToAccountPlaylist server.
func (c *Client) AddTrackToAccountPlaylist() goa.Endpoint {
	var (
		encodeRequest  = EncodeAddTrackToAccountPlaylistRequest(c.encoder)
		decodeResponse = DecodeAddTrackToAccountPlaylistResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAddTrackToAccountPlaylistRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AddTrackToAccountPlaylistDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "addTrackToAccountPlaylist", err)
		}
		return decodeResponse(resp)
	}
}

// RemoveTrackFromAccountPlaylist returns an endpoint that makes HTTP requests
// to the account service removeTrackFromAccountPlaylist server.
func (c *Client) RemoveTrackFromAccountPlaylist() goa.Endpoint {
	var (
		encodeRequest  = EncodeRemoveTrackFromAccountPlaylistRequest(c.encoder)
		decodeResponse = DecodeRemoveTrackFromAccountPlaylistResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRemoveTrackFromAccountPlaylistRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RemoveTrackFromAccountPlaylistDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("account", "removeTrackFromAccountPlaylist", err)
		}
		return decodeResponse(resp)
	}
}
