package common

import (
	"database/sql"
	"flag"
	"net/http"
	"os"
	"path/filepath"
	"service-man/logman"

	"gorm.io/gorm"
)

var (
	ServerLog  *logman.LogMan
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

	logDir := Config.LogDir
	if logDir == "" {
		logDir = "/var/log"
	}
	if !PathExists(logDir) {
		os.MkdirAll(logDir, 0777)
	}

	ServerLog = logman.NewLogMan(filepath.Join(logDir, "server.log"))
	ServerLog.SetSaveMode(logman.BySize)
	ServerLog.SetSaveVal(20)

	// Initialize the mysql db
	InitDB()
}
