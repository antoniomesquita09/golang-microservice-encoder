package services

import (
	"cloud.google.com/go/storage"
	"context"
	"encoder/application/repositories"
	"encoder/domain"
	"io"
	"log"
	"os"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(bucketName string) error {

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(v.Video.FilePath)

	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()

	body, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	log.Printf("video %s has been storage", v.Video.ID)

	return nil
}
