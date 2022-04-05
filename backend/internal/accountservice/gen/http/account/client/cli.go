// Code generated by goa v3.7.0, DO NOT EDIT.
//
// account HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package client

import (
	"encoding/json"
	"fmt"

	account "github.com/JordanRad/play-j/backend/internal/accountservice/gen/account"
)

// BuildRegisterPayload builds the payload for the account register endpoint
// from CLI flags.
func BuildRegisterPayload(accountRegisterBody string) (*account.RegisterPayload, error) {
	var err error
	var body RegisterRequestBody
	{
		err = json.Unmarshal([]byte(accountRegisterBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"confirmedPassword\": \"Ab facilis odio facere et.\",\n      \"email\": \"Voluptates id recusandae temporibus et dolore.\",\n      \"firstName\": \"Quis ut ut ipsum et molestiae.\",\n      \"lastName\": \"Dolorum et labore cumque quisquam dolorem adipisci.\",\n      \"password\": \"Numquam quos excepturi vero ad est.\"\n   }'")
		}
	}
	v := &account.RegisterPayload{
		FirstName:         body.FirstName,
		LastName:          body.LastName,
		Email:             body.Email,
		Password:          body.Password,
		ConfirmedPassword: body.ConfirmedPassword,
	}

	return v, nil
}

// BuildLoginPayload builds the payload for the account login endpoint from CLI
// flags.
func BuildLoginPayload(accountLoginBody string) (*account.LoginPayload, error) {
	var err error
	var body LoginRequestBody
	{
		err = json.Unmarshal([]byte(accountLoginBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"email\": \"Corrupti voluptas officia nostrum quia voluptatum.\",\n      \"password\": \"Tempora recusandae nobis.\"\n   }'")
		}
	}
	v := &account.LoginPayload{
		Email:    body.Email,
		Password: body.Password,
	}

	return v, nil
}
