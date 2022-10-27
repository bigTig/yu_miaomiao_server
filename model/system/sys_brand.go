package system

import "yuyu/global"

type SysBrand struct {
	global.GvaModel
	Name       string `json:"name" gorm:"comment:品牌名称;"`
	Icon       string `json:"icon" gorm:"comment:图标地址;"`
	Sort       uint   `json:"sort" gorm:"comment:排序;"`
	Type       uint   `json:"type" gorm:"default:0;comment:是否推荐 0 推荐 1 不推荐;"`
	BrandPrice int    `json:"brandPrice" gorm:"default:0;comment:起始价格 分;"`
	Status     string `json:"status" gorm:"default:ENABLE;comment:状态 ENABLE 上架 UNABLE 下架"`
	ShopId     uint   `json:"shopId" gorm:"default:0;comment:商铺id;"`
	CateId     uint   `json:"cateId" gorm:"default:0;comment:类目id;"`
}
