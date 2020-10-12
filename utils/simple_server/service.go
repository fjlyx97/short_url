package main

import (
	"context"
	"errors"
	pb "github.com/fjlyx97/short_url/proto"
	"google.golang.org/grpc"
)

var (
	GrpcAddress = "127.0.0.1:8002"
)

func (s *SimpleServer) SetUrlService(ctx context.Context, url string) (string, error) {
	con, err := grpc.Dial(GrpcAddress, grpc.WithInsecure())
	defer con.Close()
	if err != nil {
		return "", err
	}
	// grpc请求
	client := pb.NewShortUrlServiceClient(con)
	reply, err := client.SetShortUrl(ctx, &pb.SetUrlReq{
		Url: url,
	})
	if err != nil {
		return "", err
	}
	// 判断返回码
	if reply.Code != pb.Code_OK {
		return "", errors.New(reply.Code.String())
	}
	return reply.ShortUrl, nil
}

func (s *SimpleServer) GetUrlService(ctx context.Context, url string) (string, error) {
	con, err := grpc.Dial(GrpcAddress, grpc.WithInsecure())
	defer con.Close()
	if err != nil {
		return "", err
	}
	// grpc请求
	client := pb.NewShortUrlServiceClient(con)
	reply, err := client.GetAfterForwardUrl(ctx, &pb.GetAfterForwardUrlReq{
		Url: url,
	})
	if err != nil {
		return "", err
	}
	// 判断返回码
	if reply.Code != pb.Code_OK {
		return "", errors.New(reply.Code.String())
	}
	return reply.LongUrl, nil
}
