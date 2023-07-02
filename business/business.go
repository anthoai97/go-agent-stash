package business

import (
	"anquach.dev/go-agent-stash/entity"
)

type DiskStorage interface {
	Save(file *entity.FileInfo) (string, error)
}

type business struct {
	diskStorage DiskStorage
}

func NewBusiness(diskStorage DiskStorage) *business {
	return &business{
		diskStorage: diskStorage,
	}
}

func (biz *business) ExecuteBussiness(file *entity.FileInfo) (string, error) {
	result, err := biz.diskStorage.Save(file)
	if err != nil {
		return "", err
	}

	return result, nil
}
