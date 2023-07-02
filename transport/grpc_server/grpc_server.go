package grpc_server

import (
	agent_service "anquach.dev/go-agent-stash/pb"
)

type Business interface {
	ExecuteBussiness(*agent_service.SimplePackage) (string, error)
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
