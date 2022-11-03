package datamanager

import (
	"contentService/pkg/config"
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
		sqlDB.SetMaxIdleConns(16)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(256)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return dbClient
}
