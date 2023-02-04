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
	e := model.Collection{}
	err := c.ShouldBind(&e)
	if err != nil {
		util.ResponseParaError(c)
		return
	}
	err = service.JoinCollection(model.Collection{
		UserID:        e.UserID,
		ProductID:     e.ProductID,
		ProductName:   e.ProductName,
		Kind:          e.Kind,
		Title:         e.Title,
		Info:          e.Info,
		ImagePath:     e.ImagePath,
		Price:         e.Price,
		DiscountPrice: e.DiscountPrice,
		Sales:         e.Sales,
		ShopID:        e.ShopID,
		Score:         e.Score,
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
	err, e := service.LookCollections(userID)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseCollection(c, e)
}
