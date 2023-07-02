package grpc_server

import (
	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/pb"
)

type Business interface {
	ExecuteBussiness(file *entity.FileInfo) (string, error)
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
