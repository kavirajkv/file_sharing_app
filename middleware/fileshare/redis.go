package fileshare

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)


var RedisClient *redis.Client

func InitRedis(){
	addr:=fmt.Sprintf("%v:6379",os.Getenv("REDIS_URL"))
	RedisClient=redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})
}