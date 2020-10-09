package controllers

import (
	"context"
	"fmt"
	pb "github.com/fjlyx97/short_url/proto"
)

type ShortUrlServer struct {
}

func (s *ShortUrlServer) SetShortUrl(ctx context.Context, req *pb.SetUrlReq) (*pb.SetUrlRsp, error) {
	// Mock 数据
	fmt.Println(req.GetUrl())
	setUrlRsp := &pb.SetUrlRsp{
		Code:     pb.Code_OK,
		ShortUrl: "111",
	}
	return setUrlRsp, nil
}

func (s *ShortUrlServer) GetAfterForwardUrl(ctx context.Context, req *pb.GetAfterForwardUrlReq) (*pb.GetAfterForwardUrlRsp, error) {

	return nil, nil
}
