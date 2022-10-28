package request

type InsertFastCateReq struct {
	Name   string `json:"name" gorm:"NOT NULL;COMMENT:分类名称;"`                          // 分类名称
	Status string `json:"status" gorm:"DEFAULT:ENABLE;COMMENT:状态：ENABLE 启用 UNABLE 禁用"` // 状态：ENABLE 启用 UNABLE 禁用
}

type UpdateFastCateReq struct {
	Id     uint   `json:"id"`
	Name   string `json:"name" gorm:"NOT NULL;COMMENT:分类名称;"`                          // 分类名称
	Status string `json:"status" gorm:"DEFAULT:ENABLE;COMMENT:状态：ENABLE 启用 UNABLE 禁用"` // 状态：ENABLE 启用 UNABLE 禁用
}
