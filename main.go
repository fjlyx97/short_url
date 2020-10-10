package main

import (
	"fmt"
	"github.com/fjlyx97/short_url/controllers"
	"github.com/fjlyx97/short_url/dao"
	pb "github.com/fjlyx97/short_url/proto"
	conf "github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	// 捕获panic异常并记录
	defer func() {
		if err := recover(); err != nil {
			log.GLogger.Fatalf("Server crash , err : %s", err)
		}
	}()
	// 初始化配置文件
	conf.InitConfig()
	log.InitLogger()
	// 启动服务器
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.GConfig.BaseModel.Port))
	if err != nil {
		log.GLogger.Fatal(err)
		panic(err)
	}
	grpcServer := grpc.NewServer()
	shortUrlServer := controllers.NewShortUrlServer(
		&dao.MysqlDB{},
	)
	pb.RegisterShortUrlServiceServer(grpcServer, shortUrlServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		log.GLogger.Fatal(err)
		panic(err)
	}
}
