package logman

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type SaveMode int

const (
	ByDay SaveMode = iota + 1
	ByWeek
	ByMonth
	BySize
)

const (
	FATAL = iota + 1
	ERROR
	WARN
	INFO
	DEBUG
)

type LogMan struct {
	logger1 *log.Logger
	LogObj  *logFile
}

type Fields map[string]string

func joinFields(fs Fields) string {
	// Log format: $time $level $module $serverId $event $userAccount $meetingId $message
	// If a log field is empty, use "-" instead
	// You can define any format you want
	var content = []map[string]string{
		{"key": "time", "val": "-"},
		{"key": "level", "val": "-"},
		{"key": "type", "val": "-"},
		{"key": "message", "val": "-"},
	}
	// set default value for time
	content[0]["val"] = outTime()
	for k, v := range fs {
		for i, m := range content {
			if k == m["key"] {
				content[i]["val"] = v
			}
		}
	}
	ret := ""
	for _, s := range content {
		ret += s["val"]
		ret += " "
	}
	return strings.TrimSpace(ret)
	// return ret
}

func NewLogMan(fileName string) *LogMan {
	var logFd *os.File
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		logFd, err = os.Create(fileName)
		if err != nil {
			panic(fmt.Sprintf("Log file[%s] create failed.", fileName))
		}
	} else {
		logFd, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic(fmt.Sprintf("Log file[%s] create failed.", fileName))
		}
	}

	file := new(logFile)
	file.fileFd = logFd
	file.fileName = fileName
	file.logTime = time.Now().Unix()
	file.level = DEBUG
	file.saveMode = ByDay
	file.saveDays = 10
	file.saveSize = 104857600 // 100M

	l := new(LogMan)
	l.LogObj = file
	l.logger1 = log.New(io.MultiWriter(os.Stdout, file), "", 0)

	return l
}

func (l *LogMan) SetLevel(Level int) {
	l.LogObj.level = Level
}

func (l *LogMan) SetSaveMode(mode SaveMode) {
	l.LogObj.saveMode = mode
}

func (l *LogMan) GetSaveMode() SaveMode {
	return l.LogObj.saveMode
}

func (m SaveMode) String() string {
	var ret string
	switch m {
	case 1:
		ret = "ByDay"
	case 2:
		ret = "ByWeek"
	case 3:
		ret = "ByMonth"
	case 4:
		ret = "BySize"
	default:
		ret = "unknown"
	}
	return ret
}

func (l *LogMan) SetSaveVal(val int) {
	switch l.LogObj.saveMode {
	case ByDay:
		l.LogObj.saveDays = val
	case ByWeek:
		l.LogObj.saveWeeks = val
	case ByMonth:
		l.LogObj.saveMonths = val
	case BySize:
		l.LogObj.saveSize = int64(val) * 1024 * 1024
	}
}

func (l *LogMan) GetSaveVal() interface{} {
	var ret interface{}
	switch l.LogObj.saveMode {
	case ByDay:
		ret = l.LogObj.saveDays
	case ByWeek:
		ret = l.LogObj.saveWeeks
	case ByMonth:
		ret = l.LogObj.saveMonths
	case BySize:
		ret = l.LogObj.saveSize
	}
	return ret
}

func (l *LogMan) SetSaveDays(days int) {
	l.LogObj.saveDays = days
}

func (l *LogMan) SetSaveWeeks(weeks int) {
	l.LogObj.saveWeeks = weeks
}

func (l *LogMan) SetSaveMonths(months int) {
	l.LogObj.saveMonths = months
}

func (l *LogMan) SetSaveSize(size int) {
	l.LogObj.saveSize = int64(size) * 1024 * 1024
}

func (l *LogMan) Print(logContent Fields) {
	file, line := shortFile()
	lmsg := fmt.Sprintf("%s::%d %s", file, line, logContent["message"])
	logContent["message"] = lmsg
	format := joinFields(logContent)
	l.logger1.Output(2, fmt.Sprint(format))
}

func (l *LogMan) Debugf(logContent Fields) {
	if l.LogObj.level >= DEBUG {
		file, line := shortFile()
		logContent["level"] = "DEBUG"
		lmsg := fmt.Sprintf("%s::%d %s", file, line, logContent["message"])
		logContent["message"] = lmsg
		format := joinFields(logContent)
		l.logger1.Output(2, fmt.Sprint(format))
	}
}

func (l *LogMan) Infof(logContent Fields) {
	if l.LogObj.level >= INFO {
		file, line := shortFile()
		logContent["level"] = "INFO"
		lmsg := fmt.Sprintf("%s::%d %s", file, line, logContent["message"])
		logContent["message"] = lmsg
		// logContent["message"] += fmt.Sprintf("%s::%d %s", file, line, logContent["message"])
		format := joinFields(logContent)
		l.logger1.Output(2, fmt.Sprint(format))
	}
}

func (l *LogMan) Warnf(logContent Fields) {
	if l.LogObj.level >= WARN {
		file, line := shortFile()
		logContent["level"] = "WARN"
		lmsg := fmt.Sprintf("%s::%d %s", file, line, logContent["message"])
		logContent["message"] = lmsg
		format := joinFields(logContent)
		l.logger1.Output(2, fmt.Sprint(format))
	}
}

func (l *LogMan) Errorf(logContent Fields) {
	if l.LogObj.level >= ERROR {
		file, line := shortFile()
		logContent["level"] = "ERROR"
		lmsg := fmt.Sprintf("%s::%d %s", file, line, logContent["message"])
		logContent["message"] = lmsg
		format := joinFields(logContent)
		l.logger1.Output(2, fmt.Sprint(format))
	}
}

func (l *LogMan) Fatalf(logContent Fields) {
	if l.LogObj.level >= FATAL {
		file, line := shortFile()
		logContent["level"] = "FATAL"
		lmsg := fmt.Sprintf("%s::%d %s", file, line, logContent["message"])
		logContent["message"] = lmsg
		format := joinFields(logContent)
		l.logger1.Output(2, fmt.Sprint(format))
	}
}

func outTime() string {
	//var z = time.FixedZone("Asia/Shanghai", 8*3600)
	t := time.Now()
	z, _ := t.Zone()
	return fmt.Sprintf("%v(%v)", t.Format("2006-01-02 15:04:05"), z)
}

func shortFile() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	fileSlice := strings.Split(file, "/")
	file = fileSlice[len(fileSlice)-1]
	return file, line
}

// func (l *LogMan)
