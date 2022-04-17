package player

import (
	"context"
	"fmt"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
	"github.com/JordanRad/play-j/backend/internal/playerservice/gen/player"
)

type Store interface {
	GetTrackByID(id uint) (*dbmodels.Track, error)
}

type CloudStorage interface {
	ReadFileFromFolder(ctx context.Context, folderName string, fileName string) ([]byte, error)
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
	track, err := s.cloudStorage.ReadFileFromFolder(ctx, "Metallica", "One")
	if err != nil {
		return nil, fmt.Errorf("error reading the file from the bucket: %w", err)
	}

	response := &player.StreamTrackResponse{
		Track: track,
	}
	return response, nil
}
