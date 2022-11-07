package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/common/response"
	"yuyu/model/system"
	systemReq "yuyu/model/system/request"
	systemRes "yuyu/model/system/response"
	"yuyu/utils"
)

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body systemReq.Login true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.LoginVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}
	if store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.UserName, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.GvaLog.Error("登录失败, 该用户不存在或者密码错误", zap.Error(err))
			response.FailWithInternalServerError("登录失败, 该用户不存在或者密码错误", c)
			return
		}
		if user.Enable != 1 {
			global.GvaLog.Error("登录失败, 该用户被被禁止登录")
			response.FailWithInternalServerError("登录失败, 该用户被被禁止登录", c)
			return
		}
		b.TokenNext(c, *user)
		return
	}

	response.FailWithInternalServerError("验证码错误", c)
}

// Logout
// @Tags     SysUser
// @Summary  退出登录
// @Produce   application/json
// @Success  200   {object}  response.Response{data=bool,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /user/logout [post]
func (b *BaseApi) Logout(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}

	err := userService.Logout(&jwt)
	if err != nil {
		global.GvaLog.Error("退出登录失败", zap.Error(err))
		response.FailWithMessage("退出登录失败!", c)
		return
	}
	response.OkWithMessage("退出登录成功", c)
}

// ChangePassword
// @Tags      SysUser
// @Summary   用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      systemReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{data=bool,msg=string}  "用户修改密码"
// @Router    /user/changePassword [post]
func (b *BaseApi) ChangePassword(c *gin.Context) {
	var req systemReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithBadRequest(err.Error()+" 参数json 格式", c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithBadRequest(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)
	u := &system.SysUser{GvaModel: global.GvaModel{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.ConfirmPassword)
	if err != nil {
		global.GvaLog.Error("修改失败!", zap.Error(err))
		response.FailWithInternalServerError("修改失败，原密码与当前账户不符", c)
		return
	}

	token := c.Request.Header.Get("x-token")
	jwt := system.JwtBlacklist{Jwt: token}
	err = jwtService.JsonInBlacklist(&jwt)
	if err != nil {
		global.GvaLog.Error("token 加入黑名单失败!", zap.Error(err))
		response.FailWithInternalServerError("token 加入黑名单失败", c)
		return
	}

	response.OkWithMessage("修改成功", c)
}

// SetUserInfo
// @Tags      SysUser
// @Summary   设置用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body systemReq.ChangeUserInfo true  "ID, 用户名, 昵称, 头像链接"
// @Success   200   {object}  response.Response{data=bool,msg=string}  "设置用户信息"
// @Router    /user/setUserInfo [put]
func (b *BaseApi) SetUserInfo(c *gin.Context) {
	var user systemReq.ChangeUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		global.GvaLog.Error(err.Error()+"json格式", zap.Error(err))
		response.FailWithBadRequest(err.Error()+"json格式", c)
		return
	}
	err = utils.Verify(user, utils.ChangeUserInfoVerify)
	if err != nil {
		global.GvaLog.Error(err.Error(), zap.Error(err))
		response.FailWithBadRequest(err.Error(), c)
		return
	}

	err = userService.SetUserInfo(&user)

	if err != nil {
		global.GvaLog.Error(err.Error()+"用户信息更新失败", zap.Error(err))
		response.FailWithInternalServerError(err.Error(), c)
		return
	}
	response.OkWithMessage("设置成功", c)
}

// GetUserInfo
// @Tags      SysUser
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=system.SysUser,msg=string}  "获取用户信息"
// @Router    /user/getUserInfo [get]
func (b *BaseApi) GetUserInfo(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	ReqUser, err := userService.GetUserInfo(uuid)

	if err != nil {
		global.GvaLog.Error("获取失败!", zap.Error(err))
		response.FailWithInternalServerError("获取失败", c)
		return
	}

	response.OkWithDetailed(ReqUser, "获取成功", c)
}

// WxLogin
// @Tags     Base
// @Summary  授权登录
// @Produce   application/json
// @Param    data  body      systemReq.WxLogin                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.WxLoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/wxLogin [post]
func (b *BaseApi) WxLogin(c *gin.Context) {
	var l systemReq.WxLogin

	//校验参数类型
	err := c.ShouldBindJSON(&l)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//校验必填参数
	err = utils.Verify(l, utils.WxLoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	sessionKeyDto, err := code2session(l.Code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	mobile, err := decryptPhoneData(l.EncryptedData, sessionKeyDto.SessionKey, l.Iv)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	fmt.Printf("mobile %v \n", mobile)

	response.OkWithDetailed(systemRes.WxLoginResponse{
		Token:     mobile,
		ExpiresAt: 1000,
	}, "登录成功", c)
}

// TokenNext 登录后签发 jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	// 唯一签名
	j := &utils.JWT{SigningKey: []byte(global.GvaConfig.JWT.SigningKey)}
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
		Openid:   user.Openid,
		Mobile:   user.Mobile,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		global.GvaLog.Error("获取Token失败", zap.Error(err))
		response.FailWithMessage("获取Token失败", c)
		return
	}
	if !global.GvaConfig.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)

		return
	}

	//从 redis 获取 jwtStr
	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		//redis 中没有时, 存储到redis
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GvaLog.Error("设置登录状态失败", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		//响应成功
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		//从 redis 获取失败
		global.GvaLog.Error("设置登录状态失败", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		//
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr

		if err := jwtService.JsonInBlacklist(&blackJWT); err != nil {
			response.FailWithMessage("jwt 作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// 根据 code 解析  sessionKey
func code2session(code string) (systemReq.WxSessionKeyDto, error) {
	var sessionKeyDto systemReq.WxSessionKeyDto
	httpState, bytes := utils.Get(fmt.Sprintf(global.GvaConfig.Wx.Url, global.GvaConfig.Wx.Appid, global.GvaConfig.Wx.AppSecret, code))

	if httpState != 200 {
		return sessionKeyDto, errors.New("获取 sessionKey 失败, http code: " + string(httpState))
	}

	err := json.Unmarshal(bytes, &sessionKeyDto)
	if err != nil {
		return sessionKeyDto, errors.New("json 解析失败")
	}

	return sessionKeyDto, nil
}

// 解密用户的加密信息 获取手机号码
func decryptPhoneData(phoneData, sessionKey, iv string) (string, error) {
	decrypt, err := utils.AesDecrypt(phoneData, sessionKey, iv)
	if err != nil {
		global.GvaLog.Error("解密数据失败", zap.Error(err))
		return "", err
	}
	var phoneDto = systemReq.WxPhoneDto{}
	err = json.Unmarshal([]byte(decrypt), &phoneDto)
	if err != nil {
		global.GvaLog.Error("解析手机号信息失败", zap.Error(err))
		return "", err
	}
	var phone = phoneDto.PurePhoneNumber
	return phone, nil
}
