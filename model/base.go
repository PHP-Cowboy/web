package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const TableOptions string = "gorm:table_options"

func GetOptions(tableName string) string {
	return "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci comment '" + tableName + "'"
}

func AutoIncrementOptions(tableName string) string {
	return "ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci comment '" + tableName + "'"
}

type Base struct {
	Id         int       `gorm:"primaryKey;type:int(11) unsigned AUTO_INCREMENT;comment:id"`
	CreateTime time.Time `gorm:"autoCreateTime;type:datetime;comment:创建时间"`
	UpdateTime time.Time `gorm:"autoCreateTime;type:datetime;comment:更新时间"`
}

type Creator struct {
	CreatorId int    `gorm:"type:int(11) unsigned;comment:操作人id"`
	Creator   string `gorm:"type:varchar(32);comment:操作人昵称"`
}

const (
	BatchSize = 100
)

type GormList []string

func (g GormList) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (g *GormList) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

const (
	TimeFormat = "2006-01-02 15:04:05"
)

// MyTime 自定义时间
type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	now, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = MyTime(now)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format(TimeFormat))
	return []byte(stamp), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format(TimeFormat), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return time.Time(*t).Format(TimeFormat)
}
