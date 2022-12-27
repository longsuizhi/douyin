package service

import (
	"douyin/api/code"
	"douyin/dao"
	"douyin/model"
)

func UserRegister(req model.UserRegisterReq) (dao.UsersRes, error) {
	//查询是否已经存在该用户名
	res := dao.UsersRes{}
	return res, code.UserExist
	//创建该用户，获取uid
	//使用创建得到的uid进行jwt
	//存储jwt
}
