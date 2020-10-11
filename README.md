# Go短网址微服务
Go编写的短网址微服务，用于将长网址映射成指定的短网址：
```
https://github.com/fjlyx97/short_url ==> your_short_url.cn/xxxx
```

# 编译
## 编译程序
1. 确保Go环境安装正确（Go 1.13 , Go Module开启）
2. go build

## proto 文件（可以不用编译）
```shell script
protoc -I . -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.4.1 --go_out=plugins=grpc:. --validate_out="lang=go:./" short_url.proto
```

# 运行
1. 直接启动编译后的程序，确保配置文件配置正常
2. main_test.go用来测试是否成功生成短网址

# 协议
1. grpc协议见proto文件
2. 健康检查路径见配置文件

## 短链生成算法
1. Snowflake雪花算法

# 用到的第三方模块
- gRpc 通信模块
- protoc-gen-validate proto文件校验模块
- yaml.v2 配置文件解析模块
- logrus 日志模块
- file-rotatelogs 日志文件切割
- gorm 数据库模块

# 数据库建库建表
```mysql
CREATE DATABASE short_url DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci; 
CREATE TABLE links (
     uid bigint not null primary key comment '雪花算法生成的唯一id',
     url varchar(8182) not null comment '对应链接的网址',
     insert_at bigint not null comment '插入时间',
     available_at bigint not null comment '生效时间',
     available tinyint(1) not null comment '是否生效'
) ENGINE=InnoDB  DEFAULT CHARSET=utf8;
```

# LICENSE
The MIT License (MIT)