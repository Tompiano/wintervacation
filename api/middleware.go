package api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"wintervacation/service"
	"wintervacation/util"
)

//jwt中间件鉴权

func TokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这里客户端在请求头携带token,并使用Bearer作为开头
		//从前端处获取jwt的Header
		tokenString := c.Request.Header.Get("Authentication")
		mySignedKey := "token"
		//检验token是否为空
		if tokenString == "" {
			util.ResponseNormalError(c, 10008, "token is empty")
			c.Abort()
			return
		}
		//检验从前端传过来的header格式是否正确：不为空，开头前缀为Bearer
		//（string.HasPrefix提取前6位数来检验）tokenString = tokenString[6:]
		//或者采用分割的方式
		parts := strings.SplitN(tokenString, " ", 2) //分成2个部分
		if len(parts) != 2 || parts[0] != "Bearer" {
			util.ResponseNormalError(c, 10009, "token's form is wrong")
			c.Abort() //终止程序
			return
		}

		//解析token并判断解析是否成功和解析后的token是否有效
		token, claims, err := service.ParseToken(mySignedKey, tokenString, c)
		if err != nil || token.Valid == false {
			util.ResponseNormalError(c, 20000, "parse token failed")
			c.Abort() //终止程序
			return
		}
		//token通过验证后获取claims中的userID
		userID := claims.Id
		//查询数据库
		u, err := service.SearchUserID(userID)
		//验证用户是否存在
		if err != nil {
			util.ResponseNormalError(c, 20001, "user is not exit")
			return
		}
		//如果用户存在将用户信息写入上下文
		c.Set("user", u)
		c.Next() //进行后续函数

	}
}

//cookie的中间件认证

func CookieMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("cookie"); err != nil {
			//检验是否有key为cookie且其值为true的cookie，若有，则继续下面的函数
			if cookie == "true" {
				c.Next()
				return
			}
		}
		//否则就直接从中间件返回接下来的请求并返回相应的报文
		util.ResponseCookie(c)
		c.Abort() //不再进行接下来的函数
	}
}
