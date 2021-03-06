package playlist_test

import (
	"context"
	"fmt"
	"time"

	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/playlist"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/playlist/playlistfakes"
	playlistgen "github.com/JordanRad/play-j/backend/internal/accountservice/gen/playlist"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var TestToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA"
var _ = Describe("Playlist Service", func() {

	var (
		fakeStore *playlistfakes.FakeStore
		service   *playlist.Service
	)

	BeforeEach(func() {
		fakeStore = new(playlistfakes.FakeStore)
		service = playlist.NewService(fakeStore)
	})

	Describe("GetAccountPlaylistCollection", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *playlistgen.AccountPlaylistCollectionResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &playlistgen.GetAccountPlaylistCollectionPayload{
					Auth: "Bearer token",
				}
				response, err = service.GetAccountPlaylistCollection(ctx, payload)
			})

			When("account's playlists are extracted susccessfully", func() {
				BeforeEach(func() {
					var playlists []*dbmodels.Playlist
					playlist1 := &dbmodels.Playlist{
						ID:        1,
						Name:      "Hits 2022",
						AccountID: 1,
						TrackIDs:  make([]int32, 0),
						CreatedAt: time.Now(),
					}
					playlist2 := &dbmodels.Playlist{
						ID:        2,
						Name:      "Winter 2021",
						AccountID: 1,
						TrackIDs:  make([]int32, 0),
						CreatedAt: time.Now(),
					}
					playlists = append(playlists, playlist1, playlist2)
					fakeStore.GetAllAccountPlaylistsReturns(playlists, nil)
				})

				It("should return a collection of playlists", func() {
					Expect(response.Total).To(Equal(int32(2)))
					Expect(response.Resources[0].Name).To(Equal("Hits 2022"))
					Expect(response.Resources[1].Name).To(Equal("Winter 2021"))
				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})
		})

		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *playlistgen.AccountPlaylistCollectionResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &playlistgen.GetAccountPlaylistCollectionPayload{
					Auth: "Bearer token",
				}
				response, err = service.GetAccountPlaylistCollection(ctx, payload)
			})

			When("operation is unathorized", func() {

				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("CreateAccountPlaylist", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &playlistgen.CreateAccountPlaylistPayload{
					Auth: "Bearer token",
					Name: "New test playlist",
				}
				response, err = service.CreateAccountPlaylist(ctx, payload)
			})

			When("playlist is created susccessfully", func() {
				BeforeEach(func() {
					fakeStore.CreateAccountPlaylistReturns(true, nil)
				})

				It("should return a success message", func() {
					Expect(response.Message).To(Equal("Playlist created successfully"))

				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})
		})

		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &playlistgen.CreateAccountPlaylistPayload{
					Auth: "Bearer token",
					Name: "Test Playlist",
				}
				response, err = service.CreateAccountPlaylist(ctx, payload)
			})

			When("operation is unathorized", func() {

				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("GetAccountPlaylist", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *playlistgen.AccountPlaylistResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &playlistgen.GetAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
				}
				response, err = service.GetAccountPlaylist(ctx, payload)
			})

			When("playlist is extracted susccessfully", func() {
				var trackIDs []int32
				var playlist *dbmodels.Playlist
				BeforeEach(func() {
					trackIDs = make([]int32, 0)
					trackIDs = append(trackIDs, 101, 102, 302)
					playlist = &dbmodels.Playlist{
						ID:        1,
						Name:      "Hits 2022",
						AccountID: 1,
						TrackIDs:  trackIDs,
						CreatedAt: time.Now(),
					}
					fakeStore.GetAccountPlaylistReturns(playlist, nil)
				})

				It("should return a collection of playlists", func() {
					Expect(response.ID).To(Equal(int32(playlist.ID)))
					Expect(response.Name).To(Equal(playlist.Name))
					Expect(response.TrackIDs).To(Equal(playlist.TrackIDs))

				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})
		})

		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *playlistgen.AccountPlaylistResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &playlistgen.GetAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
				}
				response, err = service.GetAccountPlaylist(ctx, payload)
			})

			When("operation is unathorized", func() {

				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("RenameAccountPlaylist", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &playlistgen.RenameAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
					Name:       "Updated Name",
				}
				response, err = service.RenameAccountPlaylist(ctx, payload)
			})

			When("playlists is renamed susccessfully", func() {
				BeforeEach(func() {
					fakeStore.UpdateAccountPlaylistNameReturns(true, nil)
				})

				It("should return a success message", func() {
					Expect(response.Message).To(Equal("Playlist name updated successfully"))
				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})

		})

		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &playlistgen.RenameAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
					Name:       "wrong",
				}
				response, err = service.RenameAccountPlaylist(ctx, payload)
			})

			When("operation is unathorized", func() {

				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("AddTrackToAccountPlaylist", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &playlistgen.AddTrackToAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
					TrackID:    101,
				}
				response, err = service.AddTrackToAccountPlaylist(ctx, payload)
			})

			When("track is added to a playlist susccessfully", func() {
				BeforeEach(func() {
					fakeStore.UpdateAccountPlaylistTracksReturns(true, nil)
				})

				It("should return a success message", func() {
					Expect(response.Message).To(Equal(fmt.Sprintf("Track with ID: %d has been added to Playlist with ID: %d successfully", 101, 1)))
				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})

		})

		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &playlistgen.AddTrackToAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
					TrackID:    101,
				}
				response, err = service.AddTrackToAccountPlaylist(ctx, payload)
			})

			When("operation is unathorized", func() {

				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("RemoveTrackFromAccountPlaylist", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &playlistgen.RemoveTrackFromAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
					TrackID:    101,
				}
				response, err = service.RemoveTrackFromAccountPlaylist(ctx, payload)
			})

			When("track is removed from a playlist susccessfully", func() {
				BeforeEach(func() {
					fakeStore.UpdateAccountPlaylistTracksReturns(true, nil)
				})

				It("should return a success message", func() {
					Expect(response.Message).To(Equal(fmt.Sprintf("Track with ID: %d has been removed from Playlist with ID: %d successfully", 101, 1)))
				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})

		})

		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &playlistgen.RemoveTrackFromAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
					TrackID:    101,
				}
				response, err = service.RemoveTrackFromAccountPlaylist(ctx, payload)
			})

			When("operation is unathorized", func() {

				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("DeleteAccountPlaylist", func() {
		Describe("on given a valid token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", TestToken)
				payload := &playlistgen.DeleteAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
				}
				response, err = service.DeleteAccountPlaylist(ctx, payload)
			})

			When("track is removed from a playlist susccessfully", func() {
				BeforeEach(func() {
					fakeStore.DeleteAccountPlaylistReturns(true, nil)
				})

				It("should return a success message", func() {
					Expect(response.Message).To(Equal("Playlist deleted successfully"))
				})
				It("should NOT return any error", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})

		})

		Describe("on given an INVALID token", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGpXVCJ9.eyJkYX9.fMfUPA")
				payload := &playlistgen.DeleteAccountPlaylistPayload{
					Auth:       "Bearer token",
					PlaylistID: 1,
				}
				response, err = service.DeleteAccountPlaylist(ctx, payload)
			})

			When("operation is unathorized", func() {

				It("should return an error", func() {
					Expect(response).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})
})
