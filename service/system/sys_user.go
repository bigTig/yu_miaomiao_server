package system

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"yuyu/global"
	"yuyu/model/system"
	systemReq "yuyu/model/system/request"
	"yuyu/utils"
)

type UserService struct{}

// Login
//@author: kaifengli
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GvaDb {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.GvaDb.Where("username = ?", u.Username).First(&user).Error
	//查找到用户
	if err == nil {
		//有密码，判断密码是否正确
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}

		return &user, err
	}
	// 还未注册
	return userService.Register(u)
}

// WxLogin
//@author: kaifengli
//@function: Login
//@description: 微信授权登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func (userService *UserService) WxLogin(u *system.SysUser) (userInter *system.SysUser, err error) {
	if nil == global.GvaDb {
		return nil, fmt.Errorf("db not init")
	}

	var user system.SysUser
	err = global.GvaDb.Where("phone = ?", u.Mobile).First(&user).Error
	if err == nil {
		//有密码，判断密码是否正确
	}

	return &user, err
}

// Register
//@author: kaifengli
//@function: Register
//@description: 用户注册
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser
func (userService *UserService) Register(u *system.SysUser) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if !errors.Is(global.GvaDb.Where("mobile = ?", u.Mobile).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("该手机号码已经注册")
	}
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	u.CreatedTime = utils.SetCreatedTime()
	u.UpdatedTime = utils.SetUpdatedTime()
	err = global.GvaDb.Create(&u).Error
	return u, err
}

// SetUserInfo
//@author: kaifengli
//@function: Register
//@description: 设置用户信息
//@param: u *systemReq.ChangeUserInfo
//@return:  err error
func (userService *UserService) SetUserInfo(userInfo *systemReq.ChangeUserInfo) (err error) {
	var user system.SysUser

	err = global.GvaDb.Where("uuid = ?", userInfo.UUID).First(&user).Error
	if err != nil {
		global.GvaLog.Error("该用户不存在")
		return errors.New("该用户不存在")
	}

	user.ID = userInfo.ID
	user.UUID = userInfo.UUID
	user.Name = userInfo.Name
	user.UpdatedTime = utils.SetUpdatedTime()
	user.Mobile = userInfo.Mobile
	user.NickName = userInfo.NickName
	user.Country = userInfo.Country
	user.Province = userInfo.Province
	user.City = userInfo.City
	user.District = userInfo.District
	user.Sex = userInfo.Sex
	user.Avatar = userInfo.Avatar

	return global.GvaDb.Where("id = ?", userInfo.ID).Updates(&user).Error
}
