package s3_storage

import (
	"anquach.dev/go-agent-stash/serializer"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seqsense/s3sync"
	"google.golang.org/grpc/grpclog"
)

type s3Storage struct {
	Logger grpclog.LoggerV2
}

func News3Storage() *s3Storage {
	// Load `config`
	logger := serializer.CustomLogger()

	return &s3Storage{
		Logger: logger,
	}
}

func (s *s3Storage) newSession() *s3sync.Manager {
	region := serializer.GetEnvVar("AWS_DEFAULT_REGION", "")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		panic(err)
	}

	s3sync.SetLogger(NewCustomS3SyncLogger(s.Logger))
	return s3sync.New(sess)
}
