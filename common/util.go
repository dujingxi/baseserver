package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
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

func HTTPPOST(url string, d map[string]interface{}) (res map[string]interface{}, err error) {
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

// GetValueString 格式化输出
// 传入需要格化的对象做为参数
// 返回格式化后的值
func GetValueString(v interface{}) (r string) {
	r = ""
	switch v.(type) {
	case []byte:
		r = string(v.([]byte))
	case string:
		r = v.(string)
	case int:
		r = strconv.Itoa(v.(int))
	case int64:
		r = strconv.FormatInt(v.(int64), 10)
	case float64:
		r = fmt.Sprintf("%v", v)
		// r = strconv.FormatFloat(v.(float64), 'E', -1, 64)
	case float32:
		r = fmt.Sprintf("%v", v)
		// r = strconv.FormatFloat(v.(float64), 'E', -1, 32)
	}
	return
}
func GetValueInt(v interface{}) (r int) {
	r = 0
	switch v.(type) {
	case []byte:
		r, _ = strconv.Atoi(string(v.([]byte)))
	case string:
		r, _ = strconv.Atoi(v.(string))
	case int:
		r = v.(int)
	case int64:
		r = int(v.(int64))
	case float64:
		r = int(math.Floor(v.(float64)))
	case float32:
		r = int(math.Floor(float64(v.(float32))))
	}
	return
}

// GetDefValueString 獲取string類型的值, 為空是設置默認
func GetDefValueString(v interface{}, defval string) (r string) {
	r = ""
	switch v.(type) {
	case []byte:
		r = string(v.([]byte))
	case string:
		r = v.(string)
	case int:
		r = strconv.Itoa(v.(int))
	case int64:
		r = strconv.FormatInt(v.(int64), 10)
	case float64:
		// r = strconv.FormatFloat(v.(float64), 'E', -1, 64)
		r = strconv.Itoa(int(v.(float64)))
	case float32:
		// r = strconv.FormatFloat(v.(float64), 'E', -1, 32)
		r = strconv.Itoa(int(v.(float32)))
	case nil:
		r = defval
	}
	return
}
