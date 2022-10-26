package system

import (
	"errors"
	"gorm.io/gorm"
	"yuyu/global"
	"yuyu/model/common/request"
	"yuyu/model/system"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
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

// InsertCategory
//@author: kaifengli
//@function: InsertCategory
//@description: 添加类目
//@param: cate *systemReq.InsertCateReq
//@return: err error
func (category *CategoryService) InsertCategory(cate *systemReq.InsertCateReq) (err error) {
	if !errors.Is(global.GvaDb.Where("name = ?", cate.Name).First(&system.SysCategory{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}

	cateParams := &system.SysCategory{
		Name:      cate.Name,
		Pid:       cate.Pid,
		Icon:      cate.Icon,
		Content:   cate.Content,
		Thumbnail: cate.Thumbnail,
		Sort:      cate.Sort,
		Status:    cate.Status,
		Remarks:   cate.Remarks,
	}
	cateParams.UpdatedTime = utils.SetUpdatedTime()
	cateParams.CreatedTime = utils.SetCreatedTime()

	return global.GvaDb.Create(&cateParams).Error
}
