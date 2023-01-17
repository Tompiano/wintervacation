package model

import "github.com/dgrijalva/jwt-go"

//用户登录注册需要用到的

type User struct {
	ID       int    `json:"ID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Token    string `json:""`
}

// token要用到的

type MyStandardClaims struct {
	UserName string `json:"userName"`
	Foo      string `json:"foo"`
	jwt.StandardClaims
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
