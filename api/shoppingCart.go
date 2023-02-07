package api

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"strconv"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func Add(c *gin.Context) {
	//获取要加入的商品的数据
	userID, _ := strconv.Atoi(c.PostForm("userID"))       //用户ID
	productID, _ := strconv.Atoi(c.PostForm("productID")) //商品ID
	amount, _ := strconv.Atoi(c.PostForm("amount"))       //商品数量
	if userID == 0 || productID == 0 || amount == 0 {
		util.ResponseParaError(c)
		return
	}
	//将要加入购物车的信息用结构体封装
	information := model.ShoppingCart{
		UserID:    userID,
		ProductID: productID,
		Amount:    amount,
		Check:     1,
	}
	//判断是否登录，若未登录则用cookie储存
	if userID == 0 {

		//判断是否存在cookie
		cookie, err := c.Cookie("cart_cookie")
		//如果没有设置过cookie则设置cookie
		//设置cookie的值为productID+amount
		if err != nil { //没有cookie
			//将信息的结构体转化为string类型然后储存进cookie里面
			str, _ := json.Marshal(information)
			cartCookie := string(str)

			//设置cookie的key，cookie的值，过期时间，所在目录，域名，是否只能通过http访问，是否允许别人通过js获取自己的cookie
			c.SetCookie("cart_cookie", cartCookie, 30, "/", "127.0.0.1", false, true)
		} else { //有cookie
			//先将cookie中的信息解析出来
			str := []byte(cookie)
			var cart []model.ShoppingCart
			err = json.Unmarshal(str, &cart)
			if err != nil {
				util.ResponseNormalError(c, 20004, "unmarshal fail")
				return
			}

			//如果商品有重复的，就加数量
			Judge := false
			for _, value := range cart {
				if value.ProductID == productID {
					value.Amount += amount
					value.Check = 1 //这里改变了商品的数量即视为勾选商品
					Judge = true
				}
			}
			if Judge == false {
				//如果没有重复商品，就直接append添加
				cart = append(cart, information)
			}
			str, _ = json.Marshal(cart)
			cartCookie := string(str)

			//这里重新建立了相同名字的cookie用于覆盖原来的cookie
			c.SetCookie("cart_cookie", cartCookie, 30, "/", "127.0.0.1", false, true)

		}
	} else {
		//如果登陆了就利用数据库来储存购物车
		//先判断数据库中是否存在相同的商品
		err, s := service.SearchProductsInCart(productID)
		if productID == productID {
			//有相同商品时候就将数量增加,则现有数量=原先数量+再次准备购买的数量
			//无论是否勾选均视为勾选-->直接在dao层操作
			amount += s.Amount
			err = service.ChangeAmount(amount, productID)
			if err != nil {
				util.ResponseInternalError(c)
				return
			}

		} else {
			//没有相同商品时候就插入数据到数据库中
			err = service.AddProductsInCart(model.ShoppingCart{
				UserID:    userID,
				ProductID: productID,
				Amount:    amount,
				Check:     1,
			})
			if err != nil {
				util.ResponseInternalError(c)
				return
			}
		}
	}

	util.ResponseOK(c)
}

func Delete(c *gin.Context) {
	judge := c.PostForm("judge")                          //判断是否是全部删除
	productID, _ := strconv.Atoi(c.PostForm("productID")) //获取要删除商品的ID
	if judge == "" || productID == 0 {
		util.ResponseParaError(c)
		return
	}
	//当判断为全部删除时不加where直接删除,否则加上限制条件再删除
	if judge == "all" {
		err := service.DeleteAllProducts()
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
	} else {
		err := service.DeleteSomeProducts(productID)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
	}
	util.ResponseOK(c)
}

func Pay(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	if userID == 0 {
		util.ResponseParaError(c)
		return
	}
	//查询数据库里userID对应的购物车中的勾选商品
	err, cart := service.ListCheck(userID)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	total := 0 //总金额
	var returnCart []model.ShoppingCart
	for _, value := range cart {
		//查询数据库中对应的商品库存是否充足
		err, judge, s := service.SearchProductIfEnough(value.ProductID, value.Amount)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
		if judge == false { //若库存不充足，就返回对应的商品不充足
			ID := strconv.Itoa(value.ProductID)
			util.ResponseNormalError(c, 20002, ID+"is not enough")
			return
		} else { //若库存充足，就更改数据库中的商品存货，返回要购买商品总金额等的信息
			number := s.Number - value.Amount
			err = service.ChangeProductsNumber(value.ProductID, number)
			if err != nil {
				util.ResponseInternalError(c)
				return
			}

			returnCart = append(returnCart, model.ShoppingCart{
				CartID:    value.CartID,
				UserID:    value.UserID,
				ProductID: value.ProductID,
				Amount:    value.Amount,
				Check:     value.Check,
			})
			//计算总金额
			total += s.Price * value.Amount
		}
	}
	util.ResponseCartPay(c, total, returnCart) //返回信息

}
func Change(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	productID, _ := strconv.Atoi(c.PostForm("productID"))
	check, _ := strconv.Atoi(c.PostForm("check"))
	if userID == 0 || productID == 0 || check == 0 {
		util.ResponseParaError(c)
		return
	}
	err := service.ChangeCheck(userID, productID, check)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)

}
