package entity

import (
	"fmt"
	"strconv"

	agent_service "anquach.dev/go-agent-stash/pb"
	"anquach.dev/go-agent-stash/serializer"
)

type FileType int64

type FileInfo struct {
	FileName string
	FilePath string
	Data     []byte
	Metadata *FileMetadata
}

func NewFileFromSimplePackage(in *agent_service.SimplePackage) *FileInfo {
	fileName := GenerateFileNameFromMetadata(in.GetMetadata().GetType(), strconv.Itoa(0))
	filePath := GenerateFilePathFromMetadata(in.GetMetadata(), fileName)
	data := serializer.WriteArrayStringToByte(in.GetData())

	return &FileInfo{
		FileName: fileName,
		FilePath: filePath,
		Metadata: NewFileMetadataFromGRPC(in.Metadata),
		Data:     data,
	}
}

func NewFileFromJsonPackage(in *agent_service.JsonMsgPack) []*FileInfo {
	files := make([]*FileInfo, 0)
	for idx, data := range in.GetData() {
		fileName := GenerateFileNameFromMetadata(in.GetMetadata().GetType(), strconv.Itoa(idx))
		filePath := GenerateFilePathFromMetadata(in.GetMetadata(), fileName)
		dataInByte, err := data.MarshalJSON()
		if err != nil {
			fmt.Println(err)
		}
		files = append(files, &FileInfo{
			FileName: fileName,
			FilePath: filePath,
			Metadata: NewFileMetadataFromGRPC(in.Metadata),
			Data:     dataInByte,
		})
	}

	return files
}

func GenerateFilePathFromMetadata(metadata *agent_service.PackageMetadata, fileName string) string {
	filePath := fmt.Sprintf("%s/%s/%s", metadata.GetAgentId(), metadata.GetMessageId(), fileName)
	return filePath
}

func GenerateFileNameFromMetadata(mgsType agent_service.MessageType, suffix string) string {
	switch mgsType {
	case agent_service.MessageType_LOG:
		return fmt.Sprintf("logs_%s.txt", suffix)
	case agent_service.MessageType_RESULT:
		return fmt.Sprintf("result_%s.json", suffix)
	}
	return "unknown.txt"
}

func (f *FileInfo) GenerateFileExecuteStatus() *FileExecuteStatus {
	return &FileExecuteStatus{
		AgentID:   f.Metadata.AgentId,
		MessageID: f.Metadata.MessageId,
		Success:   true,
	}
}
