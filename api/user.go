package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func Register(c *gin.Context) {
	UserName := c.PostForm("userName")
	Password := c.PostForm("password")
	Phone := c.PostForm("phone")

	//如果这些中间有一个没有内容，说明参数请求错误
	if UserName == "" || Password == "" {
		util.ResponseParaError(c)
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
	//判断完后就添加用户的信息
	err := service.CreateUser(model.User{
		UserName: UserName,
		Password: Password,
		Phone:    Phone,
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
	claims := model.MyStandardClaims{
		UserName: UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           //在此之前不可用
			ExpiresAt: time.Now().Unix() + 60*60*2, //将jwt设置为2小时过期
			Issuer:    "Tom-Jerry",                 //发行人：Tom-Jerry
		},
	}
	claims2 := model.MyStandardClaims{
		UserName: UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),              //在此之前不可用
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //将jwt设置为7天过期
			Issuer:    "Tom-Jerry",                    //发行人：Tom-Jerry
		},
	}
	mySignedKey := []byte("token") //自定义签名
	//用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	// 使用指定的mySignedKey签名并获得完整的编码后的字符串token
	tokenString, err := token.SignedString(mySignedKey)
	refreshString, err := refresh.SignedString(mySignedKey)
	if err != nil {
		util.ResponseNormalError(c, 10006, "Token failed")
		return
	}
	//返回token和refresh_token的信息
	util.ResponseLoginOK(c, model.Token{
		Token:        tokenString,
		RefreshToken: refreshString,
	}) //返回tokenString

}
func Refresh(c *gin.Context) {

}

func Person(c *gin.Context) {
	//获取用户信息
	userName := c.PostForm("userName")
	nickName := c.PostForm("nickName")
	gender := c.PostForm("gender")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	birthday := c.PostForm("birthday")
	//入参校验
	if userName == "" || nickName == "" || gender == "" || phone == "" || email == "" || birthday == "" {
		util.ResponseParaError(c)
		return
	}
	//加入数据库
	err := service.CreatePersonInformation(model.PersonInformation{
		UserName: userName,
		NickName: nickName,
		Gender:   gender,
		Phone:    phone,
		Email:    email,
		Birthday: birthday,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
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
	u, err := service.SearchPhone(phone)
	if u.Phone == "" || u.UserName == "" || err != nil {
		util.ResponseNormalError(c, 405, "Method Not Allowed.")
		return
	}
	//若验证符合，则将新密码插入,与注册时候要求的相同
	//密码要求在6-20位
	if len(newPassword) < 6 && len(newPassword) > 20 {
		util.ResponseNormalError(c, 10001, "password error")
		return
	}
	err = service.ChangePassword(newPassword, userName)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)

}
func Avatar(c *gin.Context) {
	//用户传头像
	avatar, err := c.FormFile("avatar")
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	if userID == 0 || err != nil {
		util.ResponseParaError(c)
		return
	}
	str := strconv.Itoa(userID)
	avatarName := str + ".png"
	err = c.SaveUploadedFile(avatar, "./"+avatarName) //fileHeader和接收文件路径 "./"+avatarName
	if err != nil {
		log.Printf("upload error:%v", err)
		util.ResponseNormalError(c, 20002, "upload avatar fail")
		return
	}
	avatarPath := "./" + avatarName
	//将照片路径存入数据库
	err = service.CreatePersonAvatar(model.Avatar{
		UserID:     userID,
		AvatarName: avatarName,
		AvatarPath: avatarPath,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
}

func AddressAdd(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	address := c.PostForm("address")
	if userID == 0 || address == "" {
		util.ResponseParaError(c)
		return
	}

	err := service.CreateAddress(userID, address)
	if err != nil {
		util.ResponseOK(c)
		return
	}

}
