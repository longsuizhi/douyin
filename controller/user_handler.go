package controller

import (
	"douyin/model"
	"douyin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userRegister(c *gin.Context) {
	req := model.UserRegisterReq{}
	if err := c.Bind(&req); err != nil || req.UserName == "" || req.PassWord == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data, err := service.UserRegister(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  -1,
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
			"data": data,
		})
	}
}
