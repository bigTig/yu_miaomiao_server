package response

import "yuyu/model/system"

type LoginResponse struct {
	User      system.SysUser `json:"user"`      // 用户信息
	Token     string         `json:"token"`     // 登录凭证
	OpenId    string         `json:"openId"`    // openid
	ExpiresAt int64          `json:"expiresAt"` // 凭证过期时间
}

type WxLoginResponse struct {
	User      system.SysUser `json:"user"`      // 用户信息
	Token     string         `json:"token"`     // 登录凭证
	OpenId    string         `json:"openId"`    // openid
	ExpiresAt int64          `json:"expiresAt"` // 凭证过期时间
}
