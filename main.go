package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"anquach.dev/go-agent-stash/business"
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	reader_service "anquach.dev/go-agent-stash/proto/reader"
	"anquach.dev/go-agent-stash/repository/disk"
	"anquach.dev/go-agent-stash/serializer"
	"anquach.dev/go-agent-stash/transport/grpc_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Init Agent Stash....")

	fmt.Println("1. Start grpc server")
	startGRPCServer()
}

func startGRPCServer() {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:9090"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// repo
	diskStorage := disk.NewDiskStorage(serializer.GetEnvVar("STASH_ROOT_PATH", "stash"))

	// bussiness
	biz := business.NewBusiness(diskStorage)

	// transport gRPC
	implServer := grpc_server.NewGrpcServer(biz)
	s := grpc.NewServer()

	// Register gRPC Service
	agent_service.RegisterAgentServiceServer(s, implServer)
	reader_service.RegisterGreeterServiceServer(s, implServer)

	// Register gRPC refection
	reflection.Register(s)

	// Start gRPC Server
	fmt.Printf("Start gRPC Service at %v\n", lis.Addr())
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// Start Gateway http
	// TODO: Implement Gateway
}
