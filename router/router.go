package router

import (
	"douyin/controller"
	"douyin/logger"
	"douyin/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
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
	external := c.Group("/douyin")
	{
		//根据接口需求灵活地考虑是否加入JWT鉴权

		//--------------------- 基础api ---------------------
		// 视频流
		//external.GET("/feed", controller.FeedVideoHandler)
		// 用户注册
		external.POST("/user/register", middlewares.SHAMiddleWare(), controller.UserRegisterHandler)
		// 用户登录
		external.POST("/user/login", middlewares.SHAMiddleWare(), controller.UserLoginHandler)
		// 获取用户信息
		external.GET("/user", middlewares.JWTMiddleWare(), controller.GetUserInfoHandler)
		// 视频投稿
		//external.POST("/publish/action", controller.VideoUpload)
		// 获取用户发布列表
		//external.GET("/publish/list", controller.GetUploadHistory)

		//--------------------- 拓展接口I ---------------------
		// 赞操作
		//external.POST("/favorite/action")
		// 点赞列表
		//external.GET("/favorite/list")
		// 评论操作
		//external.POST("/comment/action")
		// 视频评论列表
		//external.GET("/comment/list")

		//--------------------- 拓展接口II ---------------------
		// 关系操作
		//external.POST("/relation/action")
		// 用户关注列表
		//external.GET("/relation/follow/list")
		// 用户粉丝列表
		//external.GET("/relation/follower/list")
	}

	/*background := s.GROUP("/api/v1/background")
	{

	}*/
}
