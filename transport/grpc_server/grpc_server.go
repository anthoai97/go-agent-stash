package grpc_server

import (
	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/pb"
)

type Business interface {
	ExecuteMsgPack(files []*entity.FileInfo) ([]*entity.FileExecuteStatus, error)
}

type grpcServer struct {
	business Business
	agent_service.UnimplementedAgentServiceServer
}

func NewGrpcServer(biz Business) *grpcServer {
	return &grpcServer{
		business: biz,
	}
}
