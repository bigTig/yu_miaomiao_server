package system

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"yuyu/global"
	"yuyu/model/system"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
)

type AddressService struct{}

// AddressList
//@author: kaifengli
//@function: AddressList
//@description: 获取收货地址列表
//@param: uuid uuid.UUID
//@return: list interface{}, total int64, err error
func (address *AddressService) AddressList(uuid uuid.UUID) (list interface{}, total int64, err error) {
	var addr []system.SysAddress

	err = global.GvaDb.Where("uuid = ?", uuid).Order("is_default desc, created_time desc").Find(&addr).Count(&total).Error

	return addr, total, err
}

// InsertAddress
//@author: kaifengli
//@function: InsertAddress
//@description: 添加收货地址
//@param: req *systemReq.InsertAddressParams, uuid uuid.UUID
//@return: err error
func (address *AddressService) InsertAddress(req systemReq.InsertAddressParams, uuid uuid.UUID) (err error) {
	if req.IsDefault == 1 {
		// 修改所有的默认地址
		err = global.GvaDb.Model(&system.SysAddress{}).Where("uuid = ?", uuid).UpdateColumn("is_default", 0).Error
		if err != nil {
			return errors.New("修改数据时发生异常")
		}
	}
	//	还没有默认地址
	addr := &system.SysAddress{
		Name:         req.Name,
		UUID:         uuid,
		Mobile:       req.Mobile,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
		Province:     req.Province,
		ProvinceCode: req.ProvinceCode,
		City:         req.City,
		CityCode:     req.CityCode,
		District:     req.District,
		DistrictCode: req.DistrictCode,
		Address:      req.Address,
		ReceiveAddr:  req.ReceiveAddr,
		Detailed:     req.Detailed,
		IsDefault:    req.IsDefault,
	}

	addr.UpdatedTime = utils.SetUpdatedTime()
	addr.CreatedTime = utils.SetCreatedTime()

	return global.GvaDb.Create(&addr).Error
}

// UpdateAddress
//@author: kaifengli
//@function: UpdateAddress
//@description: 更新收货地址
//@param: req *systemReq.UpdateAddressParams
//@return: err error
func (address *AddressService) UpdateAddress(req *systemReq.UpdateAddressParams, uuid uuid.UUID) (err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", req.Id).First(&system.SysAddress{}).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	// 如果当前的为默认
	if req.IsDefault == 1 {
		// 修改所有的默认地址
		err = global.GvaDb.Model(&system.SysAddress{}).Where("uuid = ?", uuid).UpdateColumn("is_default", 0).Error
		if err != nil {
			return errors.New("修改数据时发生异常")
		}
	}

	// 2. 更新数据库
	addr := &system.SysAddress{
		Name:         req.Name,
		UUID:         uuid,
		Mobile:       req.Mobile,
		Longitude:    req.Longitude,
		Latitude:     req.Latitude,
		Province:     req.Province,
		ProvinceCode: req.ProvinceCode,
		City:         req.City,
		CityCode:     req.CityCode,
		District:     req.District,
		DistrictCode: req.DistrictCode,
		Address:      req.Address,
		ReceiveAddr:  req.ReceiveAddr,
		Detailed:     req.Detailed,
		IsDefault:    req.IsDefault,
	}
	addr.UpdatedTime = utils.SetUpdatedTime()

	err = global.GvaDb.Where("id = ?", req.Id).Updates(&addr).Error
	return err
}

// DeleteAddress
//@author: kaifengli
//@function: DeleteAddress
//@description: 删除收货地址
//@param: id string
//@return: err error
func (address *AddressService) DeleteAddress(id string, uuid uuid.UUID) (err error) {
	//https://blog.csdn.net/Frame_X/article/details/95184201
	// 1. 根据id查找数据, 判断是否存在
	var currentData system.SysAddress
	err = global.GvaDb.Where("id = ?", id).First(&currentData).Error
	if err != nil {
		return errors.New("当前id 不存在")
	}

	// 当前用户所有数据
	var addr []system.SysAddress
	// 查看是否已经有默认地址了
	var result = global.GvaDb.Where("uuid = ?", uuid).Find(&addr)

	// !errors.Is(result.Error, gorm.ErrRecordNotFound)
	// 两条数据以上
	if result.RowsAffected > 1 && currentData.IsDefault == 1 {
		// 判断当前数据是否为默认
		err = global.GvaDb.Where("uuid = ?", uuid).Not("id = ?", id).Model(&addr).UpdateColumn("is_default", 0).Error
		if err != nil {
			return errors.New("修改数据时发生异常")
		}
		var last system.SysAddress
		err = global.GvaDb.Last(&last).Error
		if err != nil {
			return errors.New("查最新一条数据")
		}

		err = global.GvaDb.Model(&system.SysAddress{}).Where("id = ?", last.ID).UpdateColumn("is_default", 1).Error
		if err != nil {
			return errors.New("修改默认值时发生异常")
		}
	}

	err = global.GvaDb.Where("id = ?", id).Delete(&currentData).Error

	if err != nil {
		return errors.New("查询数据错误")
	}

	return err
}

// AddressDetailById
//@author: kaifengli
//@function: AddressDetailById
//@description: 地址详情
//@param: id string
//@return: addr *system.SysAddress, err error
func (address *AddressService) AddressDetailById(id string) (addr *system.SysAddress, err error) {
	// 1. 根据id查找数据, 判断是否存在
	err = global.GvaDb.Where("id = ?", id).First(&addr).Error

	return addr, err
}
