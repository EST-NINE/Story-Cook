package cache

import (
	"SparkForge/config"
	"SparkForge/pkg/util"
	"github.com/go-redis/redis"
	"strconv"
)

var RedisClient *redis.Client

func InitRedis() {
	redisDb, err := strconv.ParseInt(config.RedisDbName, 10, 64)
	if err != nil {
		util.LogrusObj.Error(err)
		return
	}
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       int(redisDb),
	})
	_, err = client.Ping().Result()
	if err != nil {
		util.LogrusObj.Infoln(err)
		panic(err)
	}
	RedisClient = client
}
