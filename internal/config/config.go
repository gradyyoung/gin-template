package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	Server Server `yaml:"server" mapstructure:"server"`
	MySQL  MySQL  `yaml:"mysql" mapstructure:"mysql"`
	Redis  Redis  `yaml:"redis" mapstructure:"redis"`
}

type Server struct {
	Port int  `yaml:"port" mapstructure:"port"`
	Auth Auth `yaml:"auth" mapstructure:"auth"`
}

type Auth struct {
	Header       string   `yaml:"header" mapstructure:"header"`
	TokenExpired int      `yaml:"token_expired" mapstructure:"token_expired"` // token 过期时间（分钟）
	ExcludeUrls  []string `yaml:"exclude_urls" mapstructure:"exclude_urls"`
}

type MySQL struct {
	DSN             string `yaml:"dsn" mapstructure:"dsn"`
	MaxIdleConns    int    `yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `yaml:"max_open_conns" mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
}

type Redis struct {
	Host      string `yaml:"host" mapstructure:"host"`
	Port      string `yaml:"port" mapstructure:"port"`
	Password  string `yaml:"password" mapstructure:"password"`
	DB        int    `yaml:"db" mapstructure:"db"`
	KeyPrefix string `yaml:"key_prefix" mapstructure:"key_prefix"`
}

// InitApplicationConfig 初始化项目配置
func InitApplicationConfig() *ApplicationConfig {
	viper.SetConfigName("application")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件失败：%s\n", err.Error())
	}
	var config ApplicationConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Fatalf("解析配置文件失败：%s\n", err.Error())
	}
	logrus.Infoln("配置文件加载完成！")
	return &config
}
