package services

import (
	"fmt"
	"github.com/fjlyx97/short_url/utils/config"
	"github.com/fjlyx97/short_url/utils/log"
	"net/http"
)

func RunHealthCheck() {
	go func() {
		http.HandleFunc(config.GConfig.BaseModel.HealthCheckUrl, func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(200)
		})
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.GConfig.BaseModel.HealthCheckPort), nil)
		if err != nil {
			log.GLogger.Error(err)
		}
	}()
}
