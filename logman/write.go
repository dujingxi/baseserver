/*
 * @Author: Dujingxi
 * @Date: 2022-02-14 16:42:44
 * @version: 1.0
 * @LastEditors: Dujingxi
 * @LastEditTime: 2022-07-05 09:28:48
 * @Descripttion:
 */
package logman

import (
	"fmt"
	"os"
	"time"
)

type logFile struct {
	fileFd     *os.File
	fileName   string
	logTime    int64    //
	level      int      // 日志等级
	saveMode   SaveMode // 保存模式
	saveDays   int      // 日志保存天数
	saveWeeks  int
	saveMonths int
	saveSize   int64 // 文件大小, 需要设置 saveMode 为 BySize 生效
}

func (f *logFile) Write(buf []byte) (n int, err error) {
	if f.fileName == "" {
		fmt.Printf("consol: %s", buf)
		return len(buf), nil
	}

	if f.saveMode == BySize {
		fileInfo, err := os.Stat(f.fileName)
		if err != nil {
			f.createLogFile()
			f.logTime = time.Now().Unix()
		} else {
			filesize := fileInfo.Size()
			if f.fileFd == nil ||
				filesize > f.saveSize {
				f.createLogFile()
				f.logTime = time.Now().Unix()
			}
		}
	} else {
		var duration int64
		switch f.saveMode {
		case ByDay:
			duration = int64(f.saveDays) * 24 * 3600
		case ByWeek:
			duration = int64(f.saveWeeks) * 7 * 24 * 3600
		case ByMonth:
			duration = int64(f.saveMonths) * 30 * 24 * 3600
		}
		if f.logTime+duration < time.Now().Unix() {
			f.createLogFile()
			f.logTime = time.Now().Unix()
		}
	}

	if f.fileFd == nil {
		fmt.Printf("log fileFd is nil !\n")
		return len(buf), nil
	}

	return f.fileFd.Write(buf)
}
