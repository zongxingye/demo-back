package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"xingyeblog/utils/redis"
)

func Test(c *gin.Context) {
	r := redis.RedisOP{}
	err := r.Set("testxingye3", "xingye", 60*time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"status": 1111,
	})
	//panic("sss")
}
func GetVersion(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "0.0.1",
	})
}
