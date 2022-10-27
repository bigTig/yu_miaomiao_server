package request

type InsertBrandReq struct {
	Name       string `json:"name"`       // 品牌名称
	Icon       string `json:"icon"`       // 图标地址
	Sort       uint   `json:"sort"`       // 排序
	Type       uint   `json:"type"`       // 是否推荐
	BrandPrice int    `json:"brandPrice"` // 起始价格
	Status     string `json:"status"`     // 状态 ENABLE 上架 UNABLE 下架
	ShopId     uint   `json:"shopId"`     // 商铺id
	CateId     uint   `json:"cateId"`     // 类目id
}

type UpdateBrandReq struct {
	Id         uint   `json:"id"`         // id
	Name       string `json:"name"`       // 品牌名称
	Icon       string `json:"icon"`       // 图标地址
	Sort       uint   `json:"sort"`       // 排序
	Type       uint   `json:"type"`       // 是否推荐
	BrandPrice int    `json:"brandPrice"` // 起始价格
	Status     string `json:"status"`     // 状态 ENABLE 上架 UNABLE 下架
	ShopId     uint   `json:"shopId"`     // 商铺id
	CateId     uint   `json:"cateId"`     // 类目id
}
