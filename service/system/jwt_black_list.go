package system

import (
	"context"
	"go.uber.org/zap"
	"yuyu/global"
	"yuyu/model/system"
	"yuyu/utils"
)

type JwtService struct{}

//@author: kaifengli
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (j *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GvaDb.Create(&jwtList).Error
	if err != nil {
		return
	}

	global.BlackCache.SetDefault(jwtList.Jwt, struct {
	}{})

	return
}

//@author: kaifengli
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (j *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

//@author: kaifengli
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: redisJWT string, err error

func (j *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GvaRedis.Get(context.Background(), userName).Result()
	return redisJWT, err
}

//@author: kaifengli
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (j *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	dr, err := utils.ParseDuration(global.GvaConfig.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GvaRedis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GvaDb.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GvaLog.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
