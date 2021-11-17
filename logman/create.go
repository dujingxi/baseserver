package logman

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func (me *logFile) createLogFile() {
	var lastSlash int
	logdir := "./"
	if index := strings.LastIndex(me.fileName, "/"); index != -1 {
		logdir = me.fileName[0:index] + "/"
		os.MkdirAll(me.fileName[0:index], os.ModePerm)
		lastSlash = index
	}

	now := time.Now()
	filename := fmt.Sprintf("%s_%04d%02d%02d_%02d%02d",
		me.fileName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
	if err := os.Rename(me.fileName, filename); err == nil {
		go func() {
			os.Chdir(logdir)
			var fName string
			if lastSlash == 0 {
				fName = filename
			} else {
				fName = filename[lastSlash+1:]
			}
			tarCmd := exec.Command("tar", "-zcf", fName+".tgz", fName, "--remove-files")
			tarCmd.Run()

			var rmCmd *exec.Cmd
			switch me.saveMode {
			case ByDay:
				rmCmd = exec.Command("/bin/sh", "-c", "find "+logdir+` -type f -mtime +`+strconv.Itoa(me.saveDays)+` -exec rm {} \;`)
			case ByWeek:
				rmCmd = exec.Command("/bin/sh", "-c", "find "+logdir+` -type f -mtime +`+strconv.Itoa(me.saveWeeks*7)+` -exec rm {} \;`)
			case ByMonth:
				rmCmd = exec.Command("/bin/sh", "-c", "find "+logdir+` -type f -mtime +`+strconv.Itoa(me.saveWeeks*30)+` -exec rm {} \;`)
			default:
				rmCmd = exec.Command("/bin/sh", "-c", "find "+logdir+` -type f -mtime +30 -exec rm {} \;`)
			}
			rmCmd.Run()
		}()
	}

	for index := 0; index < 10; index++ {
		if fd, err := os.OpenFile(me.fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); nil == err {
			me.fileFd.Sync()
			me.fileFd.Close()
			me.fileFd = fd
			break
		} else {
			fmt.Println("Open logfile error! err: ", err.Error())
		}
		me.fileFd = nil
	}
}
