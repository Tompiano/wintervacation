package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wintervacation/util"
)

func Join(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	productID, _ := strconv.Atoi(c.PostForm("productID"))
	if userID == 0 || productID == 0 {
		util.ResponseParaError(c)
		return
	}

}


