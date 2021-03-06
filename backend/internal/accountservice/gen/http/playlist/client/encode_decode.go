// Code generated by goa v3.7.0, DO NOT EDIT.
//
// playlist HTTP client encoders and decoders
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	playlist "github.com/JordanRad/play-j/backend/internal/accountservice/gen/playlist"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetAccountPlaylistCollectionRequest instantiates a HTTP request object
// with method and path set to call the "playlist" service
// "getAccountPlaylistCollection" endpoint
func (c *Client) BuildGetAccountPlaylistCollectionRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAccountPlaylistCollectionPlaylistPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("playlist", "getAccountPlaylistCollection", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetAccountPlaylistCollectionRequest returns an encoder for requests
// sent to the playlist getAccountPlaylistCollection server.
func EncodeGetAccountPlaylistCollectionRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*playlist.GetAccountPlaylistCollectionPayload)
		if !ok {
			return goahttp.ErrInvalidType("playlist", "getAccountPlaylistCollection", "*playlist.GetAccountPlaylistCollectionPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetAccountPlaylistCollectionResponse returns a decoder for responses
// returned by the playlist getAccountPlaylistCollection endpoint. restoreBody
// controls whether the response body should be restored after having been read.
func DecodeGetAccountPlaylistCollectionResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetAccountPlaylistCollectionResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("playlist", "getAccountPlaylistCollection", err)
			}
			err = ValidateGetAccountPlaylistCollectionResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("playlist", "getAccountPlaylistCollection", err)
			}
			res := NewGetAccountPlaylistCollectionAccountPlaylistCollectionResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("playlist", "getAccountPlaylistCollection", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateAccountPlaylistRequest instantiates a HTTP request object with
// method and path set to call the "playlist" service "createAccountPlaylist"
// endpoint
func (c *Client) BuildCreateAccountPlaylistRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAccountPlaylistPlaylistPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("playlist", "createAccountPlaylist", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateAccountPlaylistRequest returns an encoder for requests sent to
// the playlist createAccountPlaylist server.
func EncodeCreateAccountPlaylistRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*playlist.CreateAccountPlaylistPayload)
		if !ok {
			return goahttp.ErrInvalidType("playlist", "createAccountPlaylist", "*playlist.CreateAccountPlaylistPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		body := NewCreateAccountPlaylistRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("playlist", "createAccountPlaylist", err)
		}
		return nil
	}
}

// DecodeCreateAccountPlaylistResponse returns a decoder for responses returned
// by the playlist createAccountPlaylist endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeCreateAccountPlaylistResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateAccountPlaylistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("playlist", "createAccountPlaylist", err)
			}
			err = ValidateCreateAccountPlaylistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("playlist", "createAccountPlaylist", err)
			}
			res := NewCreateAccountPlaylistPlaylistModificationResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("playlist", "createAccountPlaylist", resp.StatusCode, string(body))
		}
	}
}

// BuildRenameAccountPlaylistRequest instantiates a HTTP request object with
// method and path set to call the "playlist" service "renameAccountPlaylist"
// endpoint
func (c *Client) BuildRenameAccountPlaylistRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		playlistID uint
	)
	{
		p, ok := v.(*playlist.RenameAccountPlaylistPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("playlist", "renameAccountPlaylist", "*playlist.RenameAccountPlaylistPayload", v)
		}
		playlistID = p.PlaylistID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RenameAccountPlaylistPlaylistPath(playlistID)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("playlist", "renameAccountPlaylist", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRenameAccountPlaylistRequest returns an encoder for requests sent to
// the playlist renameAccountPlaylist server.
func EncodeRenameAccountPlaylistRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*playlist.RenameAccountPlaylistPayload)
		if !ok {
			return goahttp.ErrInvalidType("playlist", "renameAccountPlaylist", "*playlist.RenameAccountPlaylistPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		body := NewRenameAccountPlaylistRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("playlist", "renameAccountPlaylist", err)
		}
		return nil
	}
}

// DecodeRenameAccountPlaylistResponse returns a decoder for responses returned
// by the playlist renameAccountPlaylist endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeRenameAccountPlaylistResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RenameAccountPlaylistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("playlist", "renameAccountPlaylist", err)
			}
			err = ValidateRenameAccountPlaylistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("playlist", "renameAccountPlaylist", err)
			}
			res := NewRenameAccountPlaylistPlaylistModificationResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("playlist", "renameAccountPlaylist", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteAccountPlaylistRequest instantiates a HTTP request object with
// method and path set to call the "playlist" service "deleteAccountPlaylist"
// endpoint
func (c *Client) BuildDeleteAccountPlaylistRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		playlistID uint
	)
	{
		p, ok := v.(*playlist.DeleteAccountPlaylistPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("playlist", "deleteAccountPlaylist", "*playlist.DeleteAccountPlaylistPayload", v)
		}
		playlistID = p.PlaylistID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteAccountPlaylistPlaylistPath(playlistID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("playlist", "deleteAccountPlaylist", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDeleteAccountPlaylistRequest returns an encoder for requests sent to
// the playlist deleteAccountPlaylist server.
func EncodeDeleteAccountPlaylistRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*playlist.DeleteAccountPlaylistPayload)
		if !ok {
			return goahttp.ErrInvalidType("playlist", "deleteAccountPlaylist", "*playlist.DeleteAccountPlaylistPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeDeleteAccountPlaylistResponse returns a decoder for responses returned
// by the playlist deleteAccountPlaylist endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeDeleteAccountPlaylistResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body DeleteAccountPlaylistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("playlist", "deleteAccountPlaylist", err)
			}
			err = ValidateDeleteAccountPlaylistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("playlist", "deleteAccountPlaylist", err)
			}
			res := NewDeleteAccountPlaylistPlaylistModificationResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("playlist", "deleteAccountPlaylist", resp.StatusCode, string(body))
		}
	}
}

// BuildGetAccountPlaylistRequest instantiates a HTTP request object with
// method and path set to call the "playlist" service "getAccountPlaylist"
// endpoint
func (c *Client) BuildGetAccountPlaylistRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		playlistID uint
	)
	{
		p, ok := v.(*playlist.GetAccountPlaylistPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("playlist", "getAccountPlaylist", "*playlist.GetAccountPlaylistPayload", v)
		}
		playlistID = p.PlaylistID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAccountPlaylistPlaylistPath(playlistID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("playlist", "getAccountPlaylist", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetAccountPlaylistRequest returns an encoder for requests sent to the
// playlist getAccountPlaylist server.
func EncodeGetAccountPlaylistRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*playlist.GetAccountPlaylistPayload)
		if !ok {
			return goahttp.ErrInvalidType("playlist", "getAccountPlaylist", "*playlist.GetAccountPlaylistPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeGetAccountPlaylistResponse returns a decoder for responses returned by
// the playlist getAccountPlaylist endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetAccountPlaylistResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetAccountPlaylistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("playlist", "getAccountPlaylist", err)
			}
			err = ValidateGetAccountPlaylistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("playlist", "getAccountPlaylist", err)
			}
			res := NewGetAccountPlaylistAccountPlaylistResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("playlist", "getAccountPlaylist", resp.StatusCode, string(body))
		}
	}
}

// BuildAddTrackToAccountPlaylistRequest instantiates a HTTP request object
// with method and path set to call the "playlist" service
// "addTrackToAccountPlaylist" endpoint
func (c *Client) BuildAddTrackToAccountPlaylistRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		playlistID uint
		trackID    uint
	)
	{
		p, ok := v.(*playlist.AddTrackToAccountPlaylistPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("playlist", "addTrackToAccountPlaylist", "*playlist.AddTrackToAccountPlaylistPayload", v)
		}
		playlistID = p.PlaylistID
		trackID = p.TrackID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AddTrackToAccountPlaylistPlaylistPath(playlistID, trackID)}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("playlist", "addTrackToAccountPlaylist", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAddTrackToAccountPlaylistRequest returns an encoder for requests sent
// to the playlist addTrackToAccountPlaylist server.
func EncodeAddTrackToAccountPlaylistRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*playlist.AddTrackToAccountPlaylistPayload)
		if !ok {
			return goahttp.ErrInvalidType("playlist", "addTrackToAccountPlaylist", "*playlist.AddTrackToAccountPlaylistPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeAddTrackToAccountPlaylistResponse returns a decoder for responses
// returned by the playlist addTrackToAccountPlaylist endpoint. restoreBody
// controls whether the response body should be restored after having been read.
func DecodeAddTrackToAccountPlaylistResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body AddTrackToAccountPlaylistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("playlist", "addTrackToAccountPlaylist", err)
			}
			err = ValidateAddTrackToAccountPlaylistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("playlist", "addTrackToAccountPlaylist", err)
			}
			res := NewAddTrackToAccountPlaylistPlaylistModificationResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("playlist", "addTrackToAccountPlaylist", resp.StatusCode, string(body))
		}
	}
}

// BuildRemoveTrackFromAccountPlaylistRequest instantiates a HTTP request
// object with method and path set to call the "playlist" service
// "removeTrackFromAccountPlaylist" endpoint
func (c *Client) BuildRemoveTrackFromAccountPlaylistRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		playlistID uint
		trackID    uint
	)
	{
		p, ok := v.(*playlist.RemoveTrackFromAccountPlaylistPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("playlist", "removeTrackFromAccountPlaylist", "*playlist.RemoveTrackFromAccountPlaylistPayload", v)
		}
		playlistID = p.PlaylistID
		trackID = p.TrackID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: RemoveTrackFromAccountPlaylistPlaylistPath(playlistID, trackID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("playlist", "removeTrackFromAccountPlaylist", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeRemoveTrackFromAccountPlaylistRequest returns an encoder for requests
// sent to the playlist removeTrackFromAccountPlaylist server.
func EncodeRemoveTrackFromAccountPlaylistRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*playlist.RemoveTrackFromAccountPlaylistPayload)
		if !ok {
			return goahttp.ErrInvalidType("playlist", "removeTrackFromAccountPlaylist", "*playlist.RemoveTrackFromAccountPlaylistPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeRemoveTrackFromAccountPlaylistResponse returns a decoder for responses
// returned by the playlist removeTrackFromAccountPlaylist endpoint.
// restoreBody controls whether the response body should be restored after
// having been read.
func DecodeRemoveTrackFromAccountPlaylistResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body RemoveTrackFromAccountPlaylistResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("playlist", "removeTrackFromAccountPlaylist", err)
			}
			err = ValidateRemoveTrackFromAccountPlaylistResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("playlist", "removeTrackFromAccountPlaylist", err)
			}
			res := NewRemoveTrackFromAccountPlaylistPlaylistModificationResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("playlist", "removeTrackFromAccountPlaylist", resp.StatusCode, string(body))
		}
	}
}

// unmarshalAccountPlaylistResponseResponseBodyToPlaylistAccountPlaylistResponse
// builds a value of type *playlist.AccountPlaylistResponse from a value of
// type *AccountPlaylistResponseResponseBody.
func unmarshalAccountPlaylistResponseResponseBodyToPlaylistAccountPlaylistResponse(v *AccountPlaylistResponseResponseBody) *playlist.AccountPlaylistResponse {
	res := &playlist.AccountPlaylistResponse{
		ID:        *v.ID,
		Name:      *v.Name,
		CreatedAt: v.CreatedAt,
	}
	res.TrackIDs = make([]int32, len(v.TrackIDs))
	for i, val := range v.TrackIDs {
		res.TrackIDs[i] = val
	}

	return res
}
