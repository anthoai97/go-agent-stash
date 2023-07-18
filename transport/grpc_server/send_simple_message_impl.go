package grpc_server

import (
	"context"
	"fmt"

	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcServer) SendSimpleMsgPack(ctx context.Context, in *agent_service.SimplePackage) (*agent_service.ServerReply, error) {
	if in.GetMetadata() == nil || in.GetData() == nil {
		return nil, status.Error(400, "Default error message for 400")
	}
	fmt.Printf("Receive SendJsonMsgPack from %s | %s\n", in.Metadata.GetAgent(), in.Metadata.GetMessageId())

	files := []*entity.FileInfo{entity.NewFileFromSimplePackage(in)}
	res, err := s.business.ExecuteMsgPack(files)
	if err != nil {
		fmt.Println("err ==> " + err.Error())
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	replyData := make([]*agent_service.PackageExcuteStatus, 0)
	for _, stt := range res {
		err := ""
		if stt.Error != nil {
			err = fmt.Sprintf("SendSimpleMsgPack throw error ==> %s", stt.Error.Error())
		}
		sttToGrpc := &agent_service.PackageExcuteStatus{
			Success:   stt.Success,
			AgentId:   stt.AgentID,
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
