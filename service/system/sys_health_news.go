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

type HealthNewsService struct{}

// HealthNewsList
//@author: kaifengli
//@function: BrandList
//@description: 获取新闻资讯列表
//@param: req request.PageInfo
//@return: list interface{}, total int64, err error
func (healthNews *HealthNewsService) HealthNewsList(req request.PageInfo) (list interface{}, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.GvaDb.Model(&system.SysHealthNews{})
	err = db.Count(&total).Error
	var health []system.SysHealthNews

	err = global.GvaDb.Find(&health).Error
	err = db.Limit(limit).Offset(offset).Where("status = ?", "ENABLE").Find(&health).Error

	return health, total, err
}

// InsertHealthNew
//@author: kaifengli
//@function: InsertHealthNew
//@description: 添加健康资讯
//@param: brand *systemReq.InsertHealthNewReq
//@return: err error
func (healthNews *HealthNewsService) InsertHealthNew(req *systemReq.InsertHealthNewReq) (err error) {
	if !errors.Is(global.GvaDb.Where("title = ?", req.Title).First(&system.SysHealthNews{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复title，请修改title")
	}
	health := &system.SysHealthNews{
		Sort:    req.Sort,
		Status:  req.Status,
		Author:  req.Author,
		Content: req.Content,
		Cover:   req.Cover,
		Title:   req.Title,
	}
	health.UpdatedTime = utils.SetUpdatedTime()
	health.CreatedTime = utils.SetCreatedTime()

	return global.GvaDb.Create(&health).Error
}

// UpdateHealthNew
//@author: kaifengli
//@function: InsertAdvert
//@description: 更新新闻资讯
//@param: systemReq *systemReq.UpdateHealthNewReq
//@return: err error
func (healthNews *HealthNewsService) UpdateHealthNew(req *systemReq.UpdateHealthNewReq) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", req.Id).First(&system.SysHealthNews{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}
	// 2. 更新数据库
	health := &system.SysHealthNews{
		Sort:    req.Sort,
		Status:  req.Status,
		Author:  req.Author,
		Content: req.Content,
		Cover:   req.Cover,
		Title:   req.Title,
	}
	health.UpdatedTime = utils.SetUpdatedTime()

	err = global.GvaDb.Where("id = ?", req.Id).Updates(&health).Error
	return err
}

// DeleteHealthNew
//@author: kaifengli
//@function: DeleteBrand
//@description: 删除新闻资讯
//@param: id string
//@return: err error
func (healthNews *HealthNewsService) DeleteHealthNew(id string) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", id).First(&system.SysHealthNews{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	err = global.GvaDb.Where("id = ?", id).Delete(&system.SysHealthNews{}).Error
	return err
}

// HealthNewDetail
//@author: kaifengli
//@function: HealthNewDetail
//@description: 新闻资讯详情
//@param: id string
//@return: err error
func (healthNews *HealthNewsService) HealthNewDetail(id string) (health system.SysHealthNews, err error) {
	err = global.GvaDb.Where("id = ?", id).First(&health).Error
	return health, err
}
