package model

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string //用户名
	PassWord      string //密码
	Phone         string //手机号
	Email         string //邮箱
	Identity      string //身份唯一标识
	ClientIp      string //用户ip
	ClientProt    string //用户端口
	LogInTime     uint64 //登录时间
	HeartbeatTime uint64 //心跳时间
	LogOutTime    uint64 //下线时间
	IsLogOut      bool   //是否下线
	DeviceInfo    string //设备信息
}
