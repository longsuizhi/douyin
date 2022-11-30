package http

import (
	"github.com/gin-gonic/gin"
)

func initRouter(c *gin.Engine) {
	c.Static("/static", "./public")

	external := c.Group("douyin")
	{
		//视频流
		external.GET("/feed", feedVideo)
		//用户注册
		//external.POST("/user/register", userRegister)
		//用户登录
		//external.POST("/user/login", userLogin)
		//获取用户信息
		//external.GET("/user", getUserInfo)
		//视频投稿
		//external.POST("/publish/action", videoUpload)
		//获取用户发布列表
		//external.GET("/publish/list", getUploadHistory)

	}

	/*background := s.GROUP("/api/background")
	{

	}*/
}
