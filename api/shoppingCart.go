package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func Add(c *gin.Context) {
	//获取要加入的商品的数据
	product := model.ShoppingCart{}
	err := c.ShouldBind(&product)
	if err != nil {
		util.ResponseParaError(c)
		return
	}
	//判断是否为登录状态，但是都可以加入购物车
	if product.UserID == 0 {
		temporaryID, _ := strconv.Atoi(c.PostForm("temporaryID")) //获取游客的身份ID
		err = service.AddProductsInTemporaryCart(product, temporaryID)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
	}
	//如果登陆了就利用数据库来储存购物车
	//先判断数据库中是否存在相同的商品
	err, s := service.SearchProductsInCart(product.Product.ProductID)
	if product.Product.ProductID == s.Product.ProductID {
		//有相同商品时候就将数量增加,则现有数量=原先数量+再次准备购买的数量
		product.Amount += s.Amount
		err = service.ChangeAmount(product.Amount, product.Product.ProductID)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}

	} else {
		//没有相同商品时候就插入数据到数据库中
		err = service.AddProductsInCart(product)
		if err != nil {
			util.ResponseInternalError(c)
			return
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
	userID, _ := strconv.Atoi(c.PostForm("userID"))         //用户的ID
	kindNumber, _ := strconv.Atoi(c.PostForm("kindNumber")) //有几种不同的商品
	productID := make(map[int]int, kindNumber)
	numbers := make(map[int]int, kindNumber)
	price := make(map[int]int, kindNumber)
	total := 0
	//想法是从前端获取选中商品的ID和数量
	for i := 0; i < kindNumber; i++ {
		ID, _ := strconv.Atoi(c.PostForm("productID"))  //获取被选中的商品的productID
		number, _ := strconv.Atoi(c.PostForm("number")) //获取被选中的商品的数量number
		if userID == 0 || ID == 0 || number == 0 {
			util.ResponseParaError(c)
			return
		}
		productID[i] = ID
		numbers[i] = number
		//从数据库中查询商品库存是否充足
		err, judge, p := service.SearchProductIfEnough(ID, number)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
		num := strconv.Itoa(number)
		if judge == false {
			util.ResponseNormalError(c, 30001, num+"number is not enough")
			return
		}
		total += p.Price * number   //计算总金额
		price[i] = p.Price * number //计算每一种商品的金额
	}
	util.ResponseInformation(c, userID, total)
	for i := 0; i < kindNumber; i++ {
		util.ResponsePay(c, productID[i], numbers[i], price[i])
	}

}
