package dbplayer

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/player"
)

type Store struct {
	DB *sql.DB
}

var _ player.Store = (*Store)(nil)

func (s *Store) GetTrackByID(ctx context.Context, id uint) (*dbmodels.Track, error) {
	var result dbmodels.Track

	row := s.DB.QueryRow(`SELECT * FROM tracks WHERE id = $1;`, id)
	err := row.Scan(&result.ID, &result.Name, &result.FullName,
		&result.StorageLocation, &result.CreatedAt, &result.ArtistID, &result.AlbumID)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("invalid credentials: %w", err)
	}

	return &result, nil
}

func (s Store) Search(ctx context.Context, searchInput string, limit uint) ([]*dbmodels.PlayerSearch, error) {
	query := `
	select t.id as track_id, initcap(t.name) as track_name , ar.id as artist_id, initcap(ar.name) as artist_name,al.id as album_id, initcap(al.name) as album_name
	from tracks t 
	join artists ar on t.artistid = ar.id 
	join albums al on t.albumid = al.id 
	where lower(t.name)  LIKE '%' || $1 || '%'
	or lower(t.fullname) LIKE '%' || $1 || '%'
	or lower(al.name) LIKE '%' || $1 || '%'
	or lower(ar.name) LIKE '%' || $1 || '%'
	limit $2;
	`
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing the query: %w\n", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(searchInput, limit)

	if err != nil {
		return nil, fmt.Errorf("error extracting accounts's payments: %w", err)
	}
	defer rows.Close()

	searchResults := make([]*dbmodels.PlayerSearch, 0)
	for rows.Next() {
		searchResult := &dbmodels.PlayerSearch{}

		err := rows.Scan(
			&searchResult.TrackID,
			&searchResult.TrackName,
			&searchResult.ArtistID,
			&searchResult.ArtistName,
			&searchResult.AlbumID,
			&searchResult.AlbumName,
		)

		if err != nil {
			return nil, fmt.Errorf("error mapping a payment row: %w", err)
		}

		searchResults = append(searchResults, searchResult)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return searchResults, nil
}
