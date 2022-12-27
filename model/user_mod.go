package model

type UserRegisterReq struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}
