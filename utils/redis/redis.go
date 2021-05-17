package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Redisdb *redis.Client
var ctx = context.Background()

func init() {
	Redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := Redisdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	println(pong)
}
