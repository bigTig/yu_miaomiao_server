package request

type InsertHealthNewReq struct {
	Title   string `json:"title" gorm:"comment:标题;"`
	Cover   string `json:"cover" gorm:"comment:封面图片;"`
	Sort    uint   `json:"sort" gorm:"default:0;comment:排序;"`
	Content string `json:"content" gorm:"comment:资讯内容;"`
	Status  string `json:"status" gorm:"default:ENABLE;comment:状态 ENABLE 启用 UNABLE 禁用"`
	Author  string `json:"author" gorm:"comment:作者;"`
}

type UpdateHealthNewReq struct {
	Id      uint   `json:"id"` // id
	Title   string `json:"title" gorm:"comment:标题;"`
	Cover   string `json:"cover" gorm:"comment:封面图片;"`
	Sort    uint   `json:"sort" gorm:"default:0;comment:排序;"`
	Content string `json:"content" gorm:"comment:资讯内容;"`
	Status  string `json:"status" gorm:"default:ENABLE;comment:状态 ENABLE 启用 UNABLE 禁用"`
	Author  string `json:"author" gorm:"comment:作者;"`
}
