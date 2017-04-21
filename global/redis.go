package global

import (
	"gopkg.in/redis.v5"
	"ckcaptcha/config"
)

func NewRedisClient(redisCfg *config.RedisConfig) (redisClient *redis.Client) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		DB:       redisCfg.DB,
		Password: redisCfg.Password,
	})
	return redisClient
}

func RedisClose(redisClient *redis.Client) error {
	return redisClient.Close()
}