// Code generated by goa v3.7.0, DO NOT EDIT.
//
// player views
//
// Command:
// $ goa gen github.com/JordanRad/play-j/backend/internal/design/player-service
// -o ./internal/playerservice

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// MusicFileResponse is the viewed result type that is projected based on a
// view.
type MusicFileResponse struct {
	// Type to project
	Projected *MusicFileResponseView
	// View to render
	View string
}

// MusicFileResponseView is a type that runs validations on a projected type.
type MusicFileResponseView struct {
	File []byte
}

var (
	// MusicFileResponseMap is a map indexing the attribute names of
	// MusicFileResponse by view name.
	MusicFileResponseMap = map[string][]string{
		"default": {
			"file",
		},
	}
)

// ValidateMusicFileResponse runs the validations defined on the viewed result
// type MusicFileResponse.
func ValidateMusicFileResponse(result *MusicFileResponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateMusicFileResponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateMusicFileResponseView runs the validations defined on
// MusicFileResponseView using the "default" view.
func ValidateMusicFileResponseView(result *MusicFileResponseView) (err error) {

	return
}
