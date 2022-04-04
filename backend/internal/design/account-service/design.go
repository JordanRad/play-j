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
		Path("/api/v1/account-service/accounts") // Prefix to HTTP path of all requests.
	})

	Method("register", func() {
		Payload(func() {
			Attribute("firstName", String, "First name of the user")
			Attribute("lastName", String, "Last name of the user")
			Attribute("email", String, "Email of the user")
			Attribute("password", String, "Password of the user")
			Attribute("confirmedPassword", String, "Confirmed password of the user")
			Required("email", "password", "confirmedPassword", "firstName", "lastName")
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

var _ = Service("playlist", func() {
	Description("Playlist operations")

	HTTP(func() {
		Path("/api/v1/account-service/playlists") // Prefix to HTTP path of all requests.
	})

	Method("getAccountPlaylistCollection", func() {
		Result(AccountPlaylistCollectionResponse)
		Payload(func() {
			Attribute("accountID", UInt)
			Attribute("auth", String)
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			GET("/")
		})
	})

	Method("createAccountPlaylist", func() {
		Result(PlaylistModificationResponse)
		Payload(func() {
			Attribute("auth", String, "Authorization Header")
			Attribute("name", String, "Playlist name")
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			POST("/")
		})
	})

	Method("renameAccountPlaylist", func() {
		Result(PlaylistModificationResponse)
		Payload(func() {
			Attribute("auth", String, "Authorization Header")
			Attribute("playlistID", UInt, "Playlist id to modify")
			Attribute("name", String, "New playlist name")
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			PUT("/{playlistID}")
		})
	})

	Method("deleteAccountPlaylist", func() {
		Result(PlaylistModificationResponse)
		Payload(func() {
			Attribute("auth", String)
			Attribute("playlistID", UInt)
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			DELETE("/{playlistID}")
		})
	})

	Method("getAccountPlaylist", func() {
		Result(AccountPlaylistResponse)
		Payload(func() {
			Attribute("playlistID", UInt, "Playlist ID")
			Attribute("auth", String, "Authorization Header")
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			GET("/{playlistID}")
		})
	})

	Method("addTrackToAccountPlaylist", func() {
		Result(PlaylistModificationResponse)
		Payload(func() {
			Attribute("playlistID", UInt, "Playlist ID to modify")
			Attribute("trackID", UInt, "Track ID to be added")
			Attribute("auth", String, "Authorization Header")
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			POST("/{playlistID}/tracks/{trackID}")
		})
	})

	Method("removeTrackFromAccountPlaylist", func() {
		Result(PlaylistModificationResponse)
		Payload(func() {
			Attribute("playlistID", UInt, "Playlist ID to modify")
			Attribute("trackID", UInt, "Track ID to be deleted")
			Attribute("auth", String, "Authorization Header")
		})
		HTTP(func() {
			Header("auth:Authorization", String, "JSON Web Token", func() {
				Pattern("^Bearer [^ ]+$")
			})
			DELETE("/{playlistID}/tracks/{trackID}")
		})
	})
})

var _ = Service("swagger", func() {
	Files("/openapi.json", "./gen/http/openapi.json", func() {
		Meta("swagger:generate", "false")
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

var AccountPlaylistCollectionResponse = Type("AccountPlaylistCollectionResponse", func() {
	Attribute("total", Int32, "Number of resources")
	Attribute("resources", ArrayOf(AccountPlaylistResponse), "Operation completion status")

	Required("total", "resources")
})

var AccountPlaylistResponse = Type("AccountPlaylistResponse", func() {
	Attribute("id", Int32, "Playlist id")
	Attribute("name", String, "Playlist name")
	Attribute("trackIDs", ArrayOf(Int32), "Array of TrackIDs")

	Required("id", "name", "trackIDs")
})

var PlaylistModificationResponse = Type("PlaylistModificationResponse", func() {
	Attribute("message", String, "Operation completion status")
	Required("message")
})
