package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wintervacation/model"
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
	amount, _ := strconv.Atoi(c.PostForm("amount"))
	userID, _ := strconv.Atoi(c.PostForm("userID")) //用户的ID
	address := c.PostForm("address")                //用户选择的地址
	payMethod := c.PostForm("payMethod")            //支付方式
	if address == "" || userID == 0 || payMethod == "" || amount == 0 {
		util.ResponseParaError(c)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	t := model.Orders{
		Time:          nowTime,
		UserID:        userID,
		Address:       address,
		PaymentAmount: amount,
		PayMethod:     payMethod,
	}
	util.OrdersShow(c, t, "已确认") //展示订单
}

//订单成功页面

func OrderSuccess(c *gin.Context) {
	judge := c.PostForm("judge") //判断是否成功付费
	amount, _ := strconv.Atoi(c.PostForm("amount"))
	userID, _ := strconv.Atoi(c.PostForm("userID")) //用户的ID
	address := c.PostForm("address")                //用户选择的地址
	payMethod := c.PostForm("payMethod")            //支付方式
	if address == "" || userID == 0 || payMethod == "" || amount == 0 {
		util.ResponseParaError(c)
		return
	}
	nowTime := time.Now().Unix() //当前时间
	t := model.Orders{
		Time:          nowTime,
		UserID:        userID,
		Address:       address,
		PaymentAmount: amount,
		PayMethod:     payMethod,
	}
	if judge != "true" {
		util.ResponseNormalError(c, 30002, "no pay")
	}
	util.OrdersShow(c, t, "已支付") //展示订单

}
