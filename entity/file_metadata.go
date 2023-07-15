package entity

import (
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type FileType int64

const (
	Logs FileType = iota
	Result
)

type FileMetadata struct {
	Agent     string                 `json:"agent,omitempty"`
	AgentId   string                 `json:"agent_id,omitempty"`
	Timestamp *timestamppb.Timestamp `json:"timestamp,omitempty"`
	Resend    int32                  `json:"resend,omitempty"`
	MessageId string                 `json:"message_id,omitempty"`
	Type      FileType               `json:"type,omitempty"`
}

func NewFileMetadataFromGRPC(metadata *agent_service.PackageMetadata) *FileMetadata {
	return &FileMetadata{
		Agent:     metadata.GetAgent(),
		AgentId:   metadata.GetAgentId(),
		Timestamp: metadata.GetTimestamp(),
		Resend:    metadata.GetResend(),
		MessageId: metadata.GetMessageId(),
		Type:      FileType(metadata.Type),
	}
}
