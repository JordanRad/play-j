package middleware

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/JordanRad/play-j/backend/internal/auth"
)

type ProtectedRoute struct {
	HTTPMethod  string
	URLContains string
}

var protectedRoutes = []ProtectedRoute{
	{
		HTTPMethod:  "GET",
		URLContains: "/playlists",
	},
	{
		HTTPMethod:  "POST",
		URLContains: "/playlists",
	},
	{
		HTTPMethod:  "DELETE",
		URLContains: "/playlists",
	},
	{
		HTTPMethod:  "GET",
		URLContains: "/payments",
	},
	{
		HTTPMethod:  "POST",
		URLContains: "/payments",
	},
	{
		HTTPMethod:  "GET",
		URLContains: "/subscriptions",
	},
	{
		HTTPMethod:  "GET",
		URLContains: "/profile",
	},
}

func isRouteProtected(method string, URL *url.URL) bool {
	for _, route := range protectedRoutes {
		if route.HTTPMethod == method && strings.Contains(URL.Path, route.URLContains) {
			return true
		}
	}
	return false
}

func AuthenticateRequest() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the route is protected
			isProtectedRoute := isRouteProtected(r.Method, r.URL)

			// Check if the JWT is valid
			if isProtectedRoute {
				authorizationHeader := r.Header.Get("Authorization")
				tokenString := authorizationHeader[7:]

				//Validate the token
				_, err := auth.ValidateJWT(tokenString)

				// Return 401 by any token validating error
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 - Unauthorized. \n Your token has either expired or is invalid"))
				} else {
					// Add the JWT to the context
					ctx := context.WithValue(r.Context(), "jwt", tokenString)
					requestWithContext := r.WithContext(ctx)

					h.ServeHTTP(w, requestWithContext)
				}
			} else {
				h.ServeHTTP(w, r)
			}
		})
	}
}
