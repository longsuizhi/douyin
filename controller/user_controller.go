package controller

import (
	"douyin/api/code"
	"douyin/model"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRegisterHandler(c *gin.Context) {
	req := model.UserRegisterReq{}
	if err := c.ShouldBindJSON(&req); err != nil || req.UserName == "" || req.PassWord == "" {
		zap.L().Error("UserRegisterHandler with invalid param", zap.Error(err))
		code.Response(c, err, nil)
		return
	}
	data, err := service.UserRegister(req)
	if err != nil {
		zap.L().Error("service.UserRegister failed", zap.Error(err))
	}
	code.Response(c, err, data)
}
