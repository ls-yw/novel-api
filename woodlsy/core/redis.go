package core

import (
	"fmt"
	"github.com/go-redis/redis"
	"novel/utils/log"
	"novel/woodlsy"
)

func init() {
	woodlsy.Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", woodlsy.Configs.Redis.Host, woodlsy.Configs.Redis.Port),
		Password: woodlsy.Configs.Redis.Password, // no password set
		DB:       woodlsy.Configs.Redis.Db,       // use default DB
	})

	if err := woodlsy.Redis.Ping().Err(); err != nil { //心跳测试
		log.Logger.Error("redis连接失败", err, woodlsy.Configs.Redis)
		errMsg := "failed to init redis"
		panic(errMsg)
	}
}
