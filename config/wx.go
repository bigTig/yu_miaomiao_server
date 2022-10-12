package config

type WX struct {
	Appid     string `mapstructure:"appid" json:"appid" yaml:"appid"`                // AppID(小程序ID)
	AppSecret string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"` // 小程序密钥
	Url       string `mapstructure:"url" json:"url" yaml:"url"`                      // jscode2session url
}
