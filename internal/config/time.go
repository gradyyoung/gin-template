package config

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type Time time.Time

var globalTimeFormat = "2006-01-02 15:04:05" // 全局时间格式变量

// SetTimeFormat 设置全局时间格式
func SetTimeFormat(format string) {
	if format != "" {
		globalTimeFormat = format
	}
}

// GetTimeFormat 获取时间格式配置
func GetTimeFormat(config *ApplicationConfig) string {
	if config != nil && config.Time.Format != "" {
		return config.Time.Format
	}
	return globalTimeFormat // 默认格式
}

func (t *Time) MarshalJSON() ([]byte, error) {
	format := time.Time(*t).Format(globalTimeFormat)
	return json.Marshal(format)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}
	location, err := time.ParseInLocation(globalTimeFormat, s, time.Local)
	if err != nil {
		return err
	}
	*t = Time(location)
	return nil
}

// Scan GORM Scanner 接口, 从数据库读取到类型
func (t *Time) Scan(value any) error {

	if v, ok := value.(time.Time); !ok {
		return errors.New("failed to unmarshal CustomTime value")
	} else {
		*t = Time(v)
		return nil
	}
}

// Value GORM Valuer 接口, 保存到数据库
func (t *Time) Value() (driver.Value, error) {
	if time.Time(*t).IsZero() {
		return nil, nil
	}
	return time.Time(*t), nil
}
