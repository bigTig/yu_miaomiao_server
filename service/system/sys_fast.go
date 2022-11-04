package system

import (
	"errors"
	"gorm.io/gorm"
	"yuyu/global"
	"yuyu/model/system"
	systemReq "yuyu/model/system/request"
	systemRes "yuyu/model/system/response"
	"yuyu/utils"
)

type FastService struct{}

// FastCateList
//@author: kaifengli
//@function: FastCateList
//@description: 获取禁食分类列表
//@param: req request.PageInfo
//@return: list interface{}, total int64, err error
func (fastCate *FastService) FastCateList() (list interface{}, total int64, err error) {
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
func (fastCate *FastService) InsertFastCate(req *systemReq.InsertFastCateReq) (err error) {
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
func (fastCate *FastService) UpdateFastCate(req *systemReq.UpdateFastCateReq) (err error) {
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
//@function: DeleteFastCate
//@description: 删除禁食分类
//@param: id string
//@return: err error
func (fastCate *FastService) DeleteFastCate(id string) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", id).First(&system.SysFastCate{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	err = global.GvaDb.Where("id = ?", id).Delete(&system.SysFastCate{}).Error

	return err
}

// FastList
//@author: kaifengli
//@function: BrandList
//@description: 获取禁食列表
//@param: req request.PageInfo
//@return: list interface{}, total int64, err error
func (fastCate *FastService) FastList(req *systemReq.FastListReq) (list interface{}, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.GvaDb.Model(&system.SysFast{})
	err = db.Count(&total).Error
	var fast []system.SysFast

	err = global.GvaDb.Find(&fast).Error
	err = db.Limit(limit).Offset(offset).Order("created_time desc").Where("status = ? AND cate_id = ?", "ENABLE", req.CateId).Find(&fast).Error

	return fast, total, err
}

// FastDetailById
//@author: kaifengli
//@function: BrandList
//@description: 根据id获取禁食详情
//@param: id string
//@return: fast *system.SysFast, err error
func (fastCate *FastService) FastDetailById(id string) (*systemRes.FastDetailRes, error) {
	var detail systemRes.FastDetailRes
	var fast system.SysFast
	err := global.GvaDb.Where("id = ?", id).First(&fast).Error
	if err != nil {
		global.GvaLog.Error("该记录不存在")
		return &detail, errors.New("该记录不存在")
	}

	var cate system.SysFastCate
	err = global.GvaDb.Where("id = ?", fast.CateId).Find(&cate).Error
	if err != nil {
		global.GvaLog.Error("该分类不存在")
		return &detail, errors.New("该分类不存在")
	}

	detail.CateName = cate.Name
	detail.CateId = cate.ID
	detail.Name = fast.Name
	detail.Status = fast.Status
	detail.CarDog = fast.CarDog
	detail.CanEat = fast.CanEat
	detail.Content = fast.Content
	detail.ID = fast.ID
	detail.Icon = fast.Icon
	detail.UpdatedTime = fast.UpdatedTime
	detail.CreatedTime = fast.CreatedTime

	return &detail, err
}

// InsertFast
//@author: kaifengli
//@function: InsertFast
//@description: 添加禁食
//@param: brand *systemReq.InsertFastReq
//@return: err error
func (fastCate *FastService) InsertFast(req *systemReq.InsertFastReq) (err error) {
	if !errors.Is(global.GvaDb.Where("name = ?", req.Name).First(&system.SysFast{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	fast := &system.SysFast{
		Name:    req.Name,
		Status:  req.Status,
		Icon:    req.Icon,
		CateId:  req.CateId,
		Content: req.Content,
		CanEat:  req.CanEat,
		CarDog:  req.CarDog,
	}
	fast.UpdatedTime = utils.SetUpdatedTime()
	fast.CreatedTime = utils.SetCreatedTime()

	return global.GvaDb.Create(&fast).Error
}

// UpdateFast
//@author: kaifengli
//@function: UpdateFast
//@description: 更新禁食
//@param: req *systemReq.UpdateFastReq
//@return: err error
func (fastCate *FastService) UpdateFast(req *systemReq.UpdateFastReq) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", req.Id).First(&system.SysFast{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	if !errors.Is(global.GvaDb.Where("name = ?", req.Name).First(&system.SysFast{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	// 2. 更新数据库
	fast := &system.SysFast{
		Name:    req.Name,
		Status:  req.Status,
		Icon:    req.Icon,
		CateId:  req.CateId,
		Content: req.Content,
		CanEat:  req.CanEat,
		CarDog:  req.CarDog,
	}
	fast.UpdatedTime = utils.SetUpdatedTime()
	fast.CreatedTime = utils.SetCreatedTime()

	err = global.GvaDb.Where("id = ?", req.Id).Updates(&fast).Error
	return err
}

// DeleteFast
//@author: kaifengli
//@function: DeleteFast
//@description: 删除禁食分类
//@param: id string
//@return: err error
func (fastCate *FastService) DeleteFast(id string) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", id).First(&system.SysFast{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	err = global.GvaDb.Where("id = ?", id).Delete(&system.SysFast{}).Error

	return err
}
