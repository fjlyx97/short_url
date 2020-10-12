package services

import (
	"github.com/fjlyx97/short_url/dao"
	"github.com/fjlyx97/short_url/utils"
	"github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
	"time"
)

type ShortUrlService struct {
}

func (s *ShortUrlService) SetShortUrl(db dao.DBInterface, cache dao.CacheInterface, urlConversion UrlConversion, longUrl string) (string, error) {
	uid, err := urlConversion.NextId()
	uidB62 := utils.Base10ToBase62(uid)
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
	// 写入缓存
	if cache != nil {
		err := cache.SetUrl(uidB62, longUrl)
		if err != nil {
			return "", nil
		}
	}

	shortUrl := config.GConfig.BaseModel.BaseUrl + uidB62
	return shortUrl, nil
}

func (s *ShortUrlService) GetLongUrl(db dao.DBInterface, cache dao.CacheInterface, shortUrl string) (string, error) {
	uid := utils.Base62ToBase10(shortUrl)
	// 先查缓存
	if cache != nil {
		longUrl, err := cache.GetUrl(shortUrl)
		if err != nil {
			log.GLogger.Info(err)
		} else {
			return longUrl, err
		}
	}

	// 查询数据库
	longUrl, err := db.FindLink(uid)
	if err != nil {
		return "", err
	}
	return longUrl, nil
}
