package system

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
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
			global.GvaLog.Error("密码错误")
			return nil, errors.New("密码错误")
		}

		return &user, err
	}
	// 还未注册
	return userService.Register(u)
}

// Logout
//@author: kaifengli
//@function: Logout
//@description: 退出登录
//@param: jwtList model.JwtBlacklist
//@return:  err error
func (userService *UserService) Logout(jwtList *system.JwtBlacklist) (err error) {
	err = global.GvaDb.Create(&jwtList).Error
	if err != nil {
		global.GvaLog.Error("密码错误", zap.Error(err))
		return err
	}

	return nil
}

// ChangePassword
//@author: kaifengli
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, confirmPassword string
//@return: userInter *model.SysUser,err error
func (userService *UserService) ChangePassword(u *system.SysUser, confirmPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.GvaDb.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(confirmPassword)
	err = global.GvaDb.Save(&user).Error
	return &user, err
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
//@function: SetUserInfo
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

// GetUserInfo
//@author: kaifengli
//@function: GetUserInfo
//@description: 获取用户信息
//@param: uuid uuid.UUID
//@return: err error, userInter *model.SysUser
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (userInfo *system.SysUser, err error) {
	var user system.SysUser

	err = global.GvaDb.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		global.GvaLog.Error("该用户不存在")
		return userInfo, errors.New("该用户不存在")
	}

	return &user, err
}
