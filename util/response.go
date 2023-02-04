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
func ResponseProduct(c *gin.Context, products interface{}) {
	c.JSON(http.StatusOK, products)
}
func ResponseDetail(c *gin.Context, details interface{}) {
	c.JSON(http.StatusOK, details)
}
func ResponseComments(c *gin.Context, Children interface{}) {
	c.JSON(http.StatusOK, Children)
}
func ResponseCollection(c *gin.Context, collections interface{}) {
	c.JSON(http.StatusOK, collections)
}
func ResponseOrdersPrepare(c *gin.Context, userID int, address string) {
	c.JSON(http.StatusOK, gin.H{
		"userID":  userID,
		"address": address,
	})
}
func ResponsePay(c *gin.Context, productID, number, price int) {
	c.JSON(http.StatusOK, gin.H{
		"productID": productID,
		"number":    number,
		"price":     price,
	})
}
func ResponseInformation(c *gin.Context, userID, total int) {
	c.JSON(http.StatusOK, gin.H{
		"userID": userID,
		"total":  total,
	})
}
func OrdersShow(c *gin.Context, t model.Orders, info string) {
	c.JSON(http.StatusOK, gin.H{
		"status":        info,
		"time":          t.Time,
		"userID":        t.UserID,
		"address":       t.Address,
		"paymentAmount": t.PaymentAmount,
		"payMethod":     t.PayMethod,
	})
}
