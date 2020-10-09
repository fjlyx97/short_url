package controllers

import (
	"context"
	"fmt"
	"github.com/fjlyx97/short_url/dao"
	pb "github.com/fjlyx97/short_url/proto"
	"github.com/fjlyx97/short_url/services/snowflake"
	"github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
	"strconv"
)

type ShortUrlServer struct {
	snowFlake *snowflake.Snowflake
	db        dao.DBInterface
}

func NewShortUrlServer(db dao.DBInterface) *ShortUrlServer {
	snowFlake, err := snowflake.NewSnowflake(
		config.GConfig.SnowFlakeModel.StartTimeStamp,
		config.GConfig.SnowFlakeModel.DatacenterId,
		config.GConfig.SnowFlakeModel.WorkerId,
	)
	if err != nil {
		panic(err)
	}
	// 初始化DB
	switch db.(type) {
	case *dao.MysqlDB:
		err = db.InitDB(
			config.GConfig.MysqlModel.Ip,
			config.GConfig.MysqlModel.Port,
			config.GConfig.MysqlModel.User,
			config.GConfig.MysqlModel.Password,
			config.GConfig.MysqlModel.DatabaseName,
			config.GConfig.MysqlModel.ConnectTimeout,
			config.GConfig.MysqlModel.ReadTimeout,
			config.GConfig.MysqlModel.WriteTimeout,
		)
	}
	log.GLogger.Info("Initial database success.")

	if err != nil {
		log.GLogger.Error(err)
		panic(err)
	}

	return &ShortUrlServer{
		snowFlake: snowFlake,
		db:        db,
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
