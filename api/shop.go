package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func AnnouncementWriter(c *gin.Context) {
	shopID, _ := strconv.Atoi(c.PostForm("shopID"))
	shopName := c.PostForm("shopName")
	announcement := c.PostForm("announcement")
	if shopID == 0 || shopName == "" || announcement == "" {
		util.ResponseParaError(c)
		return
	}
	err := service.WriteAnnouncement(model.Shop{
		ShopID:       shopID,
		ShopName:     shopName,
		Announcement: announcement,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
}
func AnnouncementUpdate(c *gin.Context) {
	shopID, _ := strconv.Atoi(c.PostForm("shopID"))
	shopName := c.PostForm("shopName")
	announcement := c.PostForm("announcement") //新公告
	if shopID == 0 || shopName == "" || announcement == "" {
		util.ResponseParaError(c)
		return
	}
	err := service.ChangeAnnouncement(model.Shop{
		ShopID:       shopID,
		ShopName:     shopName,
		Announcement: announcement,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)

}
func ShowShopProducts(c *gin.Context) {
	
}
