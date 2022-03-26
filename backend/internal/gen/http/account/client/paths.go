// Code generated by goa v3.5.5, DO NOT EDIT.
//
// HTTP request path constructors for the account service.
//
// Command:
// $ goa gen git.fhict.nl/I425652/jordan-portfolio-s6/backend/internal/design
// -o ./internal/

package client

// RegisterAccountPath returns the URL path to the account service register HTTP endpoint.
func RegisterAccountPath() string {
	return "/api/v1/account/register"
}

// LoginAccountPath returns the URL path to the account service login HTTP endpoint.
func LoginAccountPath() string {
	return "/api/v1/account/login"
}
