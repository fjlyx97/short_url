# 用到的第三方模块
- gRpc 通信模块
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