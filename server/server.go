package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	// Start the Rest server
	go func() {
		// mux
		mux := runtime.NewServeMux()

		// register handle
		pb.RegisterTestApiHandlerServer(context.Background(), mux, &TestApiServer{})

		// launch a http server
		log.Fatalln(http.ListenAndServe("localhost:8081", mux))

	}()

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
