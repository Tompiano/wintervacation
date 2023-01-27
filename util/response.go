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
