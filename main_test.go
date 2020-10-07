package main

import (
	"context"
	pb "github.com/fjlyx97/short_url/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
)

const (
	addr = "localhost:10011"
)

func Test_server_SayHello(t *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "Yes..",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(r.Message)
}
