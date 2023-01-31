package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wintervacation/service"
	"wintervacation/util"
)

//订单准备页面

func Prepare(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	if userID == 0 {
		util.ResponseParaError(c)
		return
	}

	//获取用户的各个address和address的个数
	err, m, address, length := service.PrepareOrders(userID)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}

	//返回用户的ID和各个address，以供用户进行选择
	for i := 0; i < length; i++ {
		util.ResponseOrdersPrepare(c, m.UserID, address[i])
	}

}

//订单展示页面

func OrdersShow(c *gin.Context) {
	userID, _ := strconv.Atoi(c.PostForm("userID")) //用户的ID
	address := c.PostForm("address")                //用户选择的地址
	if address == "" || userID == 0 {
		util.ResponseParaError(c)
		return
	}

}

//订单成功页面
