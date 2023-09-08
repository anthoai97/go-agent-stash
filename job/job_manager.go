package job

import (
	"anquach.dev/go-agent-stash/business"
	"anquach.dev/go-agent-stash/serializer"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc/grpclog"
)

type JobManager struct {
	Logger    grpclog.LoggerV2
	bussiness *business.Business
	cron      *cron.Cron
}

func NewJobManager(business *business.Business) *JobManager {
	c := cron.New()

	return &JobManager{
		Logger:    serializer.CustomLogger(),
		bussiness: business,
		cron:      c,
	}
}

func (j *JobManager) StartJobs() error {
	err := j.DailySyncToS3()
	if err != nil {
		j.Logger.Info("CRON JOB STARTED failed " + err.Error())
		return err
	}

	go func() {
		j.Logger.Info("CRON JOB STARTED ALL SUCCESSFUL")
		j.cron.Run()
	}()

	return nil
}
