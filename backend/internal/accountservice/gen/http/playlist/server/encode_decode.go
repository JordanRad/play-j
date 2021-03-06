// Code generated by goa v3.7.0, DO NOT EDIT.
//
// playlist HTTP server encoders and decoders
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	playlist "github.com/JordanRad/play-j/backend/internal/accountservice/gen/playlist"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetAccountPlaylistCollectionResponse returns an encoder for responses
// returned by the playlist getAccountPlaylistCollection endpoint.
func EncodeGetAccountPlaylistCollectionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*playlist.AccountPlaylistCollectionResponse)
		enc := encoder(ctx, w)
		body := NewGetAccountPlaylistCollectionResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAccountPlaylistCollectionRequest returns a decoder for requests
// sent to the playlist getAccountPlaylistCollection endpoint.
func DecodeGetAccountPlaylistCollectionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			auth string
			err  error
		)
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetAccountPlaylistCollectionPayload(auth)

		return payload, nil
	}
}

// EncodeCreateAccountPlaylistResponse returns an encoder for responses
// returned by the playlist createAccountPlaylist endpoint.
func EncodeCreateAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*playlist.PlaylistModificationResponse)
		enc := encoder(ctx, w)
		body := NewCreateAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateAccountPlaylistRequest returns a decoder for requests sent to
// the playlist createAccountPlaylist endpoint.
func DecodeCreateAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateAccountPlaylistRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateAccountPlaylistRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			auth string
		)
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateAccountPlaylistPayload(&body, auth)

		return payload, nil
	}
}

// EncodeRenameAccountPlaylistResponse returns an encoder for responses
// returned by the playlist renameAccountPlaylist endpoint.
func EncodeRenameAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*playlist.PlaylistModificationResponse)
		enc := encoder(ctx, w)
		body := NewRenameAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRenameAccountPlaylistRequest returns a decoder for requests sent to
// the playlist renameAccountPlaylist endpoint.
func DecodeRenameAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RenameAccountPlaylistRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRenameAccountPlaylistRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			playlistID uint
			auth       string

			params = mux.Vars(r)
		)
		{
			playlistIDRaw := params["playlistID"]
			v, err2 := strconv.ParseUint(playlistIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("playlistID", playlistIDRaw, "unsigned integer"))
			}
			playlistID = uint(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRenameAccountPlaylistPayload(&body, playlistID, auth)

		return payload, nil
	}
}

// EncodeDeleteAccountPlaylistResponse returns an encoder for responses
// returned by the playlist deleteAccountPlaylist endpoint.
func EncodeDeleteAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*playlist.PlaylistModificationResponse)
		enc := encoder(ctx, w)
		body := NewDeleteAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteAccountPlaylistRequest returns a decoder for requests sent to
// the playlist deleteAccountPlaylist endpoint.
func DecodeDeleteAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			playlistID uint
			auth       string
			err        error

			params = mux.Vars(r)
		)
		{
			playlistIDRaw := params["playlistID"]
			v, err2 := strconv.ParseUint(playlistIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("playlistID", playlistIDRaw, "unsigned integer"))
			}
			playlistID = uint(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteAccountPlaylistPayload(playlistID, auth)

		return payload, nil
	}
}

// EncodeGetAccountPlaylistResponse returns an encoder for responses returned
// by the playlist getAccountPlaylist endpoint.
func EncodeGetAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*playlist.AccountPlaylistResponse)
		enc := encoder(ctx, w)
		body := NewGetAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAccountPlaylistRequest returns a decoder for requests sent to the
// playlist getAccountPlaylist endpoint.
func DecodeGetAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			playlistID uint
			auth       string
			err        error

			params = mux.Vars(r)
		)
		{
			playlistIDRaw := params["playlistID"]
			v, err2 := strconv.ParseUint(playlistIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("playlistID", playlistIDRaw, "unsigned integer"))
			}
			playlistID = uint(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetAccountPlaylistPayload(playlistID, auth)

		return payload, nil
	}
}

// EncodeAddTrackToAccountPlaylistResponse returns an encoder for responses
// returned by the playlist addTrackToAccountPlaylist endpoint.
func EncodeAddTrackToAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*playlist.PlaylistModificationResponse)
		enc := encoder(ctx, w)
		body := NewAddTrackToAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAddTrackToAccountPlaylistRequest returns a decoder for requests sent
// to the playlist addTrackToAccountPlaylist endpoint.
func DecodeAddTrackToAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			playlistID uint
			trackID    uint
			auth       string
			err        error

			params = mux.Vars(r)
		)
		{
			playlistIDRaw := params["playlistID"]
			v, err2 := strconv.ParseUint(playlistIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("playlistID", playlistIDRaw, "unsigned integer"))
			}
			playlistID = uint(v)
		}
		{
			trackIDRaw := params["trackID"]
			v, err2 := strconv.ParseUint(trackIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("trackID", trackIDRaw, "unsigned integer"))
			}
			trackID = uint(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewAddTrackToAccountPlaylistPayload(playlistID, trackID, auth)

		return payload, nil
	}
}

// EncodeRemoveTrackFromAccountPlaylistResponse returns an encoder for
// responses returned by the playlist removeTrackFromAccountPlaylist endpoint.
func EncodeRemoveTrackFromAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*playlist.PlaylistModificationResponse)
		enc := encoder(ctx, w)
		body := NewRemoveTrackFromAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRemoveTrackFromAccountPlaylistRequest returns a decoder for requests
// sent to the playlist removeTrackFromAccountPlaylist endpoint.
func DecodeRemoveTrackFromAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			playlistID uint
			trackID    uint
			auth       string
			err        error

			params = mux.Vars(r)
		)
		{
			playlistIDRaw := params["playlistID"]
			v, err2 := strconv.ParseUint(playlistIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("playlistID", playlistIDRaw, "unsigned integer"))
			}
			playlistID = uint(v)
		}
		{
			trackIDRaw := params["trackID"]
			v, err2 := strconv.ParseUint(trackIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("trackID", trackIDRaw, "unsigned integer"))
			}
			trackID = uint(v)
		}
		auth = r.Header.Get("Authorization")
		if auth == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Authorization", "header"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewRemoveTrackFromAccountPlaylistPayload(playlistID, trackID, auth)

		return payload, nil
	}
}

// marshalPlaylistAccountPlaylistResponseToAccountPlaylistResponseResponseBody
// builds a value of type *AccountPlaylistResponseResponseBody from a value of
// type *playlist.AccountPlaylistResponse.
func marshalPlaylistAccountPlaylistResponseToAccountPlaylistResponseResponseBody(v *playlist.AccountPlaylistResponse) *AccountPlaylistResponseResponseBody {
	res := &AccountPlaylistResponseResponseBody{
		ID:        v.ID,
		Name:      v.Name,
		CreatedAt: v.CreatedAt,
	}
	if v.TrackIDs != nil {
		res.TrackIDs = make([]int32, len(v.TrackIDs))
		for i, val := range v.TrackIDs {
			res.TrackIDs[i] = val
		}
	}

	return res
}
