package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Configuration struct {
	HTTPPort  int    `json:"http_port"`
	McServer  string `json:"mc_server"`
	MsmServer string `json:"msm_server"`
	LogDir    string `json:"log_dir"`
}

func LoadConfig(name string, config *Configuration) {
	_, e := os.Lstat(name)
	notExist := os.IsNotExist(e)
	if notExist {
		log.Fatal("Config file not found, please specify config use -f.")
		// panic("Config file not found, make sure the current directory has file conf.json.")
	}
	fbytes, err := ioutil.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintf("Read config file error, %v", err))
	}
	fstring := strings.Split(string(fbytes), "\n")
	var configStr bytes.Buffer
	// dsn.WriteString(Config.MysqlUser)
	for _, fs := range fstring {
		fline := strings.TrimSpace(fs)
		if !strings.HasPrefix(fline, "//") {
			configStr.Write([]byte(fline))
		}
	}

	err = json.Unmarshal(configStr.Bytes(), &config)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error, %v", err))
	}
}
