package grpc_server

import (
	"context"
	"log"

	dsr "anquach.dev/go-agent-stash/proto/dsr_agent"
)

func (s *grpcServer) SendMessage(ctx context.Context, in *dsr.GRPCMessagePackage) (*dsr.ServerReply, error) {
	log.Printf("Received: %v", in.GetData())
	res := s.business.ExecuteBussiness()
	return &dsr.ServerReply{Message: "Hello " + in.GetAgent() + " Get from biz: " + res}, nil
}
