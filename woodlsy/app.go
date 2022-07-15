package woodlsy

import (
	"github.com/go-redis/redis"
	"novel/woodlsy/config"
)

var Configs config.Configs
var Redis *redis.Client
