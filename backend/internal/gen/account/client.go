// Code generated by goa v3.5.5, DO NOT EDIT.
//
// account client
//
// Command:
// $ goa gen git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/design
// -o ./internal/

package account

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "account" service client.
type Client struct {
	RegisterEndpoint goa.Endpoint
	LoginEndpoint    goa.Endpoint
}

// NewClient initializes a "account" service client given the endpoints.
func NewClient(register, login goa.Endpoint) *Client {
	return &Client{
		RegisterEndpoint: register,
		LoginEndpoint:    login,
	}
}

// Register calls the "register" endpoint of the "account" service.
func (c *Client) Register(ctx context.Context, p *RegisterPayload) (res *RegisterResponse, err error) {
	var ires interface{}
	ires, err = c.RegisterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*RegisterResponse), nil
}

// Login calls the "login" endpoint of the "account" service.
func (c *Client) Login(ctx context.Context, p *LoginPayload) (res *LoginResponse, err error) {
	var ires interface{}
	ires, err = c.LoginEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LoginResponse), nil
}