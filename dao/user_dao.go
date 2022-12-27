package dao

import "time"

type Users struct {
	Model
	UserName      string    `json:"user_name"`      //用户名
	PassWord      string    `json:"password"`       //用户密码
	RegisterAt    time.Time `json:"register_at"`    //注册时间
	LastLogin     time.Time `json:"last_login"`     //最后一次登录时间
	LoginNum      int       `json:"login_num"`      //登录次数
	FollowCount   int       `json:"follow_count"`   //关注数量
	FollowerCount int       `json:"follower_count"` //粉丝数量
	IsAdmin       int8      `json:"is_admin"`       //是否是管理员 0否  1是
}

type UsersRes struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}
