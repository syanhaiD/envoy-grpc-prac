package main

import (
	"fmt"
	egrpcpb "github.com/syanhaiD/envoy-grpc-prac/pkg/pb"
	egrpcserver "github.com/syanhaiD/envoy-grpc-prac/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	os.Exit(exec())
}

func exec() int {
	listenPort, err := net.Listen("tcp", ":26000")
	if err != nil {
		fmt.Println(err)
		return 1
	}
	server := grpc.NewServer()
	egrpcService := &egrpcserver.EGRPCService{}
	egrpcpb.RegisterEGRPCServer(server, egrpcService)
	reflection.Register(server)
	grpc_health_v1.RegisterHealthServer(server, &egrpcserver.HealthService{})
	err = server.Serve(listenPort)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
