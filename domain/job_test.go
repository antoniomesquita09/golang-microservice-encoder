package domain_test

import (
	"encoder/domain"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.CreatedAt = time.Now()
	video.ResourceID = "some-resource-id"
	video.FilePath = "some-path"

	job, err := domain.NewJob("path", "CONVERTED", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
