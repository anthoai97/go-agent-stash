package main

import (
	"fmt"
	"log"
	"net"

	"anquach.dev/go-agent-stash/business"
	agent_service "anquach.dev/go-agent-stash/pb"
	"anquach.dev/go-agent-stash/transport/grpc_server"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Init Agent Stash....")

	fmt.Println("1. Start grpc server")
	startGRPCServer()
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// bussiness
	biz := business.NewBusiness()

	implServer := grpc_server.NewGrpcServer(biz)
	s := grpc.NewServer()

	agent_service.RegisterAgentServiceServer(s, implServer)
	fmt.Printf("Start gRPC Service at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Println("err")
		log.Fatalf("failed to serve: %v", err)
	}
}
