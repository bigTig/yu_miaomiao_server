package system

import (
	uuid "github.com/satori/go.uuid"
	"yuyu/global"
)

type SysUser struct {
	global.GvaModel
	UUID         uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`          // 用户UUID
	Username     string    `json:"userName" gorm:"comment:用户登录名"`             // 用户账号-用户登录名
	Name         string    `json:"name" gorm:"comment:真实姓名"`                  // 真实姓名
	Password     string    `json:"-"  gorm:"comment:用户登录密码"`                  // 用户登录密码
	NickName     string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Avatar       string    `json:"avatar" gorm:"default:null;comment:用户头像"`   // 用户头像
	Mobile       string    `json:"mobile"  gorm:"comment:用户手机号"`              // 用户手机号
	IdCard       string    `json:"idCard"  gorm:"comment:身份证号"`               // 身份证号码
	BirthDay     string    `json:"birthDay"  gorm:"comment:出生日期"`             // 出生日期
	Points       float32   `json:"points" gorm:"default:0;comment:积分"`
	Sex          string    `json:"sex" gorm:"default:1;comment:性别 1 男 2 女 3 未知"`
	Enable       int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Openid       string    `json:"openid" gorm:"comment:授权ID"`
	Country      string    `json:"country" gorm:"comment:国家"`
	Province     string    `json:"province" gorm:"comment:省"`
	ProvinceCode string    `json:"provinceCode" gorm:"comment:省-编码"`
	City         string    `json:"city" gorm:"comment:市"`
	CityCode     string    `json:"cityCode" gorm:"comment:市-编码"`
	District     string    `json:"district" gorm:"comment:区"`
	DistrictCode string    `json:"districtCode" gorm:"comment:区-编码"`
	Address      string    `json:"address" gorm:"comment:省市区"`
	Detailed     string    `json:"detailed" gorm:"comment:详细地址"`
}
