package middleware

import (
	"fmt"
	"net/http"
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
}

func AuthenticateRequest() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			isProtectedRoute := false

			// Check if the route is protected
			for _, route := range protectedRoutes {
				if route.HTTPMethod == r.Method && strings.Contains(r.URL.Path, route.URLContains) {
					isProtectedRoute = true
				}
			}

			// Check if the JWT is valid
			if isProtectedRoute {
				fmt.Print("Protected Route has been requested \n")
				authorizationHeader := r.Header.Get("Authorization")

				tokenString := authorizationHeader[7:]

				//Validate the token
				_, err := auth.ValidateJWT(tokenString)

				// Return 401 by any token validating error
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("401 - Unauthorized. \n Your token has either expired or is invalid"))
				} else {
					h.ServeHTTP(w, r)
				}
			} else {
				fmt.Print("Unprotected route has been requested \n")
				h.ServeHTTP(w, r)
			}
		})
	}
}
