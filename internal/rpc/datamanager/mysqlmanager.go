package datamanager

import (
	"contentService/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var dbClient *MysqlHelper
var onceDb sync.Once

type MysqlHelper struct {
	db *gorm.DB
}

func GetMysqlInstance() *MysqlHelper {
	onceDb.Do(func() {
		dbClient = &MysqlHelper{}

		dns := config.Config.Mysql.Mysql
		if dns == "" {
			panic("init mysql err")
		}
		var err error
		dbClient.db, err = gorm.Open(mysql.New(mysql.Config{
			DSN: dns,
		}), &gorm.Config{})
		if err != nil {
			panic("mysql init err:" + err.Error())
		}

		sqlDB, err := dbClient.db.DB()
		if err != nil {
			panic("mysql get db object err:" + err.Error())
		}

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(config.Config.Mysql.MaxIdleConns)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(config.Config.Mysql.MaxOpenConns)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Duration(config.Config.Mysql.ConnMaxLifetime * time.Second.Microseconds()))
	})

	return dbClient
}

func NewMysqlHelper(conf config.GConfig) (*MysqlHelper, func(), error) {
	dbClient = &MysqlHelper{}

	dns := conf.Mysql.Mysql
	if dns == "" {
		return nil, func() {}, fmt.Errorf("init mysql err")
	}
	var err error
	dbClient.db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dns,
	}), &gorm.Config{})
	if err != nil {
		return nil, func() {}, fmt.Errorf("mysql init err:%s", err.Error())
	}

	sqlDB, err := dbClient.db.DB()
	if err != nil {
		return nil, func() {}, fmt.Errorf("mysql get db object err:%s", err.Error())
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(conf.Mysql.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(conf.Mysql.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(conf.Mysql.ConnMaxLifetime * time.Second.Microseconds()))

	return dbClient, func() {}, nil
}
