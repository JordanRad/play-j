package dbplayer

import (
	"database/sql"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/player"
)

type Store struct {
	DB *sql.DB
}

var _ player.Store = (*Store)(nil)

func (s *Store) GetTrackByID(id uint) (*dbmodels.Track, error) {
	return nil, nil
}
