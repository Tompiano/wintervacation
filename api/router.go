package api

import (
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	r.Use(TokenMiddleWare())
	user := r.Group("/user")
	{
		user.POST("/register", Register) //注册
		user.POST("/add", Person)        //添加个人信息
		user.POST("/avatar", Avatar)     //添加用户头像
		user.GET("/login", Login)        //登录
		user.PUT("/forget", Forget)      //忘记密码
	}

	item := r.Group("/product") //某一种类的商品详情页
	{
		item.POST("/creator", Creator) //创建商品(店家用于添加商品)
		item.GET("/show", Show)        //展示商品
		item.POST("/explore", Explore) //搜索商品
	}

	shoppingCart := r.Group("/shoppingCart") //购物车
	{
		shoppingCart.GET("/add") //加入购物车
	}

	shop := r.Group("")
	{
		shop.POST("") //店铺公告
		shop.GET("")  //商品展示
	}

	orders := r.Group("/orders")
	{
		orders.GET("")
	}

	r.Run()
}
