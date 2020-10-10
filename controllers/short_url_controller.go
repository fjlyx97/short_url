package controllers

import (
	"context"
	"github.com/fjlyx97/short_url/dao"
	pb "github.com/fjlyx97/short_url/proto"
	"github.com/fjlyx97/short_url/services"
	"github.com/fjlyx97/short_url/services/snowflake"
	"github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
)

type ShortUrlServer struct {
	urlConversion services.UrlConversion
	db            dao.DBInterface
}

func NewShortUrlServer(db dao.DBInterface) *ShortUrlServer {
	// 初始化短网址算法，可扩展
	snowFlake, err := snowflake.NewSnowflake(
		config.GConfig.SnowFlakeModel.StartTimeStamp,
		config.GConfig.SnowFlakeModel.DatacenterId,
		config.GConfig.SnowFlakeModel.WorkerId,
	)
	if err != nil {
		panic(err)
	}
	// 初始化DB，可扩展
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
		urlConversion: snowFlake,
		db:            db,
	}
}

func (s *ShortUrlServer) SetShortUrl(ctx context.Context, req *pb.SetUrlReq) (*pb.SetUrlRsp, error) {
	_ = ctx
	//TODO 校验参数

	url := req.GetUrl()
	service := services.ShortUrlService{}
	shortUrl, err := service.SetShortUrl(s.db, s.urlConversion, url)
	rsp := &pb.SetUrlRsp{
		Code:     0,
		ShortUrl: "",
	}
	if err != nil {
		log.GLogger.Error(err)
		rsp.Code = pb.Code_ERR_SERVICE
		return rsp, err
	}
	rsp.Code = pb.Code_OK
	rsp.ShortUrl = shortUrl
	return rsp, nil
}

func (s *ShortUrlServer) GetAfterForwardUrl(ctx context.Context, req *pb.GetAfterForwardUrlReq) (*pb.GetAfterForwardUrlRsp, error) {
	_ = ctx
	//TODO 校验参数

	url := req.GetUrl()
	service := services.ShortUrlService{}
	longUrl, err := service.GetLongUrl(s.db, url)
	rsp := &pb.GetAfterForwardUrlRsp{
		Code:    0,
		LongUrl: "",
	}
	if err != nil {
		log.GLogger.Error(err)
		rsp.Code = pb.Code_ERR_SERVICE
		return rsp, err
	}
	rsp.Code = pb.Code_OK
	rsp.LongUrl = longUrl
	return rsp, nil
}
