package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

//选择地址

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

//提交订单

func Commit(c *gin.Context) {
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
	util.OrdersShow(c, t, "已确认") //展示订单状态
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
		util.OrdersShow(c, t, "未支付")
	}
	util.OrdersShow(c, t, "已支付") //展示订单状态

}

//确认收货

func OrderComplete(c *gin.Context) {
	judge := c.PostForm("judge") //判断是否成功收货
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
		util.OrdersShow(c, t, "未收货")
	}
	util.OrdersShow(c, t, "已收货") //展示订单状态

}
