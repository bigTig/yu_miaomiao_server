package request

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
