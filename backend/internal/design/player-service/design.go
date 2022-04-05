package design

import (
	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
)

var _ = API("player", func() {
	Title("Play-J Player Service")
	Description("Microservice for streaming music files over the network")
	cors.Origin("/.*localhost.*/", func() {
		cors.Headers("Content-Type", "api_key", "Authorization")
		cors.Methods("GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS")
		cors.MaxAge(600)
	})
})

var _ = Service("player", func() {
	Description("Player operations")

	HTTP(func() {
		Path("/api/v1/player-service/tracks") // Prefix to HTTP path of all requests.
	})

	Method("getTrackByID", func() {
		Payload(func() {
			Attribute("trackID", UInt, "ID of the Track")

			Required("trackID")
		})
		Result(StreamTrackResponse)
		HTTP(func() {
			POST("/{trackID}")
		})
	})

})

var StreamTrackResponse = Type("StreamTrackResponse", func() {
	Attribute("track", ArrayOf(Bytes), "Operation completion status")
	Required("track")
})
