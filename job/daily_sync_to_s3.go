package job

import "anquach.dev/go-agent-stash/serializer"

func (j *JobManager) DailySyncToS3() error {
	if !serializer.GetEnvVar("JOB_SYNC_TO_S3", false) {
		return nil
	}

	j.Logger.Info("Register DailySyncToS3")
	dailySyncToS3Job, err := j.cron.AddFunc("*/20 * * * * *", func() {
		j.Logger.Info("Executing job DailySyncToS3")
		j.bussiness.SyncToS3(serializer.GetEnvVar("STASH_ROOT_PATH", "stash"), serializer.GetEnvVar("S3_ROOT_PATH", ""))
	})

	if err != nil {
		j.cron.Remove(dailySyncToS3Job)
		return err
	}

	return nil
}
