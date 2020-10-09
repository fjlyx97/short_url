package controllers

import (
	"context"
	"fmt"
	pb "github.com/fjlyx97/short_url/proto"
	"github.com/fjlyx97/short_url/services/snowflake"
	"github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
	"strconv"
)

type ShortUrlServer struct {
	snowFlake *snowflake.Snowflake
}

func NewShortUrlServer() *ShortUrlServer {
	snowFlake, err := snowflake.NewSnowflake(
		config.GConfig.SnowFlakeModel.StartTimeStamp,
		config.GConfig.SnowFlakeModel.DatacenterId,
		config.GConfig.SnowFlakeModel.WorkerId,
	)
	if err != nil {
		panic(err)
	}
	return &ShortUrlServer{
		snowFlake: snowFlake,
	}
}

func (s *ShortUrlServer) SetShortUrl(ctx context.Context, req *pb.SetUrlReq) (*pb.SetUrlRsp, error) {
	// Mock 数据
	fmt.Println(req.GetUrl())
	id, err := s.snowFlake.NextId()
	if err != nil {
		log.GLogger.Error(err)
	}
	setUrlRsp := &pb.SetUrlRsp{
		Code:     pb.Code_OK,
		ShortUrl: strconv.FormatInt(id, 10),
	}
	return setUrlRsp, nil
}

func (s *ShortUrlServer) GetAfterForwardUrl(ctx context.Context, req *pb.GetAfterForwardUrlReq) (*pb.GetAfterForwardUrlRsp, error) {

	return nil, nil
}
