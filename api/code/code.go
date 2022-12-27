package code

import "errors"

type ResCode int64

var codeMsgMap = map[ResCode]string{}

var errorMsgMap = map[error]ResCode{}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[errorMsgMap[ServerBusy]]
	}
	return msg
}

// 连接错误和状态码
func CodeWithError(code ResCode, msg string) error {
	codeMsgMap[code] = msg
	err := errors.New(msg)
	errorMsgMap[err] = code
	return err
}

var (
	Success          = CodeWithError(0, "success")
	InvalidParam     = CodeWithError(499, "请求参数错误")
	InternalError    = CodeWithError(500, "internal error")
	UserNameExist    = CodeWithError(1001, "用户名已存在")
	UserNameNotExist = CodeWithError(1002, "用户名不存在")
	InvalidPassword  = CodeWithError(1003, "用户名或密码错误")
	ServerBusy       = CodeWithError(1004, "服务繁忙")
	NeedLogin        = CodeWithError(1005, "需要登陆")
	InvalidToken     = CodeWithError(1006, "token错误")
	OverdueToken     = CodeWithError(1007, "token过期")
	UserNotExist     = CodeWithError(1008, "该用户不存在")
)
