package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GDBService *gorm.DB

type DBInterface interface {
	InitDB(ip string, port string, user string, password string, databaseName string,
		ConnectedTimeout string, ReadTimeout string, WriteTimeout string) error
}
