package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wintervacation/model"
)

func ResponseParaError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": 300,
		"info":   "params error",
	})
}
func ResponseNormalError(c *gin.Context, status int, info string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"status": status,
		"info":   info,
	})
}
func ResponseInternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": 500,
		"info":   "internal error",
	})
}
func ResponseOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
	})
}
func ResponseProduct(c *gin.Context, p model.Product) {
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"kind":          p.Kind,
		"productName":   p.ProductName,
		"title":         p.Title,
		"info":          p.Info,
		"imagePath":     p.ImagePath,
		"price":         p.Price,
		"discountPrice": p.DiscountPrice,
		"onSale":        p.Sales,
		"score":         p.Score,
	})
}
func ResponseDetail(c *gin.Context, d model.ProductDetail) {
	c.JSON(http.StatusOK, gin.H{
		"status":      200,
		"productName": d.ProductName,
		"DetailPath":  d.Detail,
	})
}
func ResponseComments(c *gin.Context, Children interface{}) {
	c.JSON(http.StatusOK, Children)
}
func ResponseCollection(c *gin.Context, e model.Collection) {
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"kind":          e.Kind,
		"productName":   e.ProductName,
		"title":         e.Title,
		"info":          e.Info,
		"imagePath":     e.ImagePath,
		"price":         e.Price,
		"discountPrice": e.DiscountPrice,
		"onSale":        e.Sales,
		"score":         e.Score,
	})
}
