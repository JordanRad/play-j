package dbmodels

import "time"

type Artist struct {
	ID                uint
	Name              string
	StorageFolderName string
	Albums            []Album
}

type Album struct {
	ID        uint
	Name      string
	Genre     string
	CreatedAt time.Time
	TrackIDs  []int32
	ArtistID  uint
}

type Track struct {
	ID             uint
	Name           string
	FullName       string
	StorageTrackID string
	CreatedAt      time.Time
	AlbumID        uint
	ArtistID       uint
}
