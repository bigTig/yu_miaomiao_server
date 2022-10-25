package system

import (
	"yuyu/global"
	"yuyu/model/common/request"
	"yuyu/model/system"
)

type CategoryService struct{}

// CategoryList
//@author: kaifengli
//@function: CategoryList
//@description: 获取分类列表
//@param: cateReq *model.SysCategory
//@return: list interface{}, total int64, err error
func (category *CategoryService) CategoryList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GvaDb.Model(&system.SysCategory{})
	err = db.Count(&total).Error
	var cate []system.SysCategory

	err = global.GvaDb.Find(&cate).Error
	err = db.Limit(limit).Offset(offset).Where("status = ?", "ENABLE").Find(&cate).Error

	return cate, total, nil
}
