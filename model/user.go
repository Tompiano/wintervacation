package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID                int    `json:"ID"`
	UserName          string `json:"userName"`
	Password          string `json:"password"`
	Phone             string `json:"phone"`
	PersonInformation string `json:"personInformation"`
}

// token要用到的

type MyStandardClaims struct {
	UserName string `json:"userName"`
	Foo      string `json:"foo"`
	jwt.StandardClaims
}
