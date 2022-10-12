package config

type Server struct {
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Wx  WX  `mapstructure:"wx" json:"wx" yaml:"wx"`

	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`

	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// auto
	AutoCode Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`

	Timer Timer `mapstructure:"timer" json:"timer" yaml:"timer"`
}
