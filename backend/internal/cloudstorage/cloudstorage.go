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

func (cs *CloudStorage) ReadFileFromFolder(ctx context.Context, folderName string, fileName string) ([]byte, error) {
	bucketURL := fmt.Sprintf("%s?prefix=%s/", "gs://playj-music-storage", folderName)
	bucket, err := blob.OpenBucket(ctx, bucketURL)

	if err != nil {
		return nil, fmt.Errorf("could not open bucket: %v", err)
	}
	defer bucket.Close()
	file, err := bucket.ReadAll(ctx, "One.mp3")

	if err != nil {
		return nil, fmt.Errorf("error reading the file")
	}
	return file, nil
}

// func (cs *CloudStorage) getBucket(ctx context.Context, )
