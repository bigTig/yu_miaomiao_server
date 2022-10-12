package utils

var (
	LoginVerify   = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "Captcha": {NotEmpty()}, "CaptchaId": {NotEmpty()}}
	WxLoginVerify = Rules{"Code": {NotEmpty()}, "Iv": {NotEmpty()}, "AppId": {NotEmpty()}, "EncryptedData": {NotEmpty()}}
)
