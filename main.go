package main

import (
	"context"
	"fmt"
	pb "github.com/fjlyx97/short_url/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func main() {
	listen, err := net.Listen("tcp", ":10011")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	err = s.Serve(listen)
	if err != nil {
		fmt.Println(err)
	}
}
