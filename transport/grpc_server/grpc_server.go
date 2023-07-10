package grpc_server

import (
	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Business interface {
	ExecuteMsgPack(files []*entity.FileInfo) ([]*entity.FileExecuteStatus, error)
}

type grpcServer struct {
	business Business
	agent_service.UnimplementedAgentServiceServer
	// reader_service.UnimplementedGreeterServiceServer
	grpc_health_v1.UnimplementedHealthServer
}

func NewGrpcServer(biz Business) *grpcServer {
	return &grpcServer{
		business: biz,
	}
}
