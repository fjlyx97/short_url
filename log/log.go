package log

import (
	"fmt"
	"github.com/fjlyx97/short_url/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

func InitLogger() {
	fmt.Println(config.GConfig)
	writer, err := rotatelogs.New(
		config.GConfig.LogModel.LogPath+".%Y%m%d%H%M",
		//rotatelogs.WithLinkName(config.GConfig.LogModel.LogPath),
		rotatelogs.WithRotationTime(time.Duration(config.GConfig.RotationTime)*time.Second),
	)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(writer)
	logrus.Info("Initial logger success.")
}
