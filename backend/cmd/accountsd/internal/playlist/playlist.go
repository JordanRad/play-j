// package playlist

// import (
// 	"context"
// 	"errors"
// 	"fmt"

// 	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
// 	auth "github.com/JordanRad/play-j/backend/internal/auth"
// 	"github.com/JordanRad/play-j/backend/internal/gen/account"
// )

// type PlaylistStore interface {
// 	CreateAccountPlaylist(context.Context, uint, string) (bool, error)
// 	GetUserPlaylist(context.Context, uint, uint) (string, error)
// 	GetAllUserPlaylists(context.Context, uint) ([]*dbmodels.Playlist, error)
// 	DeleteAccountPlaylist(context.Context, uint, uint) (bool, error)
// }

// type Service struct {
// 	playlistStore PlaylistStore
// }