package api

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func HTML(c *gin.Context) {
	var wr http.ResponseWriter
	var r http.Request
	//解析指定文件-->GitHub.html，生成模板对象
	temp, err := template.ParseFiles("util/GitHub.html")
	if err != nil {
		util.ResponseNormalError(c, 500, "fail to get template")
		return
	}
	//利用已知的config来渲染模板
	err = temp.Execute(wr, r)
	if err != nil {
		util.ResponseNormalError(c, 500, "fail to execute")
		return
	}
}

func Oauth(c *gin.Context) {

	//用户同意授权后，Github就会跳转到redirect_url指定的跳转网站，并带上授权码code。
	//因而从url中取出授权码code
	code := c.PostForm("code")

	//用授权码code向GitHub请求令牌token

	//1.先访问GitHub的令牌接口（一个url），且向GitHub提供三个参数client_id,client_secret,code
	TokenAuthUrl := service.GetTokenAuthUrl(code)

	//2.GitHub回应一段JSON数据，从中提取出令牌access_token
	accessToken, err := service.GetToken(TokenAuthUrl) //再根据这个url从中得到token
	if err != nil {
		util.ResponseNormalError(c, 500, "fail to get token")
		return
	}

	//通过token向api申请数据,并得到用户信息
	userInfo, err := service.GetUserInfo(accessToken)
	if err != nil {
		util.ResponseNormalError(c, 500, "fail to get user info")
		return
	}
	//fmt.Println(userInfo)
	//获取GitHub的用户名来登录此平台
	//保持登录-->token

	//生成自身的token和refresh_token
	userName := userInfo["login"].(string) //得到GitHub认证的用户名
	tokenString, refreshString, err := service.CreateTokens(userName, c)
	if err != nil {
		util.ResponseNormalError(c, 10006, "Token failed")
		return
	}
	//返回token和refresh_token的信息
	util.ResponseLoginOK(c, model.Token{
		Token:        tokenString,
		RefreshToken: refreshString,
	})

}
