/*
 * @Author: Dujingxi
 * @Date: 2022-07-27 14:01:09
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-08-03 14:18:18
 * @Descripttion:
 */
package main

import (
	"baseserver/common"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosConfig struct {
	ClientConfig constant.ClientConfig
	ServerConfig []constant.ServerConfig
	NacosClient  config_client.IConfigClient
	ConfigParam  vo.ConfigParam
	ConfFile     string
}

func NewNacosConfig(f string) (*NacosConfig, error) {

	nIp := fileConfig.NacosIp
	nScheme := fileConfig.NacosScheme
	nPort := fileConfig.NacosPort
	nNamespaceid := fileConfig.NacosNamespaceId
	nGroup := fileConfig.NacosGroup
	nDataid := fileConfig.NacosDataId
	if nIp == "" || nScheme == "" || nPort == 0 || nGroup == "" || nDataid == "" {
		err := errors.New("Get nacos config from file failed, some param wrong")
		return nil, err
	}
	cConfig := constant.ClientConfig{
		NamespaceId:         nNamespaceid, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              filepath.Join(fileConfig.NacosDir, "log"),
		CacheDir:            filepath.Join(fileConfig.NacosDir, "cache"),
		LogLevel:            "warn",
	}
	sConfig := []constant.ServerConfig{
		{
			IpAddr:      nIp,
			Port:        uint64(nPort),
			ContextPath: "/nacos",
			Scheme:      nScheme,
		},
	}
	// a more graceful way to create config client
	cClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cConfig,
			ServerConfigs: sConfig,
		},
	)

	if err != nil {
		return nil, err
	}

	p := vo.ConfigParam{
		Group:  nGroup,
		DataId: nDataid,
		// Content: "test-listen",
	}

	return &NacosConfig{
		ClientConfig: cConfig,
		ServerConfig: sConfig,
		NacosClient:  cClient,
		ConfigParam:  p,
		ConfFile:     f,
	}, nil
}

func (n *NacosConfig) GetConfigs() (string, error) {
	//get config
	n.ConfigParam.Content = ""
	content, err := n.NacosClient.GetConfig(n.ConfigParam)
	if err != nil {
		fmt.Println("nacos GetConfig error: " + err.Error())
		return "", err
	}
	// fmt.Println("GetConfig,config :" + content)
	return content, nil
}

func (n *NacosConfig) PublishConfigs(content string) error {
	n.ConfigParam.Content = content
	_, err := n.NacosClient.PublishConfig(n.ConfigParam)
	return err
}

func (n *NacosConfig) onChangeConf(namespace, group, dataId, content string) {
	// fmt.Println(" 修改后的 content: ", content)
	// 将nacos获取到的配置写入本地文件
	err := ioutil.WriteFile(n.ConfFile, []byte(content), 0666)
	if err == nil {
		fileConfigurationTemp := common.LoadConfigFile(n.ConfFile)
		if fileConfigurationTemp != nil {
			*settingConfig = *fileConfigurationTemp
		}
	}
}

func (n *NacosConfig) ListenConfigs() error {
	n.ConfigParam.Content = ""
	n.ConfigParam.OnChange = n.onChangeConf
	err := n.NacosClient.ListenConfig(n.ConfigParam)
	return err
}

// 获取　nacos 配置，如失败则使用本地文件配置
func HandleConfig(conf string) *common.Configuration {
	// 读取本地文件配置
	fileConfig = common.LoadConfigFile(conf)
	if fileConfig == nil {
		panic("Config file load failed")
	}

	// 创建nacos连接
	n, err := NewNacosConfig(conf)
	if err != nil {
		return fileConfig
	}
	content, err := n.GetConfigs()
	if err != nil || content == "" {
		fstr, err := ioutil.ReadFile(conf)
		if err != nil {
			return fileConfig
		}
		err = n.PublishConfigs(string(fstr))
		if err == nil {
			n.ListenConfigs()
		}
		return fileConfig
	} else {
		// 将nacos获取到的配置写入本地文件
		err = ioutil.WriteFile(conf, []byte(content), 0666)
		if err != nil {
			return fileConfig
		}
		// 写入后再次读取本地文件配置
		fileConfiguration2 := common.LoadConfigFile(conf)
		if fileConfiguration2 == nil {
			return fileConfig
		}
		return fileConfiguration2
	}
}
