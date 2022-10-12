package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"yuyu/global"
)

func Redis() {
	redisCfg := global.GvaConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GvaLog.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.GvaLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.GvaRedis = client
	}
}
