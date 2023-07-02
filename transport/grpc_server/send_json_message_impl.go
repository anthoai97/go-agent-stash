package grpc_server

import (
	"context"

	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcServer) SendJsonMsgPack(ctx context.Context, in *agent_service.JsonMsgPack) (*agent_service.ServerReply, error) {
	files := []*entity.FileInfo{entity.NewFileFromJsonPackage(in)}
	res, err := s.business.ExecuteMsgPack(files)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	replyData := make([]*agent_service.PackageExcuteStatus, 0)
	for _, stt := range res {
		sttToGrpc := &agent_service.PackageExcuteStatus{
			AgentId:   stt.AgentID,
			Success:   stt.Success,
			MessageId: stt.MessageID,
			Path:      stt.Path,
			Error:     stt.Error,
		}
		replyData = append(replyData, sttToGrpc)
	}

	return &agent_service.ServerReply{
		Reply: replyData,
	}, nil
}
