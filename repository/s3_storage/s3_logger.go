package s3_storage

import (
	"google.golang.org/grpc/grpclog"
)

type customS3SyncLogger struct {
	Logger grpclog.LoggerV2
}

func NewCustomS3SyncLogger(logger grpclog.LoggerV2) *customS3SyncLogger {
	return &customS3SyncLogger{
		Logger: logger,
	}
}

func (l *customS3SyncLogger) Log(v ...interface{}) {
	l.Logger.Infoln(v...)
}

func (l *customS3SyncLogger) Logf(format string, v ...interface{}) {
	l.Logger.Infof(format, v...)
}
