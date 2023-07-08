package disk

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	"anquach.dev/go-agent-stash/entity"
	"anquach.dev/go-agent-stash/serializer"
)

// Save new current file
func (store *DiskStorage) Save(fileInfo *entity.FileInfo, isAppend bool) (string, error) {
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

	prefix := 0
	files, err := serializer.FilePathWalkDir(dir)
	if err == nil {
		prefix = len(files)
	}

	if isAppend && prefix != 0 {
		prefix -= 1
	}

	finalPath := strings.Replace(fullPath, "suffix", strconv.Itoa(prefix), 1)
	fi, err := os.Stat(finalPath)
	if err == nil && isAppend {
		if fi.Size() >= serializer.GetEnvVar[int64]("MAX_FILE_SIZE", 5000) { // Default 5KB
			finalPath = strings.Replace(fullPath, "suffix", strconv.Itoa(prefix+1), 1)
		}
	}

	osFile, err := os.OpenFile(finalPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	return finalPath, nil
}

func (store *DiskStorage) SaveAppend(fileInfo *entity.FileInfo) (string, error) {
	return store.Save(fileInfo, true)
}

func (store *DiskStorage) SaveNew(fileInfo *entity.FileInfo) (string, error) {
	return store.Save(fileInfo, false)
}
