package system

import "yuyu/global"

type SysFastCate struct {
	global.GvaModel
	Name   string `json:"name" gorm:"NOT NULL;COMMENT:分类名称;"`
	Status string `json:"status" gorm:"DEFAULT:ENABLE;COMMENT:状态：ENABLE 启用 UNABLE 禁用"`
}
