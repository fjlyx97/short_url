package dao

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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
	gDbService = db
	return nil
}
func (m *MysqlDB) CreateLink(uid int64, url string, availableAt int64) error {
	link := &Links{
		Uid:         uid,
		Url:         url,
		InsertAt:    time.Now().UnixNano(),
		AvailableAt: availableAt,
		Available:   1,
	}
	err := gDbService.Create(&link).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *MysqlDB) FindLink(uid int64) (string, error) {
	link := &Links{
		Uid: uid,
	}
	var links []*Links
	err := gDbService.Find(&links, link).Error
	if err != nil {
		return "", err
	}
	if len(links) != 1 {
		errMsg := fmt.Sprintf("Too many links , num : %s", len(links))
		return "", errors.New(errMsg)
	}

	return links[0].Url, nil
}
