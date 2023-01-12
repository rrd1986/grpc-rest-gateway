package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/rrd1986/grpc-rest-gateway/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	client := pb.NewTestApiClient(conn)
	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello Rashmi"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Msg)

}
