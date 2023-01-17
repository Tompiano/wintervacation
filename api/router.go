package api

import "github.com/gin-gonic/gin"

func Entrance() {
	r := gin.Default()
	//r.Use(handleWare())
	user := r.Group("/user")
	{
		user.POST("/register", Register)             //注册
		user.POST("/add", Add)                       //添加个人信息
		user.GET("/login", TokenMiddleWare(), Login) //登录
		user.PUT("/forget", Forget)                  //忘记密码
	}
	item := r.Group("/item") //商品
	{
		item.GET("/search", Search) //搜索商品
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
