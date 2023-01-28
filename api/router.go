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

	item := r.Group("/product")
	{
		item.POST("/creator", Creator) //创建商品(店家用于添加商品)
		item.GET("/show", Show)        //展示商品
		item.POST("/explore", Explore) //搜索商品
	}

	shoppingCart := r.Group("/shoppingCart") //购物车
	{
		shoppingCart.GET("/add", Add)          //将商品加入购物车
		shoppingCart.DELETE("/delete", Delete) //删除购物车中的商品
		shoppingCart.PUT("/pay", Pay)          //将购物车内商品结账
	}

	comment := r.Group("comment") //商品的评论
	{
		comment.POST("/writer", Writer)       //写评论
		comment.PUT("/delete", DeleteComment) //删除评论
	}

	shop := r.Group("/shop")
	{
		shop.POST("/writer", AnnouncementWriter) //写店铺公告
		shop.PUT("/update", AnnouncementUpdate)  //更新店铺公告
		shop.GET("/show", ShowShopProducts)      //商品展示
	}

	orders := r.Group("/orders") //订单
	{
		orders.GET("")
	}

	r.Run()
}
