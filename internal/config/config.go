package config

import (
	"os"
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	Server Server     `yaml:"server" mapstructure:"server"`
	MySQL  MySQL      `yaml:"mysql" mapstructure:"mysql"`
	Redis  Redis      `yaml:"redis" mapstructure:"redis"`
	Log    Log        `yaml:"log" mapstructure:"log"`
	Time   TimeConfig `yaml:"time" mapstructure:"time"`
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
	LogLevel        string `yaml:"log_level" mapstructure:"log_level"`
}

type Redis struct {
	Host      string `yaml:"host" mapstructure:"host"`
	Port      string `yaml:"port" mapstructure:"port"`
	Password  string `yaml:"password" mapstructure:"password"`
	DB        int    `yaml:"db" mapstructure:"db"`
	KeyPrefix string `yaml:"key_prefix" mapstructure:"key_prefix"`
	// 连接池配置
	PoolSize     int `yaml:"pool_size" mapstructure:"pool_size"`
	MinIdleConns int `yaml:"min_idle_conns" mapstructure:"min_idle_conns"`
	// 超时配置（秒）
	DialTimeout  int `yaml:"dial_timeout" mapstructure:"dial_timeout"`
	ReadTimeout  int `yaml:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout" mapstructure:"write_timeout"`
	PoolTimeout  int `yaml:"pool_timeout" mapstructure:"pool_timeout"`
	// 重试策略
	MaxRetries      int `yaml:"max_retries" mapstructure:"max_retries"`
	MinRetryBackoff int `yaml:"min_retry_backoff" mapstructure:"min_retry_backoff"` // 毫秒
	MaxRetryBackoff int `yaml:"max_retry_backoff" mapstructure:"max_retry_backoff"` // 毫秒
}

type Log struct {
	Output          string `yaml:"output" mapstructure:"output"`
	Level           string `yaml:"level" mapstructure:"level"`
	Format          string `yaml:"format" mapstructure:"format"`
	ReportCaller    bool   `yaml:"report_caller" mapstructure:"report_caller"`
	HideKeys        bool   `yaml:"hide_keys" mapstructure:"hide_keys"`
	ShowFullLevel   bool   `yaml:"show_full_level" mapstructure:"show_full_level"`
	NoColors        bool   `yaml:"no_colors" mapstructure:"no_colors"`
	TimestampFormat string `yaml:"timestamp_format" mapstructure:"timestamp_format"`
}

type TimeConfig struct {
	Format string `yaml:"format" mapstructure:"format"`
}

// InitLogger 初始化日志配置
func InitLogger(logConfig Log) {
	// 设置输出
	switch strings.ToLower(logConfig.Output) {
	case "stdout":
		logrus.SetOutput(os.Stdout)
	case "stderr":
		logrus.SetOutput(os.Stderr)
	default:
		logrus.SetOutput(os.Stdout)
	}

	// 设置日志级别
	switch strings.ToLower(logConfig.Level) {
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	// 设置是否显示调用者信息
	logrus.SetReportCaller(logConfig.ReportCaller)

	// 设置日志格式
	switch strings.ToLower(logConfig.Format) {
	case "nested":
		logrus.SetFormatter(&nested.Formatter{
			HideKeys:        logConfig.HideKeys,
			TimestampFormat: logConfig.TimestampFormat,
			ShowFullLevel:   logConfig.ShowFullLevel,
			NoColors:        logConfig.NoColors,
		})
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: logConfig.TimestampFormat,
		})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: logConfig.TimestampFormat,
			FullTimestamp:   true,
		})
	default:
		logrus.SetFormatter(&nested.Formatter{
			HideKeys:        logConfig.HideKeys,
			TimestampFormat: logConfig.TimestampFormat,
			ShowFullLevel:   logConfig.ShowFullLevel,
			NoColors:        logConfig.NoColors,
		})
	}
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

	// 设置全局时间格式
	if config.Time.Format != "" {
		SetTimeFormat(config.Time.Format)
	}

	// 初始化日志配置
	InitLogger(config.Log)

	logrus.Infoln("配置文件加载完成！")
	return &config
}
