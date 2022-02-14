package common

import (
	"bytes"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB() {
	// connect db
	var dsn bytes.Buffer
	dsn.WriteString(Config.MysqlUser)
	dsn.WriteString(":")
	dsn.WriteString(Config.MysqlPass)
	dsn.WriteString("@(")
	dsn.WriteString(Config.MysqlHost)
	dsn.WriteString(":")
	dsn.WriteString(strconv.Itoa(Config.MysqlPort))
	dsn.WriteString(")/")
	dsn.WriteString(Config.MysqlDB)
	dsn.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	DB, Err = gorm.Open(mysql.Open(dsn.String()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//设置全局表名禁用复数
			SingularTable: true,
		},
	})
	if Err != nil {
		panic(Err)
	}
	Sqldb, Err = DB.DB()
	if Err != nil {
		// paylog.Errorf(logman.Fields{
		// 	"message": "Set database pool error.",
		// })
		panic("Set database pool error.")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	Sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	Sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	Sqldb.SetConnMaxLifetime(time.Hour)
}
