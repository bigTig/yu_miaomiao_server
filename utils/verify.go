package utils

var (
	IdVerify           = Rules{"ID": []string{NotEmpty()}}
	InsertAdvertVerify = Rules{"Name": {NotEmpty()}, "Photo": {NotEmpty()}, "Position": {NotEmpty()}, "Sort": {NotEmpty()}, "Type": {NotEmpty()}, "Status": {NotEmpty()}}
	UpdateAdvertVerify = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Photo": {NotEmpty()}, "Position": {NotEmpty()}, "Sort": {NotEmpty()}, "Type": {NotEmpty()}, "Status": {NotEmpty()}}
	LoginVerify        = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "Captcha": {NotEmpty()}, "CaptchaId": {NotEmpty()}}
	WxLoginVerify      = Rules{"Code": {NotEmpty()}, "Iv": {NotEmpty()}, "AppId": {NotEmpty()}, "EncryptedData": {NotEmpty()}}
	PageInfoVerify     = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
