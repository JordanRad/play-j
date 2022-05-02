package cloudstorage

import (
	"context"
	"fmt"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
)

type CloudStorage struct {
	url string
}

func NewCloudStorage(url string) *CloudStorage {
	return &CloudStorage{
		url: url,
	}
}

func (cs *CloudStorage) ReadFileFromFolder(ctx context.Context, filePath string) ([]byte, error) {
	bucket, err := blob.OpenBucket(ctx, "gs://playj-music-storage")

	if err != nil {
		return nil, fmt.Errorf("could not open bucket: %v", err)
	}
	defer bucket.Close()

	file, err := bucket.ReadAll(ctx, filePath)

	if err != nil {
		return nil, fmt.Errorf("error reading the file")
	}
	return file, nil
}
