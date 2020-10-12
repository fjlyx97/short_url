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
	cache         dao.CacheInterface
}

func NewShortUrlServer(db dao.DBInterface, cache dao.CacheInterface) *ShortUrlServer {
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
	if err != nil {
		log.GLogger.Error(err)
		panic(err)
	}
	log.GLogger.Info("Initial database success.")

	// 初始化缓存，可扩展
	if config.GConfig.RedisModel.Start != 0 {
		switch cache.(type) {
		case *dao.RedisCache:
			err = cache.InitCache(
				config.GConfig.RedisModel.Ip,
				config.GConfig.RedisModel.Port,
				config.GConfig.RedisModel.Password,
			)
		}
		if err != nil {
			log.GLogger.Error(err)
			panic(err)
		}
		log.GLogger.Info("Initial cache success.")
	} else {
		cache = nil
	}

	return &ShortUrlServer{
		urlConversion: snowFlake,
		db:            db,
		cache:         cache,
	}
}

func (s *ShortUrlServer) SetShortUrl(ctx context.Context, req *pb.SetUrlReq) (*pb.SetUrlRsp, error) {
	_ = ctx
	rsp := &pb.SetUrlRsp{
		Code:     0,
		ShortUrl: "",
	}
	err := req.Validate()
	if err != nil {
		log.GLogger.Error(err)
		rsp.Code = pb.Code_ERR_PARAMETER
		return rsp, err
	}
	url := req.GetUrl()
	service := services.ShortUrlService{}
	shortUrl, err := service.SetShortUrl(s.db, s.cache, s.urlConversion, url)
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
	rsp := &pb.GetAfterForwardUrlRsp{
		Code:    0,
		LongUrl: "",
	}
	err := req.Validate()
	if err != nil {
		log.GLogger.Error(err)
		rsp.Code = pb.Code_ERR_PARAMETER
		return rsp, err
	}
	url := req.GetUrl()
	service := services.ShortUrlService{}
	longUrl, err := service.GetLongUrl(s.db, s.cache, url)
	if err != nil {
		log.GLogger.Error(err)
		rsp.Code = pb.Code_ERR_SERVICE
		return rsp, err
	}
	rsp.Code = pb.Code_OK
	rsp.LongUrl = longUrl
	return rsp, nil
}
