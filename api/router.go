package api

import (
	"github.com/gin-gonic/gin"
)

func Entrance() {
	r := gin.Default()
	//r.Use(TokenMiddleWare())

	user := r.Group("/user")
	{
		user.POST("/register", Register)                        //注册
		user.GET("/login", Login)                               //登录
		user.PUT("/forget", Forget)                             //忘记密码
		user.POST("/add", TokenMiddleWare(), Person)            //添加个人信息
		user.POST("/addressAdd", TokenMiddleWare(), AddressAdd) //添加用户的地址
		user.POST("/avatar", TokenMiddleWare(), Avatar)         //添加用户头像

	}

	item := r.Group("/product")
	{
		item.POST("/creator", Creator) //创建商品(店家用于添加商品)
		item.GET("/show", Show)        //展示商品
		item.POST("/explore", Explore) //搜索商品
		item.GET("/detail", Detail)    //商品详情页的展示
	}

	shoppingCart := r.Group("/shopping_cart") //购物车
	{

		shoppingCart.POST("/add", Add)         //将商品加入购物车
		shoppingCart.PUT("/change", Change)    //改变商品的勾选状态
		shoppingCart.DELETE("/delete", Delete) //删除购物车中的商品
		shoppingCart.GET("/pay", Pay)          //将购物车内商品结账
	}

	comment := r.Group("/comment") //商品的评论
	{
		comment.POST("/writer", Writer)       //写评论
		comment.PUT("/delete", DeleteComment) //删除评论
		comment.GET("/look", LookComment)     //查看评论
	}

	shop := r.Group("/shop")
	{
		shop.POST("/writer", AnnouncementWriter)   //写店铺公告
		shop.PUT("/update", AnnouncementUpdate)    //更新店铺公告
		shop.GET("/show", ShowShopProducts)        //商品展示
		shop.POST("/detail_writer", ProductDetail) //商品详情页图片的录入
		shop.PUT("/detail_update", DetailUpdate)   //商品详情页的图片更新
	}

	collection := r.Group("/collection")
	{
		collection.POST("/join", Join)                 //加入收藏夹
		collection.DELETE("/delete", DeleteCollection) //删除收藏夹中的内容
		collection.GET("/show", LookCollection)        //查看收藏夹
	}

	orders := r.Group("/orders") //订单
	{
		orders.GET("/prepare", Prepare)        //选择收货地址
		orders.GET("/commit", Commit)          //提交订单
		orders.GET("/success", OrderSuccess)   //订单已支付状态显示
		orders.GET("/complete", OrderComplete) //订单已收货状态展示
	}

	r.Run()
}
