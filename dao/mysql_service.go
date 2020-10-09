package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MysqlDB struct {
}

func (m *MysqlDB) InitDB(ip string, port string, user string, password string, databaseName string,
	ConnectedTimeout string, ReadTimeout string, WriteTimeout string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&timeout=%ss&readTimeout=%ss&writeTimeout=%ss",
		user, password, ip, port, databaseName, ConnectedTimeout, ReadTimeout, WriteTimeout)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		errMsg := fmt.Sprintf("Initial database error, err : %s", err)
		return errors.New(errMsg)
	}
	GDBService = db
	return nil
}
