package system

import "yuyu/global"

type SysAdvert struct {
	global.GvaModel
	Name     string `json:"name" gorm:"comment:广告名称;"`
	Photo    string `json:"photo" gorm:"comment:图片地址;"`
	Sort     int    `json:"sort" gorm:"default:0;comment:排序;"`
	Type     string `json:"type" gorm:"default:news;comment:广告类型 product 产品 news 资讯 index 首页;"`
	Status   string `json:"status" gorm:"default:ENABLE;comment:状态 ENABLE 启用 UNABLE 禁用"`
	Action   string `json:"action" gorm:"comment:链接值;"`
	Position int    `json:"position" gorm:"default:1;comment:广告位置 1首页轮播;"`
}
