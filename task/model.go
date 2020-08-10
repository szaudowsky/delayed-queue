package task

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

type Task struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	CreatedDate *CustomDate `json:"createdDate"`
}

type CustomDate time.Time

const timeFormat = "2006-01-02 15:04:05"

func (t *CustomDate) UnmarshalJSON(b []byte) error {
	newTime, err := time.Parse(timeFormat, strings.Trim(string(b), "\""))
	if err != nil {
		return err
	}

	*t = CustomDate(newTime)
	return nil
}

func (t *CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", time.Time(*t).Format(timeFormat))), nil
}

func (t CustomDate) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *CustomDate) Scan(src interface{}) error {
	if val, ok := src.(time.Time); ok {
		*t = CustomDate(val)
	} else {
		return errors.New("time Scanner passed a non-time object")
	}
	return nil
}
