package main

import (
	"fmt"
	"github.com/fjlyx97/short_url/controllers"
	pb "github.com/fjlyx97/short_url/proto"
	conf "github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
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
	pb.RegisterShortUrlServiceServer(grpcServer, &controllers.ShortUrlServer{})
	err = grpcServer.Serve(listen)
	if err != nil {
		log.GLogger.Fatal(err)
		panic(err)
	}
}
