package log

import (
	"github.com/fjlyx97/short_url/utils/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

var GLogger = logrus.New()

func InitLogger() {
	writer, err := rotatelogs.New(
		config.GConfig.LogModel.LogPath+".%Y%m%d%H%M",
		//rotatelogs.WithLinkName(config.GConfig.LogModel.LogPath),
		rotatelogs.WithRotationTime(time.Duration(config.GConfig.RotationTime)*time.Second),
	)
	if err != nil {
		panic(err)
	}
	GLogger.SetOutput(writer)
	GLogger.SetFormatter(&logrus.JSONFormatter{})
	GLogger.Info("Initial logger success.")
	GLogger.Info(config.GConfig)
}
