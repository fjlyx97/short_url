package dao

import (
	"fmt"
	"github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisCache struct {
	redisPool *redis.Pool
}

const ShortUrlPrefix = "SHORT_URL_"

func (r *RedisCache) InitCache(ip string, port string, password string) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	pool := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := con.Do("AUTH", password); err != nil {
					_ = con.Close()
					return nil, err
				}
			}
			return con, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	r.redisPool = pool
	return nil
}

func (r *RedisCache) SetUrl(uidB62 string, url string) error {
	key := fmt.Sprintf("%s%s", ShortUrlPrefix, uidB62)
	con := r.GetCon()
	defer con.Close()
	reply, err := con.Do("SETEX", key, config.GConfig.RedisModel.ExpireTime, url)
	if err != nil {
		return err
	}
	log.GLogger.Info(reply)
	return nil
}

func (r *RedisCache) GetUrl(url string) (string, error) {
	key := fmt.Sprintf("%s%s", ShortUrlPrefix, url)
	con := r.GetCon()
	defer con.Close()
	reply, err := redis.String(con.Do("GET", key))
	if err != nil {
		return "", err
	}
	log.GLogger.Info(reply)
	return reply, nil
}

func (r *RedisCache) GetCon() redis.Conn {
	con := r.redisPool.Get()
	return con
}
