package config

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

type Time time.Time

var timeFormat = "2006-01-02 15:04:05"

func (t *Time) MarshalJSON() ([]byte, error) {
	format := time.Time(*t).Format(timeFormat)
	return json.Marshal(format)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return errors.Wrap(err, "json.Unmarshal")
	}
	location, err := time.ParseInLocation(timeFormat, s, time.Local)
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
