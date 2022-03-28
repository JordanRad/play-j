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

	Method("getUserPlaylists", func() {
		Result(UserPlaylistsResponse)
		Payload(func() {
			Attribute("accountID", UInt)
			Attribute("auth", String)
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			GET("/{accountID}/playlists")
		})
	})

	Method("createUserPlaylist", func() {
		Result(CreatePlaylistResponse)
		Payload(func() {
			Attribute("accountID", UInt, "Account ID")
			Attribute("auth", String, "Authorization Header")
			Attribute("name", String, "Playlist name")
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			POST("/{accountID}/playlists")
		})
	})

	Method("deleteUserPlaylist", func() {
		Result(DeletePlaylistResponse)
		Payload(func() {
			Attribute("accountID", UInt)
			Attribute("auth", String)
			Attribute("playlistID", UInt)
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			DELETE("/{accountID}/playlists/{playlistID}")
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
	Attribute("accountID", String, "User's role")
	Required("email", "token", "refresh_token", "role")
})

var UserPlaylistsResponse = Type("UserPlaylistsResponse", func() {
	Attribute("total", Int32, "Number of resources")
	Attribute("resources", ArrayOf(UserSinglePlaylistResponse), "Operation completion status")

	Required("total", "resources")
})

var UserSinglePlaylistResponse = Type("UserSinglePlaylistResponse", func() {
	Attribute("id", Int32, "Playlist id")
	Attribute("name", String, "Playlist name")
	Attribute("tracks", ArrayOf(String), "Operation completion status")

	Required("id", "name", "tracks")
})

var TrackResponse = Type("TrackResponse", func() {
	Attribute("id", Int32, "Track id")
	Attribute("name", String, "Track name")

	Required("id", "name")
})

var CreatePlaylistResponse = Type("CreatePlaylistResponse", func() {
	Attribute("message", String, "Operation completion status")
	Required("message")
})

var DeletePlaylistResponse = Type("DeletePlaylistResponse", func() {
	Attribute("message", String, "Operation completion status")
	Required("message")
})
