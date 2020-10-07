package main

import (
	conf "github.com/fjlyx97/short_url/config"
	"github.com/fjlyx97/short_url/log"
	_ "github.com/fjlyx97/short_url/log"
)

func main() {
	// 初始化配置文件
	conf.InitConfig()
	log.InitLogger()
}
