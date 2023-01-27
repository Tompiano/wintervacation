package api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"wintervacation/service"
	"wintervacation/util"
)

func TokenMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//从前端处获取jwt的Header
		tokenString := c.GetHeader("Authentication")
		mySignedKey := "token"
		//检验从前端传过来的header格式是否正确：不为空，开头前缀为Bearer（string.HasPrefix来检验）
		if tokenString == "" || strings.HasPrefix(tokenString, "Bearer") == false {
			util.ResponseNormalError(c, 10009, "token is wrong")
			c.Abort() //终止程序
			return
		}
		tokenString = tokenString[7:] //提取有效部分
		//解析token并判断解析是否成功和解析后的token是否有效
		token, claims, err := service.ParseToken(mySignedKey, tokenString, c)
		if err != nil || token.Valid == false {
			util.ResponseNormalError(c, 10008, "parse token failed")
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

	}
}
