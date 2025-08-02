package mylog

import (
	"log"
	"os"
)

var Mlogger *log.Logger

func Initlog() {
	file, err := os.OpenFile("./mylog/myblog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("创建文件失败：", err)
	}
	Mlogger = log.New(file, "", log.LstdFlags)
}
