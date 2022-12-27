package controller

import (
	"douyin/api/code"
	"douyin/model"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FeedVideoHandler(c *gin.Context) {
	req := model.FeedVideoReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		zap.L().Error("FeedVideoHandler with invalid param", zap.Error(err))
		code.Response(c, err, nil)
		return
	}
	data, err := service.FeedVideo(c, req)
	if err != nil {
		zap.L().Error("service.FeedVideo failed", zap.Error(err))
	}
	code.Response(c, err, data)
}
