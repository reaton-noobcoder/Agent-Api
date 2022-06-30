package main

import (
	pb "Agent-Ingest/tools/service"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

const (
	port = 8080
)

/*
	Template server for event listener. Once it's better understood which events we want to conusme from the agent we can fill out the 
	.proto contract
*/
type (
	HelloServer struct {
		pb.UnimplementedHelloServer
	}
)

func NewServer() *HelloServer {
	return &HelloServer{}
}

func (s *HelloServer) Hello(stream *pb.HelloRequest) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		fmt.Printf("Unable to listen: %v \n", err)
	}


	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServer(grpcServer, NewServer())

	fmt.Println("Running grpc server")
	grpcServer.Serve(lis)
}