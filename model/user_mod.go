package model

type UserRegisterLoginReq struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}