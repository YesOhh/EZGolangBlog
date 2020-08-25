package mylog

import (
	"goBlog/initialization"
	"goBlog/util"
	"log"
	"os"
)

var MyLogger *log.Logger

func init()  {
	logDir := initialization.Conf.Setting.LogDir
	if util.Exists(logDir) {
		if !util.IsDir(logDir) {
			err := os.MkdirAll(logDir, os.ModePerm)
			if err != nil {
				log.Fatal("无法创建存放日志的文件夹，请手动创建", err)
			}
		}
	} else {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			log.Fatal("无法创建存放日志的文件夹，请手动创建", err)
		}
	}
	logFile, err := os.OpenFile(logDir + string(os.PathSeparator) + "mylog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal("无法打开存放日志的文件，请手动创建", err)
	}
	MyLogger = log.New(logFile, "", log.Lshortfile|log.Ldate|log.Ltime)
}
