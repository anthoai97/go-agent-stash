package entity

import (
	"fmt"

	agent_service "anquach.dev/go-agent-stash/pb"
)

type FileType int64

type FileInfo struct {
	FileName string
	FilePath string
	Data     []byte
	Metadata *agent_service.PackageMetadata
}

func NewFileFromSimplePackage(in *agent_service.SimplePackage) *FileInfo {
	fileName := GenerateFileNameFromMetadata(in.GetMetadata().GetType())
	filePath := GenerateFilePathFromMetadata(in.GetMetadata(), fileName)
	fileData := []byte(in.GetData())

	return &FileInfo{
		FileName: fileName,
		FilePath: filePath,
		Metadata: in.GetMetadata(),
		Data:     fileData,
	}
}

func GenerateFilePathFromMetadata(metadata *agent_service.PackageMetadata, fileName string) string {
	filePath := fmt.Sprintf("%s/%s/%s", metadata.GetAgentId(), metadata.GetMessageId(), fileName)
	return filePath
}

func GenerateFileNameFromMetadata(mgsType agent_service.MessageType) string {
	switch mgsType {
	case agent_service.MessageType_LOG:
		return "logs.txt"
	case agent_service.MessageType_RESULT:
		return "result.json"
	}
	return "unknown.txt"
}
