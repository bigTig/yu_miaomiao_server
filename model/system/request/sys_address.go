package request

import uuid "github.com/satori/go.uuid"

type InsertAddressParams struct {
	Name         string  `json:"name" gorm:"comment:收货人"`
	Mobile       string  `json:"mobile" gorm:"comment:手机号"`
	Latitude     float32 `json:"latitude" gorm:"comment:经度"`
	Longitude    float32 `json:"longitude" gorm:"comment:纬度"`
	Province     string  `json:"province" gorm:"comment:省"`
	ProvinceCode string  `json:"provinceCode" gorm:"comment:省-编码"`
	City         string  `json:"city" gorm:"comment:市"`
	CityCode     string  `json:"cityCode" gorm:"comment:市-编码"`
	District     string  `json:"district" gorm:"comment:区"`
	DistrictCode string  `json:"districtCode" gorm:"comment:区-编码"`
	ReceiveAddr  string  `json:"receiveAddr" gorm:"收货地址"`
	Address      string  `json:"address" gorm:"comment:详细地址-收货地址（不加省市区）"`
	Detailed     string  `json:"detailed" gorm:"comment:省市区+详细地址"`
	IsDefault    uint    `json:"isDefault" gorm:"DEFAULT 0;comment:是否默认地址 1默认;"`
}

type UpdateAddressParams struct {
	Id           uint      `json:"id"`
	UUID         uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"` // 用户UUID
	Name         string    `json:"name" gorm:"comment:收货人"`
	Mobile       string    `json:"mobile" gorm:"comment:手机号"`
	Latitude     float32   `json:"latitude" gorm:"comment:经度"`
	Longitude    float32   `json:"longitude" gorm:"comment:纬度"`
	Province     string    `json:"province" gorm:"comment:省"`
	ProvinceCode string    `json:"provinceCode" gorm:"comment:省-编码"`
	City         string    `json:"city" gorm:"comment:市"`
	CityCode     string    `json:"cityCode" gorm:"comment:市-编码"`
	District     string    `json:"district" gorm:"comment:区"`
	DistrictCode string    `json:"districtCode" gorm:"comment:区-编码"`
	ReceiveAddr  string    `json:"receiveAddr" gorm:"收货地址"`
	Address      string    `json:"address" gorm:"comment:详细地址-收货地址（不加省市区）"`
	Detailed     string    `json:"detailed" gorm:"comment:收货地址+详细地址"`
	IsDefault    uint      `json:"isDefault" gorm:"DEFAULT 0;comment:是否默认地址 1默认;"`
}
