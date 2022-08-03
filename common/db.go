/*
 * @Author: Dujingxi
 * @Date: 2022-07-28 10:07:06
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-08-03 14:14:57
 * @Descripttion:
 */
package common

import (
	"bytes"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDB(c *Configuration) *gorm.DB {
	// connect db
	var dsn bytes.Buffer
	dsn.WriteString(c.MysqlUser)
	dsn.WriteString(":")
	dsn.WriteString(c.MysqlPass)
	dsn.WriteString("@(")
	dsn.WriteString(c.MysqlHost)
	dsn.WriteString(":")
	dsn.WriteString(strconv.Itoa(c.MysqlPort))
	dsn.WriteString(")/")
	dsn.WriteString(c.MysqlDB)
	dsn.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	DB, Err := gorm.Open(mysql.Open(dsn.String()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//设置全局表名禁用复数
			SingularTable: true,
		},
	})
	if Err != nil {
		panic(Err)
	}
	Sqldb, Err := DB.DB()
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
	return DB
}
