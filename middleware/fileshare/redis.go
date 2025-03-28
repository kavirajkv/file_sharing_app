package fileshare

import (
	"github.com/redis/go-redis/v9"
	"os"
)


var RedisClient *redis.Client

func InitRedis(){
	RedisClient=redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
		Password: "",
		DB: 0,
	})
}