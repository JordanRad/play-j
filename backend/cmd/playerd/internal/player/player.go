package player

import (
	"context"

	"github.com/JordanRad/play-j/backend/internal/playerservice/gen/player"
)

type Store interface {
}

type CloudStorage interface {
}

type Service struct {
	store        Store
	cloudStorage CloudStorage
}

func NewService(store Store, cloudStorage CloudStorage) *Service {
	return &Service{
		store:        store,
		cloudStorage: cloudStorage,
	}
}

var _ player.Service = (*Service)(nil)

func (s *Service) GetTrackByID(ctx context.Context, p *player.GetTrackByIDPayload) (*player.StreamTrackResponse, error) {
	return nil, nil
}
