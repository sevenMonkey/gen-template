package app

import (
	"github.com/go-redis/redis/v7"
)

var Redis *redis.Client
//redis.Options{
//Addr:     "localhost:6379",
//Password: "", // no password set
//DB:       0,  // use default DB
//}
func InitRedis(options redis.Options)  {
	Redis = redis.NewClient(&options)
	_, err := Redis.Ping().Result()
	if err != nil{
		panic(err)
	}
}
