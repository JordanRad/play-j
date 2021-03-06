// Code generated by goa v3.7.0, DO NOT EDIT.
//
// account endpoints
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package account

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "account" service endpoints.
type Endpoints struct {
	Register   goa.Endpoint
	Login      goa.Endpoint
	GetProfile goa.Endpoint
}

// NewEndpoints wraps the methods of the "account" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Register:   NewRegisterEndpoint(s),
		Login:      NewLoginEndpoint(s),
		GetProfile: NewGetProfileEndpoint(s),
	}
}

// Use applies the given middleware to all the "account" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Register = m(e.Register)
	e.Login = m(e.Login)
	e.GetProfile = m(e.GetProfile)
}

// NewRegisterEndpoint returns an endpoint function that calls the method
// "register" of service "account".
func NewRegisterEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RegisterPayload)
		return s.Register(ctx, p)
	}
}

// NewLoginEndpoint returns an endpoint function that calls the method "login"
// of service "account".
func NewLoginEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*LoginPayload)
		return s.Login(ctx, p)
	}
}

// NewGetProfileEndpoint returns an endpoint function that calls the method
// "getProfile" of service "account".
func NewGetProfileEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetProfilePayload)
		return s.GetProfile(ctx, p)
	}
}
