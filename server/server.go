package main

import (
	"context"
	"log"
	"net"

	pb "github.com/rrd1986/grpc-rest-gateway/gen/proto"
	"google.golang.org/grpc"
)

type TestApiServer struct {
	pb.UnimplementedTestApiServer
}

func (s *TestApiServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func (s *TestApiServer) GetUserGetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("could not open the tcp server: %s", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterTestApiServer(grpcServer, &TestApiServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalln(err)
	}
}
