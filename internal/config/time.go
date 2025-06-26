package config

import (
	"encoding/json"
	"time"
)

type Time time.Time

var timeFormat = "2006-01-02 15:04:05"

func (t *Time) MarshalJSON() ([]byte, error) {
	format := time.Time(*t).Format(timeFormat)
	return []byte(`"` + format + `"`), nil
}

func (t *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	location, err := time.ParseInLocation(timeFormat, s, time.Local)
	if err != nil {
		return err
	}
	*t = Time(location)
	return nil
}
