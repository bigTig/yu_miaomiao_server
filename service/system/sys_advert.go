package system

import (
	"errors"
	"gorm.io/gorm"
	"yuyu/global"
	"yuyu/model/common/request"
	"yuyu/model/system"
	systemReq "yuyu/model/system/request"
)

type AdvertService struct{}

// AdvertList
//@author: kaifengli
//@function: AdvertList
//@description: 获取轮播图列表
//@param: u *model.SysAdvert
//@return: err error, userInter *model.SysAdvert
func (advertService *AdvertService) AdvertList(info request.PageInfo) (list interface{}, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GvaDb.Model(&system.SysAdvert{})
	err = db.Count(&total).Error
	var advert []system.SysAdvert

	err = global.GvaDb.Find(&advert).Error
	err = db.Limit(limit).Offset(offset).Where("status = ?", "ENABLE").Find(&advert).Error

	return advert, total, err
}

// InsertAdvert
//@author: kaifengli
//@function: InsertAdvert
//@description: 添加轮播图
//@param: advert *model.SysAdvert
//@return: err error
func (advertService *AdvertService) InsertAdvert(advert *systemReq.InsertAdvert) error {

	if !errors.Is(global.GvaDb.Where("name = ?", advert.Name).First(&system.SysAdvert{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	adv := &system.SysAdvert{
		Name:     advert.Name,
		Action:   advert.Action,
		Photo:    advert.Photo,
		Type:     advert.Type,
		Sort:     advert.Sort,
		Status:   advert.Status,
		Position: advert.Position,
	}
	return global.GvaDb.Create(&adv).Error
}

// UpdateAdvert
//@author: kaifengli
//@function: InsertAdvert
//@description: 更新轮播图
//@param: advert *model.SysAdvert
//@return: err error
func (advertService *AdvertService) UpdateAdvert(advert *systemReq.UpdateAdvert) error {
	// 1. 根据id查找数据, 判断是否存在

	// 2. 更新数据库
	return nil
}
