package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strconv"
	"xingyeblog/db"
	myjwt "xingyeblog/midware/jwt"
)

//type ArticleInfo struct {
//	UserId   int    `json:"user_id"`
//	Content  string `json:"content"`
//	Title    string `json:"title"`
//	Spot     int    `json:"spot"`
//	CreateAt int    `json:"create_at"`
//	UpdateAt int    `json:"update_at"`
//}

// 获取所有文章的handler
func GetAllArticles(c *gin.Context) {
	token := c.GetHeader("token")
	j := myjwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token) // 此处拿到了token解析的user信息
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		c.Abort()
		return
	}
	//	根据uid去查数据库
	uid := claims.ID
	articleArr, err := db.GetAllArticlesByUid(int(uid))
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "查询失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"data":   articleArr,
	})
}

//获取单个文章的（by id）方法

func GetArticleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	article, err := db.GetOneArticleById(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "查询失败" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 0,
		"data":   article,
	})
}

// 创建文章
func CreateArticle(c *gin.Context) {

	//req := c.PostForm("article")
	var req db.ArticleInfo
	err := c.BindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "创建失败" + err.Error(),
		})
		return
	}
	token := c.GetHeader("token")
	j := myjwt.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token) // 此处拿到了token解析的user信息
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		c.Abort()
		return
	}

	uid := claims.ID
	fmt.Println(uid, req, reflect.TypeOf(req))
	err = db.CreateArticleWithUid(int(uid), req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "插入成功",
	})
}

// 更新文章内容属性
func UpdateArticle(c *gin.Context) {
	var req db.ArticleInfo
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.BindJSON(&req)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "更新失败" + err.Error(),
		})
		return
	}
	err = db.UpdateArticle(id, req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "更新失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "修改成功",
	})
}

// 删除指定文章内容)
func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "删除失败" + err.Error(),
		})
		return
	}
	err = db.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "删除失败" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "删除成功",
	})
}
