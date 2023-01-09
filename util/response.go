package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
