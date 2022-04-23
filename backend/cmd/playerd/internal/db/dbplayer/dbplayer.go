package dbplayer

import (
	"database/sql"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/player"
)

type Store struct {
	DB *sql.DB
}

var _ player.Store = (*Store)(nil)

func (s *Store) GetTrackByID(id uint) (*dbmodels.Track, error) {
	var result dbmodels.Track

	row := s.DB.QueryRow(`SELECT * FROM tracks WHERE id = $1;`, id)
	err := row.Scan(&result.ID, &result.Name, &result.FullName,
		&result.StorageTrackID, &result.CreatedAt, &result.ArtistID, &result.AlbumID)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}

	return &result, nil
}
