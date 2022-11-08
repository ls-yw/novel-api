package redis

import (
	"github.com/woodlsy/woodGin/log"
	"github.com/woodlsy/woodGin/redis"
	"time"
)

func Exists(key string) bool {
	v, err := redis.Redis.Exists(key).Result()
	if err != nil {
		log.Logger.Error("【redis】【Exists】key:", key, "error:", err)
		return false
	}

	return v > 0
}

func Get(key string) string {
	value, err := redis.Redis.Get(key).Result()
	if err != nil {
		log.Logger.Error("【redis】【Get】key:", key, "value:", value, "error:", err)
		return ""
	}
	return value
}

//
//func Set(key string, value interface{}, ttl int) bool {
//	err := woodlsy.Redis.Set(key, value, time.Second*time.Duration(ttl)).Err()
//	if err != nil {
//		log.Logger.Error("【redis】【SetEx】key:", key, "value:", value, "ttl:", ttl, "error:", err)
//		return false
//	}
//	return true
//}

func SetEx(key string, ttl int, value string) bool {
	err := redis.Redis.Set(key, value, time.Second*time.Duration(ttl)).Err()
	if err != nil {
		log.Logger.Error("【redis】【SetEx】key:", key, "value:", value, "ttl:", ttl, "error:", err)
		return false
	}
	return true
}

func Del(key string) bool {
	err := redis.Redis.Del(key).Err()
	if err != nil {
		log.Logger.Error("【redis】【Del】key:", key, "error:", err)
		return false
	}
	return true
}

func Close() error {
	return redis.Redis.Close()
}

func Expire(key string, ttl int) bool {
	err := redis.Redis.Expire(key, time.Second*time.Duration(ttl)).Err()
	if err != nil {
		log.Logger.Error("【redis】【Expire】key:", key, "ttl:", ttl, "error:", err)
		return false
	}
	return true
}
