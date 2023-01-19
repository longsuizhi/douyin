package service

import (
	"douyin/api/code"
	"douyin/dao"
	"douyin/define"
	"douyin/middlewares"
	"douyin/model"
	"time"

	"go.uber.org/zap"
)

func UserRegister(req model.UserRegisterLoginReq) (dao.UsersRes, error) {
	res := dao.UsersRes{}
	//校验用户名和密码长度
	err := CheckLength(req)
	if err != nil {
		zap.L().Error("CheckLength failed", zap.Error(err))
		return res, err
	}
	//查询是否已经存在该用户名
	data, err := dao.GetUserData(req.UserName)
	if err != nil {
		zap.L().Error("GetUserData failed", zap.Error(err))
		return res, err
	}
	//用户名已存在
	if data.ID > 0 {
		return res, code.UserNameExist
	}
	//创建该用户，获取uid
	newUser := dao.Users{
		UserName:   req.UserName,
		Password:   middlewares.EncodeMD5(req.PassWord),
		RegisterAt: time.Now(),
		LastLogin:  time.Now(),
	}
	createErr := dao.CreateUser(&newUser)
	if createErr != nil {
		zap.L().Error("CreateUser failed", zap.Error(createErr))
		return res, createErr
	}
	//使用创建得到的uid颁布token
	token, err := middlewares.GenToken(newUser.ID, newUser.Password)
	if err != nil {
		zap.L().Error("GenToken failed", zap.Error(err))
		return res, err
	}
	res.Token = token
	res.UserID = newUser.ID
	return res, nil
}

func CheckLength(req model.UserRegisterLoginReq) error {
	if len(req.UserName) > define.MaxUsernameLength {
		return code.UserNameTooLong
	}
	if len(req.UserName) < define.MinUsernameLength {
		return code.UserNameTooShort
	}
	if len(req.PassWord) > define.MaxPasswordLength {
		return code.PasswordTooLong
	}
	if len(req.PassWord) < define.MinPasswordLength {
		return code.PasswordTooShort
	}
	return nil
}

func UserLogin(req model.UserRegisterLoginReq) (dao.UsersRes, error) {
	res := dao.UsersRes{}
	//校验用户名和密码长度
	err := CheckLength(req)
	if err != nil {
		zap.L().Error("CheckLength failed", zap.Error(err))
		return res, err
	}
	//查询是否已经存在该用户名
	data, err := dao.GetUserData(req.UserName)
	if err != nil {
		zap.L().Error("GetUserData failed", zap.Error(err))
		return res, err
	}
	//用户不存在
	if data.ID <= 0 {
		return res, code.UserNotExist
	}
	if middlewares.EncodeMD5(req.PassWord) != data.Password {
		return res, code.InvalidPassword
	}
	//更新最后一次登录时间，登录次数加一
	data.LastLogin = time.Now()
	data.LoginNum++
	updateErr := dao.UpdateUser(&data)
	if updateErr != nil {
		zap.L().Error("UpdateUser failed", zap.Error(updateErr))
		return res, updateErr
	}
	//使用创建得到的uid颁布token
	token, err := middlewares.GenToken(data.ID, data.Password)
	if err != nil {
		zap.L().Error("GenToken failed", zap.Error(err))
		return res, err
	}
	res.Token = token
	res.UserID = data.ID
	return res, nil
}

func GetUserInfo(userID interface{})
