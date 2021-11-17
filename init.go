package main

import (
	"baseserver/logman"
	"flag"
	"os"
	"path/filepath"
)

func init() {
	cf := flag.String("f", "conf.json", "specify the config file.")
	flag.Parse()
	CONFIG = new(Configuration)
	LoadConfig(*cf, CONFIG)

	logDir := CONFIG.LogDir
	if logDir == "" {
		logDir = "/var/log"
	}
	if !PathExists(logDir) {
		os.MkdirAll(logDir, 0777)
	}
	LOGMAN = logman.NewLogMan(filepath.Join(logDir, "rc.log"))
	LOGMAN.SetSaveMode(logman.BySize)
	LOGMAN.SetSaveVal(50)
}
