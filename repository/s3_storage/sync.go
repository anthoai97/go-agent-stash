package s3_storage

import (
	"time"

	"anquach.dev/go-agent-stash/serializer"
)

func (s *s3Storage) Sync(from string, to string) error {
	startSync := time.Now()
	s.Logger.Infof("Start Sync %s to %s", from, to)
	err := s.Mananger.Sync(from, to)
	if err != nil {
		return err
	}

	syncTime := serializer.XTimeFromNToNow(startSync)
	statistic := s.Mananger.GetStatistics()
	s.Logger.Infof("Sync results: Bytes written: %d Files uploaded: %d Time spent: %d millisecond(s) Files deleted: %d", statistic.Bytes, statistic.Files, syncTime, statistic.DeletedFiles)

	return nil
}
