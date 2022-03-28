package dbplaylist

import (
	"context"
	"database/sql"
	"fmt"

	account "git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/account"
)

type Store struct {
	DB *sql.DB
}

var _ account.PlaylistStore = (*Store)(nil)

func (s *Store) CreateAccountPlaylist(ctx context.Context, accountID uint, playlistName string) (bool, error) {
	result, err := s.DB.Exec("INSERT INTO playlists (accountid,name) VALUES ($1,$2);", accountID, playlistName)

	if err != nil {
		return false, fmt.Errorf("error creating a playlist: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 1 {
		return true, fmt.Errorf("error inserting the user")
	}
	return false, nil
}

func (s *Store) GetUserPlaylist(ctx context.Context, accountID uint, playlistID uint) (string, error) {
	return "", nil
}

func (s *Store) GetAllUserPlaylists(ctx context.Context, accountID uint) ([]string, error) {
	return make([]string, 0), nil
}
