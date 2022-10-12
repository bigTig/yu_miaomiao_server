package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"yuyu/config"
	"yuyu/utils/timer"
)

var (
	GvaLog     *zap.Logger
	GvaVp      *viper.Viper
	GvaRedis   *redis.Client
	GvaConfig  config.Server
	GvaDb      *gorm.DB
	GvaDbList  map[string]*gorm.DB
	GvaTimer   timer.Timer = timer.NewTimerTask()
	BlackCache local_cache.Cache
)
