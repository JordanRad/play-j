// Code generated by goa v3.7.0, DO NOT EDIT.
//
// accounts HTTP client CLI support package
//
// Command:
// $ goa gen
// github.com/JordanRad/play-j/backend/internal/design/account-service -o
// ./internal/accountservice

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	accountc "github.com/JordanRad/play-j/backend/internal/accountservice/gen/http/account/client"
	playlistc "github.com/JordanRad/play-j/backend/internal/accountservice/gen/http/playlist/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `account (register|login)
playlist (get-account-playlist-collection|create-account-playlist|rename-account-playlist|delete-account-playlist|get-account-playlist|add-track-to-account-playlist|remove-track-from-account-playlist)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` account register --body '{
      "confirmedPassword": "Numquam quos excepturi vero ad est.",
      "email": "Dolorum et labore cumque quisquam dolorem adipisci.",
      "firstName": "Quibusdam omnis nemo provident eos quis ut.",
      "lastName": "Ipsum et molestiae.",
      "password": "Voluptates id recusandae temporibus et dolore."
   }'` + "\n" +
		os.Args[0] + ` playlist get-account-playlist-collection --auth "Illum cupiditate corporis aut vero."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		accountFlags = flag.NewFlagSet("account", flag.ContinueOnError)

		accountRegisterFlags    = flag.NewFlagSet("register", flag.ExitOnError)
		accountRegisterBodyFlag = accountRegisterFlags.String("body", "REQUIRED", "")

		accountLoginFlags    = flag.NewFlagSet("login", flag.ExitOnError)
		accountLoginBodyFlag = accountLoginFlags.String("body", "REQUIRED", "")

		playlistFlags = flag.NewFlagSet("playlist", flag.ContinueOnError)

		playlistGetAccountPlaylistCollectionFlags    = flag.NewFlagSet("get-account-playlist-collection", flag.ExitOnError)
		playlistGetAccountPlaylistCollectionAuthFlag = playlistGetAccountPlaylistCollectionFlags.String("auth", "", "")

		playlistCreateAccountPlaylistFlags    = flag.NewFlagSet("create-account-playlist", flag.ExitOnError)
		playlistCreateAccountPlaylistBodyFlag = playlistCreateAccountPlaylistFlags.String("body", "REQUIRED", "")
		playlistCreateAccountPlaylistAuthFlag = playlistCreateAccountPlaylistFlags.String("auth", "", "")

		playlistRenameAccountPlaylistFlags          = flag.NewFlagSet("rename-account-playlist", flag.ExitOnError)
		playlistRenameAccountPlaylistBodyFlag       = playlistRenameAccountPlaylistFlags.String("body", "REQUIRED", "")
		playlistRenameAccountPlaylistPlaylistIDFlag = playlistRenameAccountPlaylistFlags.String("playlist-id", "REQUIRED", "Playlist id to modify")
		playlistRenameAccountPlaylistAuthFlag       = playlistRenameAccountPlaylistFlags.String("auth", "", "")

		playlistDeleteAccountPlaylistFlags          = flag.NewFlagSet("delete-account-playlist", flag.ExitOnError)
		playlistDeleteAccountPlaylistPlaylistIDFlag = playlistDeleteAccountPlaylistFlags.String("playlist-id", "REQUIRED", "")
		playlistDeleteAccountPlaylistAuthFlag       = playlistDeleteAccountPlaylistFlags.String("auth", "", "")

		playlistGetAccountPlaylistFlags          = flag.NewFlagSet("get-account-playlist", flag.ExitOnError)
		playlistGetAccountPlaylistPlaylistIDFlag = playlistGetAccountPlaylistFlags.String("playlist-id", "REQUIRED", "Playlist ID")
		playlistGetAccountPlaylistAuthFlag       = playlistGetAccountPlaylistFlags.String("auth", "", "")

		playlistAddTrackToAccountPlaylistFlags          = flag.NewFlagSet("add-track-to-account-playlist", flag.ExitOnError)
		playlistAddTrackToAccountPlaylistPlaylistIDFlag = playlistAddTrackToAccountPlaylistFlags.String("playlist-id", "REQUIRED", "Playlist ID to modify")
		playlistAddTrackToAccountPlaylistTrackIDFlag    = playlistAddTrackToAccountPlaylistFlags.String("track-id", "REQUIRED", "Track ID to be added")
		playlistAddTrackToAccountPlaylistAuthFlag       = playlistAddTrackToAccountPlaylistFlags.String("auth", "", "")

		playlistRemoveTrackFromAccountPlaylistFlags          = flag.NewFlagSet("remove-track-from-account-playlist", flag.ExitOnError)
		playlistRemoveTrackFromAccountPlaylistPlaylistIDFlag = playlistRemoveTrackFromAccountPlaylistFlags.String("playlist-id", "REQUIRED", "Playlist ID to modify")
		playlistRemoveTrackFromAccountPlaylistTrackIDFlag    = playlistRemoveTrackFromAccountPlaylistFlags.String("track-id", "REQUIRED", "Track ID to be deleted")
		playlistRemoveTrackFromAccountPlaylistAuthFlag       = playlistRemoveTrackFromAccountPlaylistFlags.String("auth", "", "")
	)
	accountFlags.Usage = accountUsage
	accountRegisterFlags.Usage = accountRegisterUsage
	accountLoginFlags.Usage = accountLoginUsage

	playlistFlags.Usage = playlistUsage
	playlistGetAccountPlaylistCollectionFlags.Usage = playlistGetAccountPlaylistCollectionUsage
	playlistCreateAccountPlaylistFlags.Usage = playlistCreateAccountPlaylistUsage
	playlistRenameAccountPlaylistFlags.Usage = playlistRenameAccountPlaylistUsage
	playlistDeleteAccountPlaylistFlags.Usage = playlistDeleteAccountPlaylistUsage
	playlistGetAccountPlaylistFlags.Usage = playlistGetAccountPlaylistUsage
	playlistAddTrackToAccountPlaylistFlags.Usage = playlistAddTrackToAccountPlaylistUsage
	playlistRemoveTrackFromAccountPlaylistFlags.Usage = playlistRemoveTrackFromAccountPlaylistUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "account":
			svcf = accountFlags
		case "playlist":
			svcf = playlistFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "account":
			switch epn {
			case "register":
				epf = accountRegisterFlags

			case "login":
				epf = accountLoginFlags

			}

		case "playlist":
			switch epn {
			case "get-account-playlist-collection":
				epf = playlistGetAccountPlaylistCollectionFlags

			case "create-account-playlist":
				epf = playlistCreateAccountPlaylistFlags

			case "rename-account-playlist":
				epf = playlistRenameAccountPlaylistFlags

			case "delete-account-playlist":
				epf = playlistDeleteAccountPlaylistFlags

			case "get-account-playlist":
				epf = playlistGetAccountPlaylistFlags

			case "add-track-to-account-playlist":
				epf = playlistAddTrackToAccountPlaylistFlags

			case "remove-track-from-account-playlist":
				epf = playlistRemoveTrackFromAccountPlaylistFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "account":
			c := accountc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "register":
				endpoint = c.Register()
				data, err = accountc.BuildRegisterPayload(*accountRegisterBodyFlag)
			case "login":
				endpoint = c.Login()
				data, err = accountc.BuildLoginPayload(*accountLoginBodyFlag)
			}
		case "playlist":
			c := playlistc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-account-playlist-collection":
				endpoint = c.GetAccountPlaylistCollection()
				data, err = playlistc.BuildGetAccountPlaylistCollectionPayload(*playlistGetAccountPlaylistCollectionAuthFlag)
			case "create-account-playlist":
				endpoint = c.CreateAccountPlaylist()
				data, err = playlistc.BuildCreateAccountPlaylistPayload(*playlistCreateAccountPlaylistBodyFlag, *playlistCreateAccountPlaylistAuthFlag)
			case "rename-account-playlist":
				endpoint = c.RenameAccountPlaylist()
				data, err = playlistc.BuildRenameAccountPlaylistPayload(*playlistRenameAccountPlaylistBodyFlag, *playlistRenameAccountPlaylistPlaylistIDFlag, *playlistRenameAccountPlaylistAuthFlag)
			case "delete-account-playlist":
				endpoint = c.DeleteAccountPlaylist()
				data, err = playlistc.BuildDeleteAccountPlaylistPayload(*playlistDeleteAccountPlaylistPlaylistIDFlag, *playlistDeleteAccountPlaylistAuthFlag)
			case "get-account-playlist":
				endpoint = c.GetAccountPlaylist()
				data, err = playlistc.BuildGetAccountPlaylistPayload(*playlistGetAccountPlaylistPlaylistIDFlag, *playlistGetAccountPlaylistAuthFlag)
			case "add-track-to-account-playlist":
				endpoint = c.AddTrackToAccountPlaylist()
				data, err = playlistc.BuildAddTrackToAccountPlaylistPayload(*playlistAddTrackToAccountPlaylistPlaylistIDFlag, *playlistAddTrackToAccountPlaylistTrackIDFlag, *playlistAddTrackToAccountPlaylistAuthFlag)
			case "remove-track-from-account-playlist":
				endpoint = c.RemoveTrackFromAccountPlaylist()
				data, err = playlistc.BuildRemoveTrackFromAccountPlaylistPayload(*playlistRemoveTrackFromAccountPlaylistPlaylistIDFlag, *playlistRemoveTrackFromAccountPlaylistTrackIDFlag, *playlistRemoveTrackFromAccountPlaylistAuthFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// accountUsage displays the usage of the account command and its subcommands.
func accountUsage() {
	fmt.Fprintf(os.Stderr, `Account operations
Usage:
    %[1]s [globalflags] account COMMAND [flags]

COMMAND:
    register: Register implements register.
    login: Login implements login.

Additional help:
    %[1]s account COMMAND --help
`, os.Args[0])
}
func accountRegisterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account register -body JSON

Register implements register.
    -body JSON: 

Example:
    %[1]s account register --body '{
      "confirmedPassword": "Numquam quos excepturi vero ad est.",
      "email": "Dolorum et labore cumque quisquam dolorem adipisci.",
      "firstName": "Quibusdam omnis nemo provident eos quis ut.",
      "lastName": "Ipsum et molestiae.",
      "password": "Voluptates id recusandae temporibus et dolore."
   }'
`, os.Args[0])
}

func accountLoginUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] account login -body JSON

Login implements login.
    -body JSON: 

Example:
    %[1]s account login --body '{
      "email": "Iusto non mollitia qui non culpa laborum.",
      "password": "Corrupti voluptas officia nostrum quia voluptatum."
   }'
`, os.Args[0])
}

// playlistUsage displays the usage of the playlist command and its subcommands.
func playlistUsage() {
	fmt.Fprintf(os.Stderr, `Playlist operations
Usage:
    %[1]s [globalflags] playlist COMMAND [flags]

COMMAND:
    get-account-playlist-collection: GetAccountPlaylistCollection implements getAccountPlaylistCollection.
    create-account-playlist: CreateAccountPlaylist implements createAccountPlaylist.
    rename-account-playlist: RenameAccountPlaylist implements renameAccountPlaylist.
    delete-account-playlist: DeleteAccountPlaylist implements deleteAccountPlaylist.
    get-account-playlist: GetAccountPlaylist implements getAccountPlaylist.
    add-track-to-account-playlist: AddTrackToAccountPlaylist implements addTrackToAccountPlaylist.
    remove-track-from-account-playlist: RemoveTrackFromAccountPlaylist implements removeTrackFromAccountPlaylist.

Additional help:
    %[1]s playlist COMMAND --help
`, os.Args[0])
}
func playlistGetAccountPlaylistCollectionUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] playlist get-account-playlist-collection -auth STRING

GetAccountPlaylistCollection implements getAccountPlaylistCollection.
    -auth STRING: 

Example:
    %[1]s playlist get-account-playlist-collection --auth "Illum cupiditate corporis aut vero."
`, os.Args[0])
}

func playlistCreateAccountPlaylistUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] playlist create-account-playlist -body JSON -auth STRING

CreateAccountPlaylist implements createAccountPlaylist.
    -body JSON: 
    -auth STRING: 

Example:
    %[1]s playlist create-account-playlist --body '{
      "name": "Cum corporis autem repellendus."
   }' --auth "Non veniam est."
`, os.Args[0])
}

func playlistRenameAccountPlaylistUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] playlist rename-account-playlist -body JSON -playlist-id UINT -auth STRING

RenameAccountPlaylist implements renameAccountPlaylist.
    -body JSON: 
    -playlist-id UINT: Playlist id to modify
    -auth STRING: 

Example:
    %[1]s playlist rename-account-playlist --body '{
      "name": "Voluptatem sequi adipisci iure dolorem nesciunt itaque."
   }' --playlist-id 5048963703141644532 --auth "Praesentium commodi ratione."
`, os.Args[0])
}

func playlistDeleteAccountPlaylistUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] playlist delete-account-playlist -playlist-id UINT -auth STRING

DeleteAccountPlaylist implements deleteAccountPlaylist.
    -playlist-id UINT: 
    -auth STRING: 

Example:
    %[1]s playlist delete-account-playlist --playlist-id 16587001386555198618 --auth "Qui qui sed."
`, os.Args[0])
}

func playlistGetAccountPlaylistUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] playlist get-account-playlist -playlist-id UINT -auth STRING

GetAccountPlaylist implements getAccountPlaylist.
    -playlist-id UINT: Playlist ID
    -auth STRING: 

Example:
    %[1]s playlist get-account-playlist --playlist-id 9750426786919826982 --auth "Eius et quasi repudiandae."
`, os.Args[0])
}

func playlistAddTrackToAccountPlaylistUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] playlist add-track-to-account-playlist -playlist-id UINT -track-id UINT -auth STRING

AddTrackToAccountPlaylist implements addTrackToAccountPlaylist.
    -playlist-id UINT: Playlist ID to modify
    -track-id UINT: Track ID to be added
    -auth STRING: 

Example:
    %[1]s playlist add-track-to-account-playlist --playlist-id 7897115789263760436 --track-id 14345441535750740261 --auth "Deserunt consectetur ipsa ab quibusdam enim ratione."
`, os.Args[0])
}

func playlistRemoveTrackFromAccountPlaylistUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] playlist remove-track-from-account-playlist -playlist-id UINT -track-id UINT -auth STRING

RemoveTrackFromAccountPlaylist implements removeTrackFromAccountPlaylist.
    -playlist-id UINT: Playlist ID to modify
    -track-id UINT: Track ID to be deleted
    -auth STRING: 

Example:
    %[1]s playlist remove-track-from-account-playlist --playlist-id 8873083792410562693 --track-id 7249225158853866608 --auth "A libero dolore aliquid voluptas ut."
`, os.Args[0])
}
