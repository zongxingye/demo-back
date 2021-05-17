package redis

import (
	"fmt"
	"time"
)

type RedisOP struct {
}

//var ctx = context.Background()

func (RedisOP) Set(key string, data interface{}, time time.Duration) error {
	Redisdb.Set(ctx, key, data, time)
	//if err != nil {
	//	return err
	//}
	return nil
}
func (RedisOP) Get(key string) (string, error) {
	val, err := Redisdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return val, nil
}
func (RedisOP) ClearKey(key string) error {
	err := Redisdb.Del(ctx, key).Err()
	if err != nil {

		return err
	}
	return nil
}
