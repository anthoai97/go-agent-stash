package grpc_server

import (
	"context"
	"log"

	agent_service "anquach.dev/go-agent-stash/proto"
)

func (s *grpcServer) SendSimpleMsgPack(ctx context.Context, in *agent_service.SimplePackage) (*agent_service.ServerReply, error) {
	log.Printf("Received: %v", in.GetData())
	log.Printf("Received: %v", in.GetMetadata().GetAgentId())
	res := s.business.ExecuteBussiness()

	return &agent_service.ServerReply{Message: "Hello " + " Get from biz: " + res}, nil
}
