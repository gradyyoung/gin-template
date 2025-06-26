package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"os"
	"time"
	"ygang.top/gin-template/cmd/wire"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:        true,          // 隐藏键
		TimestampFormat: time.DateTime, // 时间格式
		ShowFullLevel:   false,         // 显示完整级别名称
		NoColors:        false,         // 不使用颜色
	})
}

func main() {
	wire.InitApplication()
}
