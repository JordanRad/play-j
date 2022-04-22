package player

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/JordanRad/play-j/backend/cmd/playerd/internal/db/dbmodels"
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

var _ PlayerService = (*Service)(nil)

func (s *Service) StreamTrackByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	musicFile, err := s.cloudStorage.ReadFileFromFolder(ctx, "Metallica", "One.mp3")

	if err != nil {
		fmt.Printf("error getting the file: %v", err)
		w.Write([]byte("Cannot find the file: "))
	}

	fmt.Println(len(musicFile))

	flusher, ok := w.(http.Flusher)

	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}

	w.Header().Set("Content-Type", "audio/mpegurl")
	w.Header().Set("Transfer-encoding", "chunked")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	for i := 0; i < 5; i++ {
		start := (i * len(musicFile) / 5)
		end := ((i + 1) * len(musicFile)) / 5

		fmt.Printf("Chunk #%v sent...\n", i+1)
		w.Write(musicFile[start:end])
		flusher.Flush() // Trigger "chunked" encoding and send a chunk...
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Service) Search(w http.ResponseWriter, r *http.Request) {

}
