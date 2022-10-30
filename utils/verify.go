package utils

var (
	IdVerify              = Rules{"ID": []string{NotEmpty()}}
	InsertAdvertVerify    = Rules{"Name": {NotEmpty()}, "Photo": {NotEmpty()}, "Position": {NotEmpty()}, "Sort": {Ge("0")}, "Type": {NotEmpty()}, "Status": {NotEmpty()}}
	UpdateAdvertVerify    = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Photo": {NotEmpty()}, "Position": {NotEmpty()}, "Sort": {NotEmpty()}, "Type": {NotEmpty()}, "Status": {NotEmpty()}}
	LoginVerify           = Rules{"Username": {NotEmpty()}, "Password": {NotEmpty()}, "Captcha": {NotEmpty()}, "CaptchaId": {NotEmpty()}}
	ChangePasswordVerify  = Rules{"Password": {NotEmpty()}, "ConfirmPassword": {NotEmpty()}}
	WxLoginVerify         = Rules{"Code": {NotEmpty()}, "Iv": {NotEmpty()}, "AppId": {NotEmpty()}, "EncryptedData": {NotEmpty()}}
	InsertCateVerify      = Rules{"Name": {NotEmpty()}, "Icon": {NotEmpty()}, "Content": {NotEmpty()}, "Sort": {Ge("0")}, "Status": {NotEmpty()}, "Remarks": {NotEmpty()}}
	UpdateCateVerify      = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Icon": {NotEmpty()}, "Content": {NotEmpty()}, "Sort": {Ge("0")}, "Status": {NotEmpty()}, "Remarks": {NotEmpty()}}
	InsertBrandVerify     = Rules{"Name": {NotEmpty()}, "Icon": {NotEmpty()}, "CateId": {Ge("0")}, "ShopId": {Ge("0")}, "BrandPrice": {Ge("0")}, "Sort": {Ge("0")}, "Status": {NotEmpty()}}
	UpdateBrandVerify     = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Icon": {NotEmpty()}, "CateId": {Ge("0")}, "ShopId": {Ge("0")}, "BrandPrice": {Ge("0")}, "Sort": {Ge("0")}, "Status": {NotEmpty()}}
	InsertHealthNewVerify = Rules{"Title": {NotEmpty()}, "Author": {NotEmpty()}, "Status": {NotEmpty()}, "Cover": {NotEmpty()}, "Content": {NotEmpty()}}
	UpdateHealthNewVerify = Rules{"Id": {NotEmpty()}, "Title": {NotEmpty()}, "Author": {NotEmpty()}, "Status": {NotEmpty()}, "Cover": {NotEmpty()}, "Content": {NotEmpty()}}
	InsertFastCateVerify  = Rules{"Name": {NotEmpty()}, "Status": {NotEmpty()}}
	UpdateFastCateVerify  = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Status": {NotEmpty()}}
	FastListVerify        = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "CateId": {Ge("0")}, "CarDog": {Ge("0")}}
	InsertFastVerify      = Rules{"Name": {NotEmpty()}, "Status": {NotEmpty()}, "CateId": {Ge("0")}, "CarDog": {Ge("0")}, "CanEat": {Ge("0")}, "Content": {NotEmpty()}, "Icon": {NotEmpty()}}
	UpdateFastVerify      = Rules{"Id": {NotEmpty()}, "Name": {NotEmpty()}, "Status": {NotEmpty()}, "CateId": {Ge("0")}, "CarDog": {Ge("0")}, "CanEat": {Ge("0")}, "Content": {NotEmpty()}, "Icon": {NotEmpty()}}
	PageInfoVerify        = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
