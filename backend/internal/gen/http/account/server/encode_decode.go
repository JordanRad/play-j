// Code generated by goa v3.5.5, DO NOT EDIT.
//
// account HTTP server encoders and decoders
//
// Command:
// $ goa gen git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/design
// -o ./internal/

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	account "git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/gen/account"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeRegisterResponse returns an encoder for responses returned by the
// account register endpoint.
func EncodeRegisterResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*account.RegisterResponse)
		enc := encoder(ctx, w)
		body := NewRegisterResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeRegisterRequest returns a decoder for requests sent to the account
// register endpoint.
func DecodeRegisterRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body RegisterRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateRegisterRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewRegisterPayload(&body)

		return payload, nil
	}
}

// EncodeLoginResponse returns an encoder for responses returned by the account
// login endpoint.
func EncodeLoginResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*account.LoginResponse)
		enc := encoder(ctx, w)
		body := NewLoginResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeLoginRequest returns a decoder for requests sent to the account login
// endpoint.
func DecodeLoginRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body LoginRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		payload := NewLoginPayload(&body)

		return payload, nil
	}
}

// EncodeGetAccountPlaylistCollectionResponse returns an encoder for responses
// returned by the account getAccountPlaylistCollection endpoint.
func EncodeGetAccountPlaylistCollectionResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*account.AccountPlaylistCollectionResponse)
		enc := encoder(ctx, w)
		body := NewGetAccountPlaylistCollectionResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAccountPlaylistCollectionRequest returns a decoder for requests
// sent to the account getAccountPlaylistCollection endpoint.
func DecodeGetAccountPlaylistCollectionRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			accountID uint
			auth      *string
			err       error

			params = mux.Vars(r)
		)
		{
			accountIDRaw := params["accountID"]
			v, err2 := strconv.ParseUint(accountIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("accountID", accountIDRaw, "unsigned integer"))
			}
			accountID = uint(v)
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetAccountPlaylistCollectionPayload(accountID, auth)

		return payload, nil
	}
}

// EncodeCreateAccountPlaylistResponse returns an encoder for responses
// returned by the account createAccountPlaylist endpoint.
func EncodeCreateAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*account.CreatePlaylistResponse)
		enc := encoder(ctx, w)
		body := NewCreateAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateAccountPlaylistRequest returns a decoder for requests sent to
// the account createAccountPlaylist endpoint.
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

		var (
			accountID uint
			auth      *string

			params = mux.Vars(r)
		)
		{
			accountIDRaw := params["accountID"]
			v, err2 := strconv.ParseUint(accountIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("accountID", accountIDRaw, "unsigned integer"))
			}
			accountID = uint(v)
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateAccountPlaylistPayload(&body, accountID, auth)

		return payload, nil
	}
}

// EncodeDeleteAccountPlaylistResponse returns an encoder for responses
// returned by the account deleteAccountPlaylist endpoint.
func EncodeDeleteAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*account.DeletePlaylistResponse)
		enc := encoder(ctx, w)
		body := NewDeleteAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDeleteAccountPlaylistRequest returns a decoder for requests sent to
// the account deleteAccountPlaylist endpoint.
func DecodeDeleteAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			accountID  uint
			playlistID uint
			auth       *string
			err        error

			params = mux.Vars(r)
		)
		{
			accountIDRaw := params["accountID"]
			v, err2 := strconv.ParseUint(accountIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("accountID", accountIDRaw, "unsigned integer"))
			}
			accountID = uint(v)
		}
		{
			playlistIDRaw := params["playlistID"]
			v, err2 := strconv.ParseUint(playlistIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("playlistID", playlistIDRaw, "unsigned integer"))
			}
			playlistID = uint(v)
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeleteAccountPlaylistPayload(accountID, playlistID, auth)

		return payload, nil
	}
}

// EncodeGetAccountPlaylistResponse returns an encoder for responses returned
// by the account getAccountPlaylist endpoint.
func EncodeGetAccountPlaylistResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*account.AccountPlaylistResponse)
		enc := encoder(ctx, w)
		body := NewGetAccountPlaylistResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAccountPlaylistRequest returns a decoder for requests sent to the
// account getAccountPlaylist endpoint.
func DecodeGetAccountPlaylistRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			accountID  uint
			playlistID uint
			auth       *string
			err        error

			params = mux.Vars(r)
		)
		{
			accountIDRaw := params["accountID"]
			v, err2 := strconv.ParseUint(accountIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("accountID", accountIDRaw, "unsigned integer"))
			}
			accountID = uint(v)
		}
		{
			playlistIDRaw := params["playlistID"]
			v, err2 := strconv.ParseUint(playlistIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("playlistID", playlistIDRaw, "unsigned integer"))
			}
			playlistID = uint(v)
		}
		authRaw := r.Header.Get("Authorization")
		if authRaw != "" {
			auth = &authRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetAccountPlaylistPayload(accountID, playlistID, auth)

		return payload, nil
	}
}

// marshalAccountAccountPlaylistResponseToAccountPlaylistResponseResponseBody
// builds a value of type *AccountPlaylistResponseResponseBody from a value of
// type *account.AccountPlaylistResponse.
func marshalAccountAccountPlaylistResponseToAccountPlaylistResponseResponseBody(v *account.AccountPlaylistResponse) *AccountPlaylistResponseResponseBody {
	res := &AccountPlaylistResponseResponseBody{
		ID:   v.ID,
		Name: v.Name,
	}
	if v.TrackIDs != nil {
		res.TrackIDs = make([]int32, len(v.TrackIDs))
		for i, val := range v.TrackIDs {
			res.TrackIDs[i] = val
		}
	}

	return res
}
