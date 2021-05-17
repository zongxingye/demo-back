package cmd

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"xingyeblog/handler"
	"xingyeblog/midware/jwt"
	"xingyeblog/midware/zaplogger"
)

func xingyeblog() {
	router := gin.New()
	//router.Use(logger.LoggerToFile())
	logger, _ := zap.NewProduction()
	router.Use(zaplogger.Ginzap(logger, time.RFC3339, true), zaplogger.RecoveryWithZap(logger, true))
	router.GET("/test", handler.Test)      // 测试路由
	router.POST("/login", handler.Login)   // 登陆
	router.POST("/regist", handler.Regist) // 注册
	blogRouter := router.Group("/api/v1")
	//blogRouter.Use(logger.LoggerToFile())
	blogRouter.Use(jwt.JWTAuth())
	{
		blogRouter.GET("/version", handler.GetVersion)
		blogRouter.POST("/logout", handler.Logout)
		blogRouter.GET("/all_articles", handler.GetAllArticles)
		blogRouter.GET("/article/:id", handler.GetArticleById)
		blogRouter.POST("/articles", handler.CreateArticle)
		blogRouter.PUT("/articles/:id", handler.UpdateArticle)
		blogRouter.DELETE("/articles/:id", handler.DeleteArticle)
	}
	router.Run(":9999")

}
