package code

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
{
	"code": 1001, // 程序中的错误码
	"message": ""， // 提示信息
	"data": {},    // 数据
}
*/

type ResponseData struct {
	Code ResCode `json:"code"`
	Msg  any     `json:"msg"`
	Data any     `json:"data,omitempty"`
}

func Response(c *gin.Context, err error, data interface{}) {
	var code ResCode
	if err == nil {
		code = 0
	} else {
		code = errorMsgMap[err]
	}
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}

/*func ResponseError(c *gin.Context, code ResCode, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: data,
	})
}*/

/*func ResponseErrorWithMsg(c *gin.Context, code ResCode, msg any, data interface{}) {
	c.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}*/
