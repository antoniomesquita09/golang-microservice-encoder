package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.VideoRepositoryDb{Db: db}
	repoVideo.Insert(video)

	jobOutput := "output_path"
	jobStatus := "Pending"
	job, err := domain.NewJob(jobOutput, jobStatus, video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	j, err := repoJob.Insert(job)

	//a := fmt.Sprintf("job %s with output path: %s", j.ID, j.OutputBucketPath)
	//b := fmt.Sprintf("video %s with path: %s", j.Video.ID, j.Video.FilePath)
	//fmt.Println(a)
	//fmt.Println(b)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
	require.Equal(t, j.Video.ID, video.ID)
	require.Equal(t, j.OutputBucketPath, jobOutput)
	require.Equal(t, j.Status, jobStatus)
}

func TestVideoRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repoVideo := repositories.VideoRepositoryDb{Db: db}
	repoVideo.Insert(video)

	jobOutput := "output_path"
	jobStatus := "Pending"
	job, err := domain.NewJob(jobOutput, jobStatus, video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	repoJob.Insert(job)

	job.Status = "Complete"
	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Status, job.Status)
}
