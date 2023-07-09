package grpc_server

import (
	"context"
	"fmt"

	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcServer) SendJsonMsgPack(ctx context.Context, in *agent_service.JsonMsgPack) (*agent_service.ServerReply, error) {
	fmt.Printf("Receive SendJsonMsgPack from %s | %s\n", in.Metadata.GetAgent(), in.Metadata.GetMessageId())
	files := entity.NewFileFromJsonPackage(in)
	res, err := s.business.ExecuteMsgPack(files)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	replyData := make([]*agent_service.PackageExcuteStatus, 0)
	for _, stt := range res {
		err := ""
		if stt.Error != nil {
			err = fmt.Sprintf("SendSimpleMsgPack throw error ==> %s", stt.Error.Error())
		}
		sttToGrpc := &agent_service.PackageExcuteStatus{
			AgentId:   stt.AgentID,
			Success:   stt.Success,
			MessageId: stt.MessageID,
			Path:      stt.Path,
			Error:     err,
		}
		replyData = append(replyData, sttToGrpc)
	}

	return &agent_service.ServerReply{
		Reply: replyData,
	}, nil
}
