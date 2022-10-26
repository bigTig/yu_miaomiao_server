package utils

var (
	IdVerify             = Rules{"ID": []string{NotEmpty()}}
	InsertAdvertVerify   = Rules{"Name": {NotEmpty()}, "Photo": {NotEmpty()}, "Position": {NotEmpty()}, "Sort": {NotEmpty()}, "Type": {NotEmpty()}, "Status": {NotEmpty()}}
	UpdateAdvertVerify   = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Photo": {NotEmpty()}, "Position": {NotEmpty()}, "Sort": {NotEmpty()}, "Type": {NotEmpty()}, "Status": {NotEmpty()}}
	LoginVerify          = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "Captcha": {NotEmpty()}, "CaptchaId": {NotEmpty()}}
	ChangePasswordVerify = Rules{"Password": {NotEmpty()}, "ConfirmPassword": {NotEmpty()}}
	WxLoginVerify        = Rules{"Code": {NotEmpty()}, "Iv": {NotEmpty()}, "AppId": {NotEmpty()}, "EncryptedData": {NotEmpty()}}
	InsertCateVerify     = Rules{"Name": {NotEmpty()}, "Icon": {NotEmpty()}, "Content": {NotEmpty()}, "Sort": {NotEmpty()}, "Status": {NotEmpty()}, "Remarks": {NotEmpty()}}
	PageInfoVerify       = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
