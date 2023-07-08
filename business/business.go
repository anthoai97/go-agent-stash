package business

import (
	"anquach.dev/go-agent-stash/entity"
)

type DiskStorage interface {
	Save(file *entity.FileInfo, isAppend bool) (string, error)
	SaveAppend(file *entity.FileInfo) (string, error)
	SaveNew(file *entity.FileInfo) (string, error)
}

type business struct {
	diskStorage DiskStorage
}

func NewBusiness(diskStorage DiskStorage) *business {
	return &business{
		diskStorage: diskStorage,
	}
}
