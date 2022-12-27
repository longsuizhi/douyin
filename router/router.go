package router

import (
	"douyin/controller"
	"douyin/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(c *gin.Engine) {
	c.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 连通测试
	c.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})

	//静态文件服务，存储视频图片
	c.Static("/static", "./static")
	//设置统一接口
	external := c.Group("/api/v1")
	{
		//根据接口需求灵活地考虑是否加入JWT鉴权

		//--------------------- 基础api ---------------------
		// 视频流
		//external.GET("/feed", controller.FeedVideoHandler)
		// 用户注册
		external.POST("/user/register", controller.userRegister)
		// 用户登录
		//external.POST("/user/login", userLogin)
		// 获取用户信息
		//external.GET("/user", getUserInfo)
		// 视频投稿
		//external.POST("/publish/action", videoUpload)
		// 获取用户发布列表
		//external.GET("/publish/list", getUploadHistory)

	}

	/*background := s.GROUP("/api/background")
	{

	}*/
}
