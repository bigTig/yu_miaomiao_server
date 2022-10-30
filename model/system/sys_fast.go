package system

import "yuyu/global"

type SysFast struct {
	global.GvaModel
	Name    string `json:"name" gorm:"NOT NULL;COMMENT:分类名称;"`
	Status  string `json:"status" gorm:"DEFAULT:ENABLE;COMMENT:状态：ENABLE 启用 UNABLE 禁用;"`
	CateId  uint   `json:"cateId" gorm:"DEFAULT:0;COMMENT:分类id;"`
	CanEat  uint   `json:"canEat" gorm:"DEFAULT:0;COMMENT:程度 0 禁食 1 慎食 2 可食;"`
	CarDog  uint   `json:"carDog" gorm:"DEFAULT:0;COMMENT:所属 0 猫 1 狗 2 两者;"`
	Content string `json:"content" gorm:"COMMENT:详细描述;"`
	Icon    string `json:"icon" gorm:"COMMENT:图标;"`
}

type SysFastCate struct {
	global.GvaModel
	Name   string `json:"name" gorm:"NOT NULL;COMMENT:分类名称;"`
	Status string `json:"status" gorm:"DEFAULT:ENABLE;COMMENT:状态：ENABLE 启用 UNABLE 禁用;"`
}
