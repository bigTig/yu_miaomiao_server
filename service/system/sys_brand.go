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

type BrandService struct{}

// BrandList
//@author: kaifengli
//@function: BrandList
//@description: 获取品牌列表
//@param: req request.PageInfo
//@return: list interface{}, total int64, err error
func (b *BrandService) BrandList(req request.PageInfo) (list interface{}, total int64, err error) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.GvaDb.Model(&system.SysBrand{})
	err = db.Count(&total).Error
	var brand []system.SysBrand

	err = global.GvaDb.Find(&brand).Error
	err = db.Limit(limit).Offset(offset).Where("status = ?", "ENABLE").Find(&brand).Error

	return brand, total, err
}

// InsertBrand
//@author: kaifengli
//@function: InsertBrand
//@description: 添加品牌
//@param: brand *systemReq.InsertBrandReq
//@return: err error
func (b *BrandService) InsertBrand(brand *systemReq.InsertBrandReq) error {
	if !errors.Is(global.GvaDb.Where("name = ?", brand.Name).First(&system.SysBrand{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	sysBrand := &system.SysBrand{
		Name:       brand.Name,
		Icon:       brand.Icon,
		CateId:     brand.CateId,
		ShopId:     brand.ShopId,
		BrandPrice: brand.BrandPrice,
		Type:       brand.Type,
		Sort:       brand.Sort,
		Status:     brand.Status,
	}
	sysBrand.UpdatedTime = utils.SetUpdatedTime()
	sysBrand.CreatedTime = utils.SetCreatedTime()

	return global.GvaDb.Create(&sysBrand).Error
}

// UpdateBrand
//@author: kaifengli
//@function: InsertAdvert
//@description: 更新品牌
//@param: req *systemReq.UpdateBrandReq
//@return: err error
func (b *BrandService) UpdateBrand(req *systemReq.UpdateBrandReq) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", req.Id).First(&system.SysBrand{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}
	// 2. 更新数据库
	brand := &system.SysBrand{
		Name:       req.Name,
		Type:       req.Type,
		Icon:       req.Icon,
		Sort:       req.Sort,
		Status:     req.Status,
		ShopId:     req.ShopId,
		CateId:     req.CateId,
		BrandPrice: req.BrandPrice,
	}
	brand.UpdatedTime = utils.SetUpdatedTime()

	err = global.GvaDb.Where("id = ?", req.Id).Updates(&brand).Error
	return err
}

// DeleteBrand
//@author: kaifengli
//@function: DeleteBrand
//@description: 删除品牌
//@param: id string
//@return: err error
func (b *BrandService) DeleteBrand(id string) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", id).First(&system.SysBrand{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	err = global.GvaDb.Where("id = ?", id).UpdateColumns(&system.SysBrand{
		Status: "UNABLE",
		GvaModel: global.GvaModel{
			UpdatedTime: utils.SetUpdatedTime(),
		},
	}).Error
	return err
}
