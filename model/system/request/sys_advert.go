package request

type InsertAdvert struct {
	Name     string `json:"name"`     // 广告名称
	Photo    string `json:"photo"`    // 图片地址
	Sort     int    `json:"sort"`     // 排序
	Type     string `json:"type"`     // 广告类型 product 产品 news 资讯 index 首页
	Status   string `json:"status"`   // 状态
	Action   string `json:"action"`   // 链接值
	Position int    `json:"position"` // 广告位置 1首页轮播;
}

type UpdateAdvert struct {
	Id       string `json:"id"`       // id
	Name     string `json:"name"`     // 广告名称
	Photo    string `json:"photo"`    // 图片地址
	Sort     int    `json:"sort"`     // 排序
	Type     string `json:"type"`     // 广告类型 product 产品 news 资讯 index 首页
	Status   string `json:"status"`   // 状态
	Action   string `json:"action"`   // 链接值
	Position int    `json:"position"` // 广告位置 1首页轮播;
}
