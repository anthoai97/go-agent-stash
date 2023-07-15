package business

import (
	"fmt"
	"sync"

	"anquach.dev/go-agent-stash/entity"
)

func (biz *Business) ExecuteMsgPack(files []*entity.FileInfo) ([]*entity.FileExecuteStatus, error) {
	if len(files) < 1 {
		return nil, fmt.Errorf("files is empty")
	}

	resp := make([]*entity.FileExecuteStatus, 0)
	resChan := make(chan *entity.FileExecuteStatus, 1000)
	var wg = new(sync.WaitGroup)

	for _, file := range files {
		wg.Add(1)

		go func(file *entity.FileInfo) {
			defer wg.Done()
			var path string
			var err error

			if file.Metadata.Type == 0 {
				path, err = biz.diskStorage.SaveAppend(file)
			} else {
				path, err = biz.diskStorage.SaveNew(file)
			}

			status := file.GenerateFileExecuteStatus()

			if err != nil {
				status.Success = false
				status.Error = err
				resChan <- status
				return
			}

			status.Path = path
			resChan <- status
		}(file)
	}

	for i := 0; i < len(files); i++ {
		stt := <-resChan
		resp = append(resp, stt)
	}

	wg.Wait()
	close(resChan)

	return resp, nil
}
