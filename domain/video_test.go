package domain_test

import (
	"encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestValidateIfVideoIdIsNotUUID(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "dummy-id"
	video.FilePath = "file.path"
	video.ResourceID = "some_resource_id"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.NewV4().String()
	video.FilePath = "file.path"
	video.ResourceID = "some_resource_id"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
