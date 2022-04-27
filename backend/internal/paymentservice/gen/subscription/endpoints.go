// Code generated by goa v3.7.0, DO NOT EDIT.
//
// subscription endpoints
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/payment-service -o
// ./internal/paymentservice

package subscription

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "subscription" service endpoints.
type Endpoints struct {
	GetAccountSubscriptionStatus goa.Endpoint
}

// NewEndpoints wraps the methods of the "subscription" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetAccountSubscriptionStatus: NewGetAccountSubscriptionStatusEndpoint(s),
	}
}

// Use applies the given middleware to all the "subscription" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetAccountSubscriptionStatus = m(e.GetAccountSubscriptionStatus)
}

// NewGetAccountSubscriptionStatusEndpoint returns an endpoint function that
// calls the method "getAccountSubscriptionStatus" of service "subscription".
func NewGetAccountSubscriptionStatusEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetAccountSubscriptionStatusPayload)
		return s.GetAccountSubscriptionStatus(ctx, p)
	}
}