package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"anquach.dev/go-agent-stash/business"
	"anquach.dev/go-agent-stash/job"
	agent_service "anquach.dev/go-agent-stash/proto/agent"
	"anquach.dev/go-agent-stash/repository/disk"
	"anquach.dev/go-agent-stash/repository/s3_storage"
	"anquach.dev/go-agent-stash/serializer"
	"anquach.dev/go-agent-stash/transport/grpc_server"
	gateway "anquach.dev/go-agent-stash/transport/http_server"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Env string

const (
	Develop    Env = "develop"
	Production Env = "production"
)

func main() {
	fmt.Println("Init Agent Stash....")

	env := serializer.GetEnvVar("ENV", "develop")
	if env == string(Develop) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// repo
	diskStorage := disk.NewDiskStorage(serializer.GetEnvVar("STASH_ROOT_PATH", "stash"))
	s3Storage := s3_storage.News3Storage()

	// bussiness
	biz := business.NewBusiness(diskStorage, s3Storage)

	// Schedule Jobs
	jobManager := job.NewJobManager(biz)
	if err := jobManager.StartJobs(); err != nil {
		log.Fatal(err.Error())
	}

	startGRPCServerAndGateway(biz)
}

func startGRPCServerAndGateway(biz *business.Business) {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	addr := "0.0.0.0:9090"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// transport gRPC
	implServer := grpc_server.NewGrpcServer(biz)
	s := grpc.NewServer()

	// Register gRPC Service
	agent_service.RegisterAgentServiceServer(s, implServer)
	grpc_health_v1.RegisterHealthServer(s, implServer)

	// Register gRPC refection
	reflection.Register(s)

	// Start gRPC Server
	log.Info("Start gRPC Service at ", lis.Addr().String())
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// Start Gateway http
	err = gateway.Run("dns:///" + addr)
	log.Fatalln(err)
}
