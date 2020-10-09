package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	GConfig    *Config
	configPath = "./config.yaml"
)

func init() {
	flag.StringVar(&configPath, "conf", "./config.yaml", "Set the yaml config path")
}

// 解析基础配置文件
func InitConfig() {
	GConfig = &Config{}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		//配置文件解析失败就没有继续往下的必要了
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, GConfig)
	if err != nil {
		panic(err)
	}
}
