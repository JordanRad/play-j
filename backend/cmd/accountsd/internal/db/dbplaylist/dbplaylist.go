package dbplaylist

import (
	"context"
	"database/sql"
	"fmt"

	account "git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/account"
	"git.fhict.nl/I425652/jordan-portfolio-s6/backend/cmd/accountsd/internal/db/dbmodels"
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
		return true, nil
	}

	return false, fmt.Errorf("error inserting a playlist")
}

func (s *Store) DeleteAccountPlaylist(ctx context.Context, accountID uint, playlistID uint) (bool, error) {
	result, err := s.DB.Exec("DELETE FROM playlists WHERE id= $1 and accountid = $2;", playlistID, accountID)

	if err != nil {
		return false, fmt.Errorf("error deleting a playlist: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 1 {
		return true, nil
	}

	return false, fmt.Errorf("error removing a playlist")
}

func (s *Store) GetUserPlaylist(ctx context.Context, accountID uint, playlistID uint) (string, error) {

	return "", nil
}

func (s *Store) GetAllUserPlaylists(ctx context.Context, accountID uint) ([]*dbmodels.Playlist, error) {
	rows, err := s.DB.Query("SELECT * FROM playlists WHERE accountId = $1;", accountID)

	if err != nil {
		return make([]*dbmodels.Playlist, 0), fmt.Errorf("error extracting user's playlists: %w", err)
	}
	defer rows.Close()

	playlists := make([]*dbmodels.Playlist, 0)
	for rows.Next() {
		playlist := &dbmodels.Playlist{}
		tracks := []string{"song1", "song2"}
		err := rows.Scan(&playlist.ID, &playlist.Name, &playlist.AccountID)
		if err != nil {
			return make([]*dbmodels.Playlist, 0), fmt.Errorf("error mapping a playlist row: %w", err)
		}
		playlist.Tracks = tracks
		playlists = append(playlists, playlist)
	}
	if err := rows.Err(); err != nil {
		return make([]*dbmodels.Playlist, 0), fmt.Errorf("rows err: %w", err)
	}
	return playlists, nil
}
