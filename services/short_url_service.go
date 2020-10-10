package services

import (
	"github.com/fjlyx97/short_url/dao"
	"github.com/fjlyx97/short_url/utils"
	"github.com/fjlyx97/short_url/utils/config"
	"time"
)

type ShortUrlService struct {
}

func (s *ShortUrlService) SetShortUrl(db dao.DBInterface, urlConversion UrlConversion, longUrl string) (string, error) {
	uid, err := urlConversion.NextId()
	if err != nil {
		return "", err
	}
	// 写入数据库
	oneDayAfter, err := time.ParseDuration("24h")
	if err != nil {
		return "", err
	}
	availableTime := time.Now().Add(oneDayAfter).UnixNano()
	err = db.CreateLink(uid, longUrl, availableTime)
	if err != nil {
		return "", err
	}
	shortUrl := config.GConfig.BaseModel.BaseUrl + utils.Base10ToBase62(uid)
	return shortUrl, nil
}

func (s *ShortUrlService) GetLongUrl(db dao.DBInterface, shortUrl string) (string, error) {
	uid := utils.Base62ToBase10(shortUrl)
	// TODO 先查redis

	// 查询数据库
	longUrl, err := db.FindLink(uid)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}
