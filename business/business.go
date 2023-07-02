package business

import (
	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/pb"
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

func (biz *business) ExecuteBussiness(in *agent_service.SimplePackage) (string, error) {
	file := entity.NewFileFromSimplePackage(in)
	result, err := biz.diskStorage.Save(file)
	if err != nil {
		return "", err
	}

	return result, nil
}
