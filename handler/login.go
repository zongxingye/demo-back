package handler

import (
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
	"xingyeblog/db"
	myjwt "xingyeblog/midware/jwt"
	"xingyeblog/model"
	"xingyeblog/utils/redis"
)

//注册信息
type RegistInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	CreateAt int    `json:"create_at"`
	UpdateAt int    `json:"update_at"`
}

// 登陆信息
type LoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// 登陆后的结果
type LoginResult struct {
	Token string `json:"token"`
	model.User
}

// 协程池子
var wg sync.WaitGroup

// 注册
func Regist(c *gin.Context) {
	var registInfo RegistInfo
	err := c.BindJSON(&registInfo)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "注册失败" + err.Error(),
		})
		return
	}
	// 没问题则插入mysql，此处加入入库方法
	err = db.Register(registInfo.UserName, registInfo.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "注册失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "注册成功",
	})
	return
}

// login
func Login(c *gin.Context) {
	var loginReq LoginInfo
	if c.BindJSON(&loginReq) != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "json解析失败",
		})
		return
	}
	isPass, user, err := db.LoginCheck(loginReq.UserName, loginReq.Password)
	fmt.Println(user)
	if !isPass {
		c.JSON(200, gin.H{
			"status":  -1,
			"message": "验证失败",
			"err":     err,
		})
		return
	}
	generateToken(c, user)
}

// 生成token(根据model的user结构生成jwttoken)
func generateToken(c *gin.Context, user model.User) {
	var expireTime = 60 * 60 * time.Second // 过期时间60s

	j := myjwt.JWT{[]byte("newtrekWang1")}
	claims := myjwt.CustomClaims{
		user.Id,
		user.UserName,
		user.Pwd,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 60*60), // 过期时间60秒
			Issuer:    "newtrekWang",
		},
	}
	fmt.Println("ididid", claims.ID)
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("token is :", token)
	// 此处携程，对生成的token插入redis
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := redis.RedisOP{}
		uid := strconv.FormatInt(user.Id, 10)
		err = r.Set(uid, token, expireTime)
		if err != nil {
			fmt.Println(err)

		}
	}()

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})
	wg.Wait()
	return
}
func Logout(c *gin.Context) {
	// 销毁redis指定的uid
	r := redis.RedisOP{}
	token := c.GetHeader("token")
	j := myjwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token) // 此处拿到了token解析的user信息
	if err != nil {
		if err == myjwt.TokenExpired {
			err := r.ClearKey(strconv.FormatInt(claims.ID, 10))
			if err != nil {
				fmt.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "授权已过期,logout成功",
			})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		c.Abort()
		return
	}
	fmt.Println(claims.ID)
	err = r.ClearKey(strconv.FormatInt(claims.ID, 10))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "登出失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登出成功",
	})
}
