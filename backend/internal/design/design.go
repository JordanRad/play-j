package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("accounts", func() {
	Title("Accounts Backend Service")
	Description("Microservice for account management & operations")
})

var _ = Service("account", func() {
	Description("Account operations")

	HTTP(func() {
		Path("/api/v1/account") // Prefix to HTTP path of all requests.
	})

	Method("register", func() {
		Payload(func() {
			Attribute("email", String, "Email of the user")
			Attribute("password", String, "Password of the user")
			Attribute("confirmedPassword", String, "Confirmed password of the user")
			Required("email", "password", "confirmedPassword")
		})
		Result(RegisterResponse)
		HTTP(func() {
			POST("/register")
		})
	})

	Method("login", func() {
		Payload(func() {
			Attribute("email", String, "Email of the user")
			Attribute("password", String, "Password of the user")
		})
		Result(LoginResponse)
		HTTP(func() {
			POST("/login")
		})
	})
})

var RegisterResponse = Type("RegisterResponse", func() {
	Attribute("message", String, "Operation completion status")
	Required("message")
})

var LoginResponse = Type("LoginResponse", func() {
	Attribute("email", String, "User's email")
	Attribute("token", String, "JWT Token")
	Attribute("refresh_token", String, "Refresh token")
	Attribute("role", String, "User's role")
	Required("email", "token", "refresh_token", "role")
})

var UserPlaylistsResponse = Type("UserPlaylistsResponse", func() {
	Attribute("total", Int32, "Number of resources")
	Attribute("resources", ArrayOf(Int32), "Operation completion status")

	Required("total", "resources")
})
