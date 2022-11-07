package request

import uuid "github.com/satori/go.uuid"

// Login User Login structure
type Login struct {
	UserName  string `json:"userName"`  // 用户名-手机号码
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// WxLogin User wxLogin structure
type WxLogin struct {
	Code          string `json:"code"`          // 微信用户登录凭证（有效期五分钟）
	EncryptedData string `json:"encryptedData"` // 微信用户信息的加密数据
	Iv            string `json:"iv"`            // 加密算法的初始向量
	AppId         string `json:"appid"`         // 微信应用id
}

// WxSessionKeyDto 封装code2session接口返回数据
type WxSessionKeyDto struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// WxPhoneDto 封装手机号信息数据
type WxPhoneDto struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
}

type ChangePasswordReq struct {
	ID              uint   `json:"-"`               // 从 JWT 中提取 user id，避免越权
	Password        string `json:"password"`        // 旧密码
	ConfirmPassword string `json:"confirmPassword"` // 新密码
}

type ChangeUserInfo struct {
	ID           uint      `gorm:"primarykey"`
	UUID         uuid.UUID `json:"uuid" gorm:"index;comment:用户UUID"`          // 用户UUID
	Name         string    `json:"name" gorm:"comment:真实姓名"`                  // 真实姓名
	NickName     string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
	Mobile       string    `json:"mobile"  gorm:"comment:用户手机号"`              // 用户手机号
	Sex          string    `json:"sex" gorm:"default:1;comment:性别 1 男 2 女 3 未知"`
	IdCard       string    `json:"idCard"`
	BirthDay     string    `json:"birthDay"`
	Province     string    `json:"province" gorm:"comment:省"`
	ProvinceCode string    `json:"provinceCode" gorm:"comment:省-编码"`
	City         string    `json:"city" gorm:"comment:市"`
	CityCode     string    `json:"cityCode" gorm:"comment:市-编码"`
	District     string    `json:"district" gorm:"comment:区"`
	DistrictCode string    `json:"districtCode" gorm:"comment:区-编码"`
	Address      string    `json:"address" gorm:"comment:省市区"`
	Detailed     string    `json:"detailed" gorm:"comment:详细地址"`
}
