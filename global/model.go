package global

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type GvaModel struct {
	ID          uint       `json:"id" gorm:"primarykey;comment:主键ID"`                         // 主键ID
	CreatedTime string     `json:"created_time" gorm:"column:created_time"`                   // 创建时间
	UpdatedTime string     `json:"updated_time" gorm:"column:updated_time"`                   // 更新时间
	DeletedTime *LocalTime `json:"deleted_time" sql:"index" gorm:"column:deleted_time;index"` // 删除时间
}

type LocalTime struct {
	time.Time
}

// MarshalJSON 读取数据时将将时间数据格式化
func (l *LocalTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", l.Format("2006-01-02 15:04:05"))), nil
}

// Value 在存储时调⽤，将该⽅法的返回值进⾏存储，该⽅法可以实现数据存储前对数据进⾏相关操作
func (l LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if l.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return l.Time, nil
}

// Scan 实现在数据查询出来之前对数据进⾏相关操作
func (l *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*l = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
