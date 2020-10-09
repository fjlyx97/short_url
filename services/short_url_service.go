package services

import (
	"github.com/fjlyx97/short_url/dao"
	"github.com/fjlyx97/short_url/services/snowflake"
	"github.com/fjlyx97/short_url/utils/config"
	"time"
)

type ShortUrlService struct {
}

func (s *ShortUrlService) SetShortUrl(db dao.DBInterface, snowflake *snowflake.Snowflake, longUrl string) (string, error) {
	uid, err := snowflake.NextId()
	if err != nil {
		return "", err
	}
	// 写入数据库
	oneDayAfter, err := time.ParseDuration("24h")
	availableTime := time.Now().Add(oneDayAfter).UnixNano()
	if err != nil {
		return "", err
	}
	err = db.CreateLink(uid, longUrl, availableTime)
	if err != nil {
		return "", err
	}
	shortUrl := config.GConfig.BaseModel.BaseUrl + "123"
	return shortUrl, nil
}
