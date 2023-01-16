package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"wintervacation/dao"
	"wintervacation/model"
)

//注册相关----------------------------------------------------------------------------------------------------------------

func CreateUserInformation(userName string) (u model.User) {
	u = dao.SelectUserInformation(userName)
	return u
}

func CreateUser(u model.User) (err error) {
	err = dao.InsertUser(u)
	return err
}

//登录相关---------------------------------------------------------------------------------------------------------------

func CheckUsernameAndPassword(userName string) (u model.User) {
	u = dao.SelectUserNameIfFirPassword(userName)
	return
}

//解析token

func ParseToken(mySignedKey string, tokenString string, c *gin.Context) (*jwt.Token, model.MyStandardClaims, error) {
	token, claims, err := dao.Authentication(mySignedKey, tokenString, c)
	return token, claims, err
}

//忘记密码相关------------------------------------------------------------------------------------------------------------

func SearchPhone(phone string) (u model.User) {
	u = dao.SelectPhoneIfExist(phone)
	return
}
func ChangePassword(password, userName string) (err error) {
	err = dao.UpdatePassword(password, userName)
	return
}
