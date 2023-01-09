package api

import (
	"github.com/gin-gonic/gin"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func Register(c *gin.Context) {
	UserName := c.PostForm("userName")
	Password := c.PostForm("password")
	Phone := c.PostForm("phone")
	PersonInformation := c.PostForm("personInformation")

	//如果这些中间有一个没有内容，说明参数请求错误
	if UserName == "" || Password == "" || PersonInformation == "" {
		util.ResponseParaError(c)
		return
	}
	//添加用户信息(用户信息不能超过500字)
	if len(PersonInformation) > 500 {
		util.ResponseNormalError(c, 10005, "userInformation error")
		return
	}
	//密码要求在6-20位
	if len(Password) < 6 && len(Password) > 20 {
		util.ResponseNormalError(c, 10001, "password error")
		return
	}
	//用户名不可超过20个字符,且不可以是别人用过的
	if len(UserName) > 20 {
		util.ResponseNormalError(c, 10003, "userName error")
		return
	} else {
		u := service.CreateUserInformation(UserName)
		if u.UserName != "" {
			util.ResponseNormalError(c, 10002, "userName repeat")
			return
		}
	}
	//电话号码必须是11位
	if len(Phone) != 11 {
		util.ResponseNormalError(c, 10004, "phone error")
		return
	}
	//用户信息未在500个字符以内
	if len(PersonInformation) > 500 {
		util.ResponseNormalError(c, 10005, "personInformation error")
		return
	}
	//判断完后就添加用户的信息
	err := service.CreateUser(model.User{
		UserName:          UserName,
		Password:          Password,
		Phone:             Phone,
		PersonInformation: PersonInformation,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	} else {
		c.SetCookie("cookie", "test", 3600, "/",
			"localhost", false, true)
		util.ResponseOK(c)
	}

}
