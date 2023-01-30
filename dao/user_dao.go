package dao

import (
	"douyin/define"
	"errors"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

type Users struct {
	Model      `json:"model,omitempty"`
	UserName   string    `json:"user_name,omitempty"`   //用户名
	Password   string    `json:"password,omitempty"`    //用户密码
	RegisterAt time.Time `json:"register_at,omitempty"` //注册时间
	LastLogin  time.Time `json:"last_login,omitempty"`  //最后一次登录时间
	LoginNum   int       `json:"login_num,omitempty"`   //登录次数
	IsAdmin    int8      `json:"is_admin,omitempty"`    //是否是管理员 0否  1是
}

type UsersRes struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfo struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserWithName struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
}

func GetUserData(userName string) (Users, error) {
	res := Users{}
	err := SvDB.Model(Users{}).Where("user_name = ?", userName).First(&res).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, nil
	}
	return res, err
}

func CreateUser(newUser *Users) error {
	return SvDB.Model(Users{}).Create(newUser).Error
}

func UpdateUser(newUser *Users) error {
	return SvDB.Model(Users{}).Where("id = ?", newUser.ID).Updates(newUser).Error
}

func GetUserInfo(userID interface{}) (UserWithName, error) {
	data := UserWithName{}
	err := SvDB.Table("users").Select("id, user_name").Where("id = ? and deleted_at is null", userID).First(&data).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, nil
	}
	return data, err
}

func userFollowKey(userID interface{}) string {
	return fmt.Sprintf("%s%d", define.UserFollow, userID)
}

func userFollowerKey(userID interface{}) string {
	return fmt.Sprintf("%s%d", define.UserFollower, userID)
}

func GetUserFollowInfo(userID uint, myID uint) (int64, int64, bool, error) {
	followKey := userFollowKey(userID)
	followerKey := userFollowerKey(userID)
	myFollowKey := userFollowKey(myID)
	if err := SvRedis.Send("SCARD", followKey); err != nil {
		return 0, 0, false, err
	}
	if err := SvRedis.Send("SCARD", followerKey); err != nil {
		return 0, 0, false, err
	}
	if err := SvRedis.Send("SISMEMBER", myFollowKey, userID); err != nil {
		return 0, 0, false, err
	}
	if err := SvRedis.Flush(); err != nil {
		return 0, 0, false, err
	}
	follow, err := redis.Int64(SvRedis.Receive())
	if err != nil {
		return 0, 0, false, err
	}
	follower, err := redis.Int64(SvRedis.Receive())
	if err != nil {
		return 0, 0, false, err
	}
	isFollow, err := redis.Bool(SvRedis.Receive())
	if err != nil {
		return 0, 0, false, err
	}
	return follow, follower, isFollow, nil
}
