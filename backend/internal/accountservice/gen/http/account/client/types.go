// Code generated by goa v3.7.0, DO NOT EDIT.
//
// account HTTP client types
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package client

import (
	account "github.com/JordanRad/play-j/backend/internal/accountservice/gen/account"
	goa "goa.design/goa/v3/pkg"
)

// RegisterRequestBody is the type of the "account" service "register" endpoint
// HTTP request body.
type RegisterRequestBody struct {
	// First name of the user
	FirstName string `form:"firstName" json:"firstName" xml:"firstName"`
	// Last name of the user
	LastName string `form:"lastName" json:"lastName" xml:"lastName"`
	// Email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Password of the user
	Password string `form:"password" json:"password" xml:"password"`
	// Confirmed password of the user
	ConfirmedPassword string `form:"confirmedPassword" json:"confirmedPassword" xml:"confirmedPassword"`
}

// LoginRequestBody is the type of the "account" service "login" endpoint HTTP
// request body.
type LoginRequestBody struct {
	// Email of the user
	Email string `form:"email" json:"email" xml:"email"`
	// Password of the user
	Password string `form:"password" json:"password" xml:"password"`
}

// RegisterResponseBody is the type of the "account" service "register"
// endpoint HTTP response body.
type RegisterResponseBody struct {
	// Operation completion status
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// LoginResponseBody is the type of the "account" service "login" endpoint HTTP
// response body.
type LoginResponseBody struct {
	// User's email
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// JWT Token
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	// Refresh token
	RefreshToken *string `form:"refresh_token,omitempty" json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// User's role
	Role *string `form:"role,omitempty" json:"role,omitempty" xml:"role,omitempty"`
	// User's role
	AccountID *string `form:"accountID,omitempty" json:"accountID,omitempty" xml:"accountID,omitempty"`
}

// GetProfileResponseBody is the type of the "account" service "getProfile"
// endpoint HTTP response body.
type GetProfileResponseBody struct {
	// Operation completion status
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// First Name
	FirstName *string `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	// Last Name
	LastName *string `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
	// Username
	Username *string `form:"username,omitempty" json:"username,omitempty" xml:"username,omitempty"`
	// Array of last payments
	LastPayments []*PaymentResponseResponseBody `form:"last_payments,omitempty" json:"last_payments,omitempty" xml:"last_payments,omitempty"`
}

// PaymentResponseResponseBody is used to define fields on response body types.
type PaymentResponseResponseBody struct {
	ID            *uint    `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	CreatedAt     *string  `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	PaymentNumber *string  `form:"paymentNumber,omitempty" json:"paymentNumber,omitempty" xml:"paymentNumber,omitempty"`
	Amount        *float32 `form:"amount,omitempty" json:"amount,omitempty" xml:"amount,omitempty"`
}

// NewRegisterRequestBody builds the HTTP request body from the payload of the
// "register" endpoint of the "account" service.
func NewRegisterRequestBody(p *account.RegisterPayload) *RegisterRequestBody {
	body := &RegisterRequestBody{
		FirstName:         p.FirstName,
		LastName:          p.LastName,
		Email:             p.Email,
		Password:          p.Password,
		ConfirmedPassword: p.ConfirmedPassword,
	}
	return body
}

// NewLoginRequestBody builds the HTTP request body from the payload of the
// "login" endpoint of the "account" service.
func NewLoginRequestBody(p *account.LoginPayload) *LoginRequestBody {
	body := &LoginRequestBody{
		Email:    p.Email,
		Password: p.Password,
	}
	return body
}

// NewRegisterResponseOK builds a "account" service "register" endpoint result
// from a HTTP "OK" response.
func NewRegisterResponseOK(body *RegisterResponseBody) *account.RegisterResponse {
	v := &account.RegisterResponse{
		Message: *body.Message,
	}

	return v
}

// NewLoginResponseOK builds a "account" service "login" endpoint result from a
// HTTP "OK" response.
func NewLoginResponseOK(body *LoginResponseBody) *account.LoginResponse {
	v := &account.LoginResponse{
		Email:        *body.Email,
		Token:        *body.Token,
		RefreshToken: *body.RefreshToken,
		Role:         *body.Role,
		AccountID:    body.AccountID,
	}

	return v
}

// NewGetProfileProfileResponseOK builds a "account" service "getProfile"
// endpoint result from a HTTP "OK" response.
func NewGetProfileProfileResponseOK(body *GetProfileResponseBody) *account.ProfileResponse {
	v := &account.ProfileResponse{
		Email:     *body.Email,
		FirstName: *body.FirstName,
		LastName:  *body.LastName,
		Username:  *body.Username,
	}
	v.LastPayments = make([]*account.PaymentResponse, len(body.LastPayments))
	for i, val := range body.LastPayments {
		v.LastPayments[i] = unmarshalPaymentResponseResponseBodyToAccountPaymentResponse(val)
	}

	return v
}

// ValidateRegisterResponseBody runs the validations defined on
// RegisterResponseBody
func ValidateRegisterResponseBody(body *RegisterResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateLoginResponseBody runs the validations defined on LoginResponseBody
func ValidateLoginResponseBody(body *LoginResponseBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "body"))
	}
	if body.RefreshToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("refresh_token", "body"))
	}
	if body.Role == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("role", "body"))
	}
	return
}

// ValidateGetProfileResponseBody runs the validations defined on
// GetProfileResponseBody
func ValidateGetProfileResponseBody(body *GetProfileResponseBody) (err error) {
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first_name", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last_name", "body"))
	}
	if body.Username == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("username", "body"))
	}
	if body.LastPayments == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last_payments", "body"))
	}
	for _, e := range body.LastPayments {
		if e != nil {
			if err2 := ValidatePaymentResponseResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidatePaymentResponseResponseBody runs the validations defined on
// PaymentResponseResponseBody
func ValidatePaymentResponseResponseBody(body *PaymentResponseResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.CreatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("createdAt", "body"))
	}
	if body.PaymentNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("paymentNumber", "body"))
	}
	if body.Amount == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("amount", "body"))
	}
	return
}
