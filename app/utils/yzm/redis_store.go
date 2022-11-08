package yzm

import (
	"errors"
	"novel/app/utils/redis"
)

type RedisStore struct {
}

const TTL = 10 * 60

func (r RedisStore) Set(id string, value string) error {
	row := redis.SetEx(redis.ImageCaptchaPrefix+id, TTL, value)
	if !row {
		return errors.New("验证码存入缓存失败")
	}
	return nil
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (r RedisStore) Get(id string, clear bool) string {
	value := redis.Get(redis.ImageCaptchaPrefix + id)
	if clear {
		redis.Del(id)
	}
	return value
}

//Verify captcha's answer directly
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	value := r.Get(id, clear)
	return value == answer
}
