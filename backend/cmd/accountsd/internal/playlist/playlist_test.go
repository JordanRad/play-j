package playlist_test

import (
	"context"
	"time"

	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/playlist"
	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/playlist/playlistfakes"
	playlistgen "github.com/JordanRad/play-j/backend/internal/accountservice/gen/playlist"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

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
		Describe("Given a valid JWT", func() {
			var (
				ctx      context.Context
				response *playlistgen.AccountPlaylistCollectionResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
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

		Describe("Given an INVALID JWT", func() {
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
		Describe("Given a valid JWT", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
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

		Describe("Given an INVALID JWT", func() {
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
		Describe("Given a valid JWT", func() {
			var (
				ctx      context.Context
				response *playlistgen.AccountPlaylistResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
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

		Describe("Given an INVALID JWT", func() {
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
		Describe("Given a valid JWT", func() {
			var (
				ctx      context.Context
				response *playlistgen.PlaylistModificationResponse
				err      error
			)

			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, "jwt", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImFjY291bnRJRCI6IjEiLCJlbWFpbCI6ImpvZXNtaXRoQGFidi5iZyJ9LCJleHAiOjM5OTk5OTE1MjgsImlzcyI6InBsYXlqLWFjY291bnQtc2VydmljZSJ9.fMaYAI_qnb97cfI7bsYT-WK8U9FAT0_DGXLJi5ejmPA")
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

		Describe("Given an INVALID JWT", func() {
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
})
