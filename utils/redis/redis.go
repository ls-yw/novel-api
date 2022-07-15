package redis

import (
	"novel/utils/log"
	"novel/woodlsy"
	"time"
)

func Exists(key string) bool {
	v, err := woodlsy.Redis.Exists(key).Result()
	if err != nil {
		log.Logger.Error("【redis】【Exists】key:", key, "error:", err)
		return false
	}

	return v > 0
}

func Get(key string) string {
	value, err := woodlsy.Redis.Get(key).Result()
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
	err := woodlsy.Redis.Set(key, value, time.Second*time.Duration(ttl)).Err()
	if err != nil {
		log.Logger.Error("【redis】【SetEx】key:", key, "value:", value, "ttl:", ttl, "error:", err)
		return false
	}
	return true
}

func Del(key string) bool {
	err := woodlsy.Redis.Del(key).Err()
	if err != nil {
		log.Logger.Error("【redis】【Del】key:", key, "error:", err)
		return false
	}
	return true
}

func Close() error {
	return woodlsy.Redis.Close()
}

func Expire(key string, ttl int) bool {
	err := woodlsy.Redis.Expire(key, time.Second*time.Duration(ttl)).Err()
	if err != nil {
		log.Logger.Error("【redis】【Expire】key:", key, "ttl:", ttl, "error:", err)
		return false
	}
	return true
}
