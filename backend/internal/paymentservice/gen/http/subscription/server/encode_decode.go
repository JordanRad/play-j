// Code generated by goa v3.7.0, DO NOT EDIT.
//
// subscription HTTP server encoders and decoders
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/payment-service -o
// ./internal/paymentservice

package server

import (
	"context"
	"net/http"

	subscription "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/subscription"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetAccountSubscriptionStatusResponse returns an encoder for responses
// returned by the subscription getAccountSubscriptionStatus endpoint.
func EncodeGetAccountSubscriptionStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*subscription.SubscriptionStatusResponse)
		enc := encoder(ctx, w)
		body := NewGetAccountSubscriptionStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetAccountSubscriptionStatusRequest returns a decoder for requests
// sent to the subscription getAccountSubscriptionStatus endpoint.
func DecodeGetAccountSubscriptionStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
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
		payload := NewGetAccountSubscriptionStatusPayload(auth)

		return payload, nil
	}
}
