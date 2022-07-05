/*
 * @Author: Dujingxi
 * @Date: 2022-06-17 10:19:08
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-07-05 11:19:13
 * @Descripttion:
 */
package common

import (
	"database/sql"
	"flag"
	"net/http"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

var (
	// ServerLog  *logman.LogMan
	Config     *Configuration
	DB         *gorm.DB
	Sqldb      *sql.DB
	Err        error
	HTTPClient *http.Client
)

func init() {
	cf := flag.String("f", "conf.json", "specify the config file.")
	flag.Parse()
	Config = new(Configuration)
	// cfile := "conf.json"
	// LoadConfig(cfile, Config)
	LoadConfig(*cf, Config)

	HTTPClient = &http.Client{}

	// rootDir, err := osext.ExecutableFolder()
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	} else {
		Config.RootDir = rootDir
	}
	logDir := Config.LogDir
	if logDir == "" {
		logDir = filepath.Join(Config.RootDir, "log")
		Config.LogDir = logDir
	}
	if !PathExists(logDir) {
		os.MkdirAll(logDir, 0777)
	}

	// Initialize the mysql db
	InitDB()
}
