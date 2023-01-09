package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

//注册相关

func CreateUserInformation(userName string) (u model.User) {
	u = dao.SelectUserInformation(userName)
	return u
}

func CreateUser(u model.User) (err error) {
	err = dao.InsertUser(u)
	return err
}

//登录相关
