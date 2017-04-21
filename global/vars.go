package global

import (
	"ckcaptcha/config"
	"gopkg.in/redis.v5"
)

var Config config.Config

var RedisClient *redis.Client