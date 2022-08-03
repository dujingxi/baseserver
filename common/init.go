/*
 * @Author: Dujingxi
 * @Date: 2022-06-17 10:19:08
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-08-03 14:12:17
 * @Descripttion:
 */
package common

import (
	"net/http"
)

var (
	// ServerLog  *logman.LogMan
	// ConfigFile *Configuration
	// DB         *gorm.DB
	// Sqldb      *sql.DB
	// Err        error
	HTTPClient *http.Client
)

func init() {
	// cf := flag.String("f", "conf.json", "specify the config file.")
	// flag.Parse()
	// ConfigFile = LoadConfigFile(*cf)

	HTTPClient = &http.Client{}

	// // Initialize the mysql db
	// InitDB()
}
