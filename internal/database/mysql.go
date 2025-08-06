package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
	"ygang.top/gin-template/internal/config"
)

// NewDB 创建数据库连接
func NewDB(config *config.ApplicationConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.MySQL.DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败，%s", err.Error())
		return nil
	}
	// 获取底层SQL连接以设置连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("无法获取数据库连接: %s", err.Error())
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(config.MySQL.MaxIdleConns)                                    // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(config.MySQL.MaxOpenConns)                                    // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Duration(config.MySQL.ConnMaxLifetime) * time.Second) // 设置了连接可复用的最大时间
	logrus.Infof("成功连接到数据库！")
	return db
}
