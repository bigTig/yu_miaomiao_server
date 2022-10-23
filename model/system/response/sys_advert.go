package response

type SysAdvertResponse struct {
	Name        string `json:"name"`         // 广告名称
	Photo       string `json:"photo"`        // 图片地址
	Sort        int    `json:"sort"`         // 排序
	Type        string `json:"type"`         // 广告类型 product 产品 news 资讯 index 首页
	Action      string `json:"action"`       // 链接值
	Position    int    `json:"position"`     // 广告位置 1首页轮播
	ID          uint   `json:"id" `          // 主键ID
	CreatedTime string `json:"created_time"` // 创建时间
	UpdatedTime string `json:"updated_time"` // 更新时间
	DeletedTime string `json:"deleted_time"` // 删除时间
}
