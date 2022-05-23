/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-05-23 16:53:19
 * @Descripttion:
 */
package common

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
	HTTPPort   int    `json:"http_port"`
	LogDir     string `json:"log_dir"`
	MysqlHost  string `json:"mysql_host"`
	MysqlPort  int    `json:"mysql_port"`
	MysqlUser  string `json:"mysql_user"`
	MysqlPass  string `json:"mysql_pass"`
	MysqlDB    string `json:"mysql_db"`
	CrosConfig bool   `json:"cros_config"`
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
