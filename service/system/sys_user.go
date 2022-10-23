package system

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
	"yuyu/global"
	"yuyu/model/system"
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
	u.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
	u.UpdatedTime = time.Now().Format("2006-01-02 15:04:05")
	err = global.GvaDb.Create(&u).Error
	return u, err
}
