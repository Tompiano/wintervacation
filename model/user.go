package model

import "github.com/dgrijalva/jwt-go"

//用户登录注册需要用到的

type User struct {
	ID       string `json:"ID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

// token要用到的

type MyStandardClaims struct {
	UserName string `json:"userName"`
	jwt.StandardClaims
}
type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

//添加用户信息相关的信息

type PersonInformation struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Gender   string `json:"Gender"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
}

//添加用户头像相关的信息

type Avatar struct {
	UserID     int    `json:"userID"`
	AvatarName string `json:"avatarName"`
	AvatarPath string `json:"avatarPath"`
}
