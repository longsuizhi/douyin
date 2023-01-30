package controller

import (
	"douyin/api/code"
	"douyin/model"
	"douyin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func UserRegisterHandler(c *gin.Context) {
	req := model.UserRegisterLoginReq{}
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

func UserLoginHandler(c *gin.Context) {
	req := model.UserRegisterLoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil || req.UserName == "" || req.PassWord == "" {
		zap.L().Error("UserRegisterHandler with invalid param", zap.Error(err))
		code.Response(c, err, nil)
		return
	}
	data, err := service.UserLogin(req)
	if err != nil {
		zap.L().Error("service.UserLoginHandler failed", zap.Error(err))
	}
	code.Response(c, err, data)
}

func GetUserInfoHandler(c *gin.Context) {
	//获取token中解析出来的user_id
	userID := c.Query("user_id")
	rawID, ok := c.Get("user_id")
	myID, ok := rawID.(uint)
	if !ok {
		zap.L().Error("解析userId出错")
		code.Response(c, nil, nil)
		return
	}
	if userID == "" || myID <= 0 {
		zap.L().Error("UserRegisterHandler with invalid param", zap.Error(code.InvalidParam))
		code.Response(c, code.InvalidParam, nil)
		return
	}
	data, err := service.GetUserInfo(userID, myID)
	if err != nil {
		zap.L().Error("service.UserLoginHandler failed", zap.Error(err))
	}
	code.Response(c, err, data)
}
