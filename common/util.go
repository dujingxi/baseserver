package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func EnsureDir(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
}

func SOTPHTTPPOST(url string, d map[string]interface{}) (res map[string]interface{}, err error) {
	// url := common.Config.SotpUrl.LoginUrl
	// d := map[string]interface{}{
	// 	"account":  name,
	// 	"password": pwd,
	// 	"role":     role,
	// }
	dBytes, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	dJson := bytes.NewReader(dBytes)
	req, err := http.NewRequest("POST", url, dJson)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	resp, err := HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(respBytes, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
