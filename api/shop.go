package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func AnnouncementWriter(c *gin.Context) {
	shopID, _ := strconv.Atoi(c.PostForm("shopID"))
	shopName := c.PostForm("shopName")
	announcement := c.PostForm("announcement")
	if shopID == 0 || announcement == "" || shopName == "" {
		util.ResponseParaError(c)
		return
	}
	err := service.WriteAnnouncement(model.Shop{
		ShopID:       shopID,
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
	announcement := c.PostForm("announcement") //新公告
	if shopID == 0 || announcement == "" {
		util.ResponseParaError(c)
		return
	}
	err := service.ChangeAnnouncement(announcement, shopID)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)

}
func ShowShopProducts(c *gin.Context) {
	way := c.PostForm("way")                        //选择展示方式：销量、价格、评价
	shopID, _ := strconv.Atoi(c.PostForm("shopID")) //需要展示的店铺
	kind := c.PostForm("kind")                      //要展示的种类
	if way == "" || shopID == 0 || kind == "" {
		util.ResponseParaError(c)
		return
	}
	if kind == "all" {
		//展示所有的商品
		err, p := service.SearchAllProductsByShopID(way, shopID)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
		util.ResponseProduct(c, p)
	} else {
		//分类展示商品
		err, p := service.ShowProductByShopID(way, kind, shopID)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
		util.ResponseProduct(c, p)
	}

}
func ProductDetail(c *gin.Context) {
	file, err := c.FormFile("detailPhoto") //商品详情页的图片
	productName := c.PostForm("productName")
	productID, _ := strconv.Atoi(c.PostForm("productID"))
	if err != nil || productName == "" || productID == 0 {
		util.ResponseParaError(c)
		return
	}
	fileName := productName + ".detail.png"
	err = c.SaveUploadedFile(file, "./"+fileName) //上传文件到本地
	if err != nil {
		log.Printf("when upload detail picture error:%v", err)
		util.ResponseNormalError(c, 20003, "upload shop picture fail") //上传商品图片失败
		return
	}
	detailPath := "./" + fileName
	err = service.CreateProductDetail(model.ProductDetail{
		ProductID:   productID,
		ProductName: productName,
		URL:         detailPath,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)

}
func DetailUpdate(c *gin.Context) {
	file, err := c.FormFile("detail") //新的图片信息
	productName := c.PostForm("productName")
	productID, _ := strconv.Atoi(c.PostForm("productID"))
	//入参校验
	if err != nil || productName == "" || productID == 0 {
		util.ResponseParaError(c)
		return
	}
	fileName := productName + ".detail.png"
	err = c.SaveUploadedFile(file, "./"+fileName) //上传文件
	if err != nil {
		log.Printf("when upload detail picture error:%v", err)
		util.ResponseNormalError(c, 20003, "upload shop picture fail") //上传商品图片失败
		return
	}
	detailPath := "./" + fileName
	err = service.ChangeProductDetail(model.ProductDetail{
		ProductID:   productID,
		ProductName: productName,
		URL:         detailPath,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
}
