package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gDbService *gorm.DB

type DBInterface interface {
	InitDB(ip string, port string, user string, password string, databaseName string,
		ConnectedTimeout string, ReadTimeout string, WriteTimeout string) error
	CreateLink(uid int64, url string, availableAt int64) error
}
