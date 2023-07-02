package grpc_server

import (
	"context"

	agent_service "anquach.dev/go-agent-stash/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcServer) SendSimpleMsgPack(ctx context.Context, in *agent_service.SimplePackage) (*agent_service.ServerReply, error) {
	res, err := s.business.ExecuteBussiness(in)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &agent_service.ServerReply{Message: res}, nil
}
