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
	ID              uint
	Name            string
	FullName        string
	StorageLocation string
	CreatedAt       time.Time
	AlbumID         uint
	ArtistID        uint
}

type PlayerSearch struct {
	TrackID    uint
	TrackName  string
	ArtistID   uint
	ArtistName string
	AlbumID    uint
	AlbumName  string
}
