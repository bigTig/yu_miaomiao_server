package system

import (
	"errors"
	"gorm.io/gorm"
	"yuyu/global"
	"yuyu/model/system"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
)

type FastCateService struct{}

// FastCateList
//@author: kaifengli
//@function: FastCateList
//@description: 获取禁食分类列表
//@param: req request.PageInfo
//@return: list interface{}, total int64, err error
func (fastCate *FastCateService) FastCateList() (list interface{}, total int64, err error) {
	db := global.GvaDb.Model(&system.SysFastCate{})
	err = db.Count(&total).Error
	var cate []system.SysFastCate

	err = global.GvaDb.Find(&cate).Error
	err = db.Where("status = ?", "ENABLE").Find(&cate).Error

	return cate, total, err
}

// InsertFastCate
//@author: kaifengli
//@function: InsertFastCate
//@description: 添加禁食分类
//@param: brand *systemReq.InsertFastCateReq
//@return: err error
func (fastCate *FastCateService) InsertFastCate(req *systemReq.InsertFastCateReq) (err error) {
	if !errors.Is(global.GvaDb.Where("name = ?", req.Name).First(&system.SysFastCate{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	cate := &system.SysFastCate{
		Name:   req.Name,
		Status: req.Status,
	}
	cate.UpdatedTime = utils.SetUpdatedTime()
	cate.CreatedTime = utils.SetCreatedTime()

	return global.GvaDb.Create(&cate).Error
}

// UpdateFastCate
//@author: kaifengli
//@function: InsertAdvert
//@description: 更新禁食分类
//@param: req *systemReq.UpdateFastCateReq
//@return: err error
func (fastCate *FastCateService) UpdateFastCate(req *systemReq.UpdateFastCateReq) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", req.Id).First(&system.SysFastCate{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}
	if !errors.Is(global.GvaDb.Where("name = ?", req.Name).First(&system.SysFastCate{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	// 2. 更新数据库
	cate := &system.SysFastCate{
		Name:   req.Name,
		Status: req.Status,
	}
	cate.UpdatedTime = utils.SetUpdatedTime()

	err = global.GvaDb.Where("id = ?", req.Id).Updates(&cate).Error
	return err
}

// DeleteFastCate
//@author: kaifengli
//@function: DeleteBrand
//@description: 删除禁食分类
//@param: id string
//@return: err error
func (fastCate *FastCateService) DeleteFastCate(id string) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", id).First(&system.SysFastCate{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	err = global.GvaDb.Where("id = ?", id).UpdateColumns(&system.SysFastCate{
		Status: "UNABLE",
		GvaModel: global.GvaModel{
			UpdatedTime: utils.SetUpdatedTime(),
		},
	}).Error

	return err
}
