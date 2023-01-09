package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func handleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("中间件")
	}
}
