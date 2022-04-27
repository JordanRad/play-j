// Code generated by goa v3.7.0, DO NOT EDIT.
//
// subscription HTTP client types
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/payment-service -o
// ./internal/paymentservice

package client

import (
	subscription "github.com/JordanRad/play-j/backend/internal/paymentservice/gen/subscription"
	goa "goa.design/goa/v3/pkg"
)

// GetAccountSubscriptionStatusResponseBody is the type of the "subscription"
// service "getAccountSubscriptionStatus" endpoint HTTP response body.
type GetAccountSubscriptionStatusResponseBody struct {
	// Is the subscrion paid or not
	HasPaid *bool `form:"hasPaid,omitempty" json:"hasPaid,omitempty" xml:"hasPaid,omitempty"`
	// Subscription end date
	ValidUntil *string `form:"validUntil,omitempty" json:"validUntil,omitempty" xml:"validUntil,omitempty"`
	// Subscription type
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// NewGetAccountSubscriptionStatusSubscriptionStatusResponseOK builds a
// "subscription" service "getAccountSubscriptionStatus" endpoint result from a
// HTTP "OK" response.
func NewGetAccountSubscriptionStatusSubscriptionStatusResponseOK(body *GetAccountSubscriptionStatusResponseBody) *subscription.SubscriptionStatusResponse {
	v := &subscription.SubscriptionStatusResponse{
		HasPaid:    *body.HasPaid,
		ValidUntil: *body.ValidUntil,
		Type:       *body.Type,
	}

	return v
}

// ValidateGetAccountSubscriptionStatusResponseBody runs the validations
// defined on GetAccountSubscriptionStatusResponseBody
func ValidateGetAccountSubscriptionStatusResponseBody(body *GetAccountSubscriptionStatusResponseBody) (err error) {
	if body.HasPaid == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("hasPaid", "body"))
	}
	if body.ValidUntil == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("validUntil", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}