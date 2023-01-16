package api

import "github.com/gin-gonic/gin"

func Entrance() {
	r := gin.Default()
	//r.Use(handleWare())
	u := r.Group("/user")
	{
		u.POST("/register", Register)             //注册
		u.GET("/login", TokenMiddleWare(), Login) //登录
		u.PUT("/forget", Forget)                  //忘记密码
	}
	i := r.Group("/item")
	{
		i.GET("/search", Search) //搜索商品
	}
	r.Run()
}
