package dbmodels

import "time"

type Account struct {
	ID                    uint
	FirstName             string
	LastName              string
	Username              string
	Email                 string
	Password              string
	IsActive              bool
	HasActiveSubscription bool
	Playlists             []Playlist
}

type Playlist struct {
	ID        uint
	Name      string
	AccountID uint
	CreatedAt time.Time
	TrackIDs  []int32
}
