package disk

import (
	"sync"
)

type DiskStorage struct {
	mutex    sync.RWMutex
	RootPath string
}

func NewDiskStorage(rootPath string) *DiskStorage {
	return &DiskStorage{
		RootPath: rootPath,
	}
}
