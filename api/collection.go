package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

//加入收藏

func Join(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	productID, _ := strconv.Atoi(c.PostForm("productID"))
	if userID == 0 || productID == 0 {
		util.ResponseParaError(c)
		return
	}
	//将userID和productID插入到数据库中
	err := service.JoinCollection(model.Collection{
		UserID:    userID,
		ProductID: productID,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)

}

//取消收藏

func DeleteCollection(c *gin.Context) {
	ProductID, _ := strconv.Atoi(c.PostForm("productID"))
	if ProductID == 0 {
		util.ResponseInternalError(c)
		return
	}
	err := service.DeleteCollectionByID(ProductID)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
}

//查看收藏列表

func LookCollection(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	if userID == 0 {
		util.ResponseParaError(c)
		return
	}
	//获取该用户收藏夹中的所有productID
	err, collections := service.LookCollections(userID)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	var details []*model.Product //用map来装所有的product的详细信息
	for _, products := range collections {
		//利用收藏夹中的productID找到商品的全部信息，将商品的全部信息返回
		err, productDetails := service.SearchProducts(products.ProductID)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}

		details = append(details, productDetails...)

	}
	//返回该用户下的所有收藏的商品的信息
	util.ResponseCollection(c, details)
}
