package main

import (
	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

func init() {
	log := logs.NewLogger(1000)
	log.SetLogger("console", "")
}

func main() {
	log.Error("error")
	log.Info("info")
}
