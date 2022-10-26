package request

type InsertCateReq struct {
	Pid       uint   `json:"pid" gorm:"default:0;comment:父级分类id;"`
	Name      string `json:"name" gorm:"comment:栏目名称;"`
	Icon      string `json:"icon" gorm:"comment:图标地址;"`
	Sort      uint   `json:"sort" gorm:"default:0;comment:排序;"`
	Content   string `json:"content" gorm:"comment:栏目简介;"`
	Status    string `json:"status" gorm:"default:ENABLE;comment:状态 ENABLE 上架 UNABLE 下架"`
	Thumbnail string `json:"thumbnail" gorm:"comment:缩略图;"`
	Remarks   string `json:"remarks" gorm:"comment:备注;"`
}
