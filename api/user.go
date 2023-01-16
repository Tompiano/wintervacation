package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
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
		//c.SetCookie("cookie", "test", 3600, "/",
		//"localhost", false, true)
		util.ResponseOK(c)
	}

}
func Login(c *gin.Context) {

	UserName := c.PostForm("userName")
	Password := c.PostForm("password")
	//如果这些中间有一个没有内容，说明参数请求错误
	if UserName == "" || Password == "" {
		util.ResponseParaError(c)
		return
	}
	//检验用户是否存在,和密码是否正确
	u := service.CheckUsernameAndPassword(UserName)
	if UserName != u.UserName || Password != u.Password {
		util.ResponseNormalError(c, 404, "Not Found")
		return
	}
	util.ResponseOK(c)
	//生成token
	mySignedKey := []byte("token") //自定义jwt密钥
	claims := model.MyStandardClaims{
		UserName: "userName",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           //在此之前不可用
			ExpiresAt: time.Now().Unix() + 60*60*2, //将jwt设置为2小时过期
			Issuer:    "user",                      //发行人：user
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySignedKey) //用HS256对jwt签名
	if err != nil {
		util.ResponseNormalError(c, 10006, "Token failed")
		return
	}
	fmt.Println(ss)

}

func Forget(c *gin.Context) {

	userName := c.PostForm("userName")
	newPassword := c.PostForm("new_password")
	phone := c.PostForm("phone")
	//检查参数是否错误
	if userName == "" || newPassword == "" || phone == "" {
		util.ResponseParaError(c)
		return
	}
	//查询phone是否与username相符合
	u := service.SearchPhone(phone)
	if u.Phone == "" || u.UserName == "" {
		util.ResponseNormalError(c, 405, "Method Not Allowed.")
		return
	}
	//若验证符合，则将新密码插入,与注册时候要求的相同
	//密码要求在6-20位
	if len(newPassword) < 6 && len(newPassword) > 20 {
		util.ResponseNormalError(c, 10001, "password error")
		return
	}
	err := service.ChangePassword(newPassword, userName)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)

}
