package models

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

const (
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)

type Mode struct {
	ID        uint `gorm:"primarykey;comment:'Self-increment numbering'" json:"id" form:"id"`
	CreatedAt LocalTime
	UpdatedAt LocalTime
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type LocalTime struct {
	time.Time
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	// ""空值不进行解析
	if len(data) == 2 {
		*t = LocalTime{Time: time.Time{}}
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+SecLocalTimeFormat+`"`, string(data))
	*t = LocalTime{Time: now}
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(SecLocalTimeFormat))
	return []byte(output), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	/*
		gorm 写入 mysql 时调用
	*/
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to LocalTime", v)
}

func (t LocalTime) DateString() string {
	return t.Format(DateLocalTimeFormat)
}
