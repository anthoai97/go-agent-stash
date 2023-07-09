package grpc_server

import (
	"anquach.dev/go-agent-stash/entity"
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	reader_service "anquach.dev/go-agent-stash/proto/reader"
)

type Business interface {
	ExecuteMsgPack(files []*entity.FileInfo) ([]*entity.FileExecuteStatus, error)
}

type grpcServer struct {
	business Business
	agent_service.UnimplementedAgentServiceServer
	reader_service.UnimplementedGreeterServiceServer
}

func NewGrpcServer(biz Business) *grpcServer {
	return &grpcServer{
		business: biz,
	}
}
