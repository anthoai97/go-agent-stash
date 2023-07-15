package grpc_server

import (
	"anquach.dev/go-agent-stash/business"
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type grpcServer struct {
	business *business.Business
	agent_service.UnimplementedAgentServiceServer
	// reader_service.UnimplementedGreeterServiceServer
	grpc_health_v1.UnimplementedHealthServer
}

func NewGrpcServer(biz *business.Business) *grpcServer {
	return &grpcServer{
		business: biz,
	}
}
