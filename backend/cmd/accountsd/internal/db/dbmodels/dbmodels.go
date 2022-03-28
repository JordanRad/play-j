package dbmodels

type Account struct {
	ID        uint
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Playlists []Playlist
}

type Playlist struct {
	ID        uint
	Name      string
	AccountID uint
	Tracks    []string
}