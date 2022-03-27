package dbmodels

type Playlist struct {
	ID     uint
	Name   string
	Tracks []string
}
type Account struct {
	ID        uint
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Playlists []Playlist
}
