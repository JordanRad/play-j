// Code generated by goa v3.7.0, DO NOT EDIT.
//
// payment HTTP client encoders and decoders
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/payment-service -o
// ./internal/paymentservice

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	payment "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/payment"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetAccountPaymentsRequest instantiates a HTTP request object with
// method and path set to call the "payment" service "getAccountPayments"
// endpoint
func (c *Client) BuildGetAccountPaymentsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetAccountPaymentsPaymentPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("payment", "getAccountPayments", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeGetAccountPaymentsRequest returns an encoder for requests sent to the
// payment getAccountPayments server.
func EncodeGetAccountPaymentsRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*payment.GetAccountPaymentsPayload)
		if !ok {
			return goahttp.ErrInvalidType("payment", "getAccountPayments", "*payment.GetAccountPaymentsPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		values := req.URL.Query()
		values.Add("limit", fmt.Sprintf("%v", p.Limit))
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeGetAccountPaymentsResponse returns a decoder for responses returned by
// the payment getAccountPayments endpoint. restoreBody controls whether the
// response body should be restored after having been read.
func DecodeGetAccountPaymentsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body GetAccountPaymentsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("payment", "getAccountPayments", err)
			}
			err = ValidateGetAccountPaymentsResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("payment", "getAccountPayments", err)
			}
			res := NewGetAccountPaymentsPaymentListResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("payment", "getAccountPayments", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateAccountPaymentRequest instantiates a HTTP request object with
// method and path set to call the "payment" service "createAccountPayment"
// endpoint
func (c *Client) BuildCreateAccountPaymentRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateAccountPaymentPaymentPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("payment", "createAccountPayment", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateAccountPaymentRequest returns an encoder for requests sent to
// the payment createAccountPayment server.
func EncodeCreateAccountPaymentRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*payment.CreateAccountPaymentPayload)
		if !ok {
			return goahttp.ErrInvalidType("payment", "createAccountPayment", "*payment.CreateAccountPaymentPayload", v)
		}
		{
			head := p.Auth
			req.Header.Set("Authorization", head)
		}
		return nil
	}
}

// DecodeCreateAccountPaymentResponse returns a decoder for responses returned
// by the payment createAccountPayment endpoint. restoreBody controls whether
// the response body should be restored after having been read.
func DecodeCreateAccountPaymentResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body CreateAccountPaymentResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("payment", "createAccountPayment", err)
			}
			err = ValidateCreateAccountPaymentResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("payment", "createAccountPayment", err)
			}
			res := NewCreateAccountPaymentTransactionResponseOK(&body)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("payment", "createAccountPayment", resp.StatusCode, string(body))
		}
	}
}

// unmarshalPaymentResponseResponseBodyToPaymentPaymentResponse builds a value
// of type *payment.PaymentResponse from a value of type
// *PaymentResponseResponseBody.
func unmarshalPaymentResponseResponseBodyToPaymentPaymentResponse(v *PaymentResponseResponseBody) *payment.PaymentResponse {
	res := &payment.PaymentResponse{
		ID:            *v.ID,
		CreatedAt:     *v.CreatedAt,
		PaymentNumber: *v.PaymentNumber,
		Amount:        *v.Amount,
	}

	return res
}
