package playlist

import (
	"context"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/accountsd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/internal/accountservice/gen/playlist"
	"github.com/JordanRad/play-j/backend/internal/auth"
)

// You only need **one** of these per package!
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Store

type Store interface {
	CreateAccountPlaylist(context.Context, uint, string) (bool, error)
	GetAccountPlaylist(context.Context, uint, uint) (*dbmodels.Playlist, error)
	GetAllAccountPlaylists(context.Context, uint) ([]*dbmodels.Playlist, error)
	DeleteAccountPlaylist(context.Context, uint, uint) (bool, error)
	UpdateAccountPlaylistTracks(context.Context, uint, uint, string) (bool, error)
	UpdateAccountPlaylistName(context.Context, uint, string) (bool, error)
}

type Service struct {
	store Store
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
	}
}

// Compile-time assertion that *Service implements the user.Service interface.
var _ playlist.Service = (*Service)(nil)

func (s *Service) GetAccountPlaylistCollection(ctx context.Context, p *playlist.GetAccountPlaylistCollectionPayload) (*playlist.AccountPlaylistCollectionResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	// Get playlists by user
	playlists, err := s.store.GetAllAccountPlaylists(ctx, tokenClaims.AccountID)

	if err != nil {
		return &playlist.AccountPlaylistCollectionResponse{}, fmt.Errorf("error displaying user's playlists: %w", err)
	}

	resources := make([]*playlist.AccountPlaylistResponse, 0, len(playlists))

	for _, singlePlaylist := range playlists {

		resource := &playlist.AccountPlaylistResponse{
			ID:       int32(singlePlaylist.ID),
			Name:     singlePlaylist.Name,
			TrackIDs: singlePlaylist.TrackIDs,
		}

		resources = append(resources, resource)
	}

	response := &playlist.AccountPlaylistCollectionResponse{
		Total:     int32(len(playlists)),
		Resources: resources,
	}

	return response, nil
}

func (s *Service) CreateAccountPlaylist(ctx context.Context, p *playlist.CreateAccountPlaylistPayload) (*playlist.PlaylistModificationResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}
	_, err = s.store.CreateAccountPlaylist(ctx, tokenClaims.AccountID, p.Name)
	if err != nil {
		return nil, fmt.Errorf("error creating new playlist: %w", err)
	}

	response := &playlist.PlaylistModificationResponse{
		Message: "Playlist created successfully",
	}

	return response, nil
}

func (s *Service) DeleteAccountPlaylist(ctx context.Context, p *playlist.DeleteAccountPlaylistPayload) (*playlist.PlaylistModificationResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	_, err = s.store.DeleteAccountPlaylist(ctx, tokenClaims.AccountID, p.PlaylistID)
	if err != nil {
		return nil, fmt.Errorf("error deleting a playlist: %w", err)
	}

	response := &playlist.PlaylistModificationResponse{
		Message: "Playlist deleted successfully",
	}

	return response, nil
}

func (s *Service) GetAccountPlaylist(ctx context.Context, p *playlist.GetAccountPlaylistPayload) (*playlist.AccountPlaylistResponse, error) {
	// Extract claims from token
	tokenClaims, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	singlePlaylist, err := s.store.GetAccountPlaylist(ctx, tokenClaims.AccountID, p.PlaylistID)

	if err != nil {
		return nil, fmt.Errorf("[service] error getting a playlis: %w", err)
	}

	response := &playlist.AccountPlaylistResponse{
		ID:       int32(singlePlaylist.ID),
		Name:     singlePlaylist.Name,
		TrackIDs: singlePlaylist.TrackIDs,
	}

	return response, nil
}

func (s *Service) RenameAccountPlaylist(ctx context.Context, p *playlist.RenameAccountPlaylistPayload) (*playlist.PlaylistModificationResponse, error) {
	// Extract claims from token
	_, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}
	_, err = s.store.UpdateAccountPlaylistName(ctx, p.PlaylistID, p.Name)

	if err != nil {
		return nil, fmt.Errorf("error renaimg a playlist: %w", err)
	}

	response := &playlist.PlaylistModificationResponse{
		Message: "Playlist name updated successfully",
	}

	return response, nil
}

func (s *Service) AddTrackToAccountPlaylist(ctx context.Context, p *playlist.AddTrackToAccountPlaylistPayload) (*playlist.PlaylistModificationResponse, error) {
	// Extract claims from token
	_, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	_, err = s.store.UpdateAccountPlaylistTracks(ctx, p.PlaylistID, p.TrackID, "ADD")

	if err != nil {
		return nil, fmt.Errorf("service error adding track to a playlist: %w", err)
	}

	response := &playlist.PlaylistModificationResponse{
		Message: fmt.Sprintf("Track with ID: %d has been added to Playlist with ID: %d successfully", p.TrackID, p.PlaylistID),
	}

	return response, nil
}

func (s *Service) RemoveTrackFromAccountPlaylist(ctx context.Context, p *playlist.RemoveTrackFromAccountPlaylistPayload) (*playlist.PlaylistModificationResponse, error) {
	// Extract claims from token
	_, err := auth.ExtractJWTCLaims(ctx.Value("jwt").(string))
	if err != nil {
		return nil, fmt.Errorf("error extracting token claims: %w", err)
	}

	_, err = s.store.UpdateAccountPlaylistTracks(ctx, p.PlaylistID, p.TrackID, "REMOVE")

	if err != nil {
		return nil, fmt.Errorf("service error removing a track from a playlist: %w", err)
	}

	response := &playlist.PlaylistModificationResponse{
		Message: fmt.Sprintf("Track with ID: %d has been removed from Playlist with ID: %d successfully", p.TrackID, p.PlaylistID),
	}

	return response, nil
}
