/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-08-05 11:33:24
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
	"path/filepath"
	"strings"
)

type Configuration struct {
	HttpTls          bool   `json:"http_tls"`
	TlsCrt           string `json:"tls_crt"`
	TlsKey           string `json:"tls_key"`
	HTTPBind         string `json:"http_bind"`
	HTTPPort         int    `json:"http_port"`
	MysqlHost        string `json:"mysql_host"`
	MysqlPort        int    `json:"mysql_port"`
	MysqlUser        string `json:"mysql_user"`
	MysqlPass        string `json:"mysql_pass"`
	MysqlDB          string `json:"mysql_db"`
	CrosConfig       bool   `json:"cros_config"`
	RootDir          string `json:"root_dir"`
	LogDir           string `json:"log_dir"`
	NacosDir         string `json:"nacos_dir"`
	NacosIp          string `json:"nacos_ip"` // about nacos
	NacosScheme      string `json:"nacos_scheme"`
	NacosPort        int    `json:"nacos_port"`
	NacosNamespaceId string `json:"nacos_namespace_id"`
	NacosGroup       string `json:"nacos_group"`
	NacosDataId      string `json:"nacos_data_id"`
	NacosUsername    string `json:"nacos_username"`
	NacosPassword    string `json:"nacos_password"`
}

func (c *Configuration) Map() map[string]interface{} {
	bs, err := json.Marshal(c)
	if err != nil {
		return map[string]interface{}{
			"err": err.Error(),
		}
	}
	var m = make(map[string]interface{})
	json.Unmarshal(bs, &m)
	return m
}

func (c *Configuration) String() string {
	bs, err := json.Marshal(c)
	if err != nil {
		return fmt.Sprintf("{\"err\": \"%s\"}", string(bs))
	}
	return string(bs)
}

func (c *Configuration) InitSomeParm() {
	// rootDir, err := osext.ExecutableFolder()
	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	} else {
		c.RootDir = rootDir
	}
	logDir := c.LogDir
	if logDir == "" {
		logDir = filepath.Join(c.RootDir, "log")
		c.LogDir = logDir
	}
	nacosDir := c.NacosDir
	if nacosDir == "" {
		nacosDir = filepath.Join(c.RootDir, "nacos")
		c.NacosDir = nacosDir
	}

	if c.HTTPBind == "" {
		c.HTTPBind = "0.0.0.0"
	}

	if !PathExists(c.LogDir) {
		os.MkdirAll(c.LogDir, 0777)
	}
}

func LoadConfigFile(name string) *Configuration {
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
	var config *Configuration
	err = json.Unmarshal(configStr.Bytes(), &config)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error, %v", err))
	}
	// init some arg, e.g. logdir/nacosdir...
	config.InitSomeParm()
	return config
}
