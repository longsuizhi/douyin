package dao

import (
	"douyin/define"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Users struct {
	Model
	UserName   string    `json:"user_name"`   //用户名
	Password   string    `json:"password"`    //用户密码
	RegisterAt time.Time `json:"register_at"` //注册时间
	LastLogin  time.Time `json:"last_login"`  //最后一次登录时间
	LoginNum   int       `json:"login_num"`   //登录次数
	IsAdmin    int8      `json:"is_admin"`    //是否是管理员 0否  1是
}

type UsersRes struct {
	UserID uint   `json:"user_id"`
	Token  string `json:"token"`
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

func userFollowKey(userID uint) string {
	return fmt.Sprintf("%s%d", define.UserFollowNum, userID)
}

func GetUserFollowInfo(userID uint) (string, error) {
	key := userFollowKey(userID)

}
