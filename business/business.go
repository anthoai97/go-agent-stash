package business

import (
	"anquach.dev/go-agent-stash/entity"
)

type DiskStorage interface {
	Save(file *entity.FileInfo, isAppend bool) (string, error)
	SaveAppend(file *entity.FileInfo) (string, error)
	SaveNew(file *entity.FileInfo) (string, error)
}

type S3Storage interface {
	Sync(from string, to string) error
}

type Business struct {
	diskStorage DiskStorage
	s3Storage   S3Storage
}

func NewBusiness(diskStorage DiskStorage, s3Storage S3Storage) *Business {
	return &Business{
		diskStorage: diskStorage,
		s3Storage:   s3Storage,
	}
}
