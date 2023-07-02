package grpc_server

import (
	agent_service "anquach.dev/go-agent-stash/pb"
)

type Business interface {
	ExecuteBussiness() string
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
