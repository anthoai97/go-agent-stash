package disk

import (
	"fmt"
	"os"
	"path"

	"anquach.dev/go-agent-stash/entity"
)

// Save override current file
func (store *DiskStorage) Save(fileInfo *entity.FileInfo) (string, error) {
	fullPath := fmt.Sprintf("%s/%s", store.RootPath, fileInfo.FilePath)
	dir := path.Dir(fullPath)
	ext := path.Ext(fullPath)
	if ext != ".json" && ext != ".txt" {
		return "", fmt.Errorf("not support file extention: %s", ext)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0775)
		if err != nil {
			return "", err
		}
	}

	osFile, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", fmt.Errorf("cannot create file: %w", err)
	}
	defer osFile.Close()
	writeData := append(fileInfo.Data, '\n')
	_, err = osFile.Write(writeData)
	if err != nil {
		return "", fmt.Errorf("cannot write to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	return fullPath, nil
}
