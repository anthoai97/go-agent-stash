package grpc_server

import (
	dsr "anquach.dev/go-agent-stash/proto/dsr_agent"
)

type Business interface {
	ExecuteBussiness() string
}

type grpcServer struct {
	business Business
	dsr.UnimplementedDsrAgentServer
}

func NewGrpcServer(biz Business) *grpcServer {
	return &grpcServer{
		business: biz,
	}
}
