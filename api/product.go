package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func Creator(c *gin.Context) {

	kind := c.PostForm("kind")                                    //商品种类
	productName := c.PostForm("productName")                      //商品名称
	title := c.PostForm("title")                                  //商品标题
	info := c.PostForm("info")                                    //商品简要描述
	price, _ := strconv.Atoi(c.PostForm("price"))                 //商品的价格
	discountPrice, _ := strconv.Atoi(c.PostForm("discountPrice")) //商品打折后的价格
	shopID, _ := strconv.Atoi(c.PostForm("shopID"))               //店铺的唯一标识符
	file, err := c.FormFile("file")                               //商品图片文件
	//入参校验
	if err != nil || kind == "" || productName == "" || title == "" || info == "" || price == 0 || discountPrice == 0 || shopID == 0 {
		util.ResponseParaError(c)
		return
	}
	Sales := 0 //商品的销量初始添加时候为0
	fileName := productName + ".png"
	err = c.SaveUploadedFile(file, "./"+fileName) //上传文件到本地
	if err != nil {
		log.Printf("when upload shop picture error:%v", err)
		util.ResponseNormalError(c, 20003, "upload shop picture fail") //上传商品图片失败
		return
	}
	ImagePath := "http://localhost:8080" + "./" + fileName
	//传入数据到数据库前对参数的要求
	if len(productName) > 50 {
		util.ResponseNormalError(c, 20004, "productName excessive")
		return
	}
	if len(title) > 50 {
		util.ResponseNormalError(c, 20005, "tile excessive")
		return
	}
	if len(info) > 500 {
		util.ResponseNormalError(c, 20006, "info excessive")
		return
	}
	if strings.Contains(info, "毒品") || strings.Contains(title, "毒品") || strings.Contains(productName, "毒品") {
		util.ResponseNormalError(c, 20007, " contains sensitive words")
		return
	}

	err = service.CreateProduct(model.Product{
		Kind:          kind,
		ProductName:   productName,
		Title:         title,
		Info:          info,
		Price:         price,
		DiscountPrice: discountPrice,
		Sales:         Sales,
		ShopID:        shopID,
		ImagePath:     ImagePath,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
}

func Show(c *gin.Context) {
	//获取用户想要的商品展示方式，想要展示的种类，和前端展示的页数和每页的容量
	//way,kind,pageNumber,pageSize
	Kind := c.PostForm("kind")
	Way := c.PostForm("way")
	PageNumber, _ := strconv.Atoi(c.PostForm("pageNumber"))
	PageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	if Kind == "" || Way == "" || PageNumber == 0 || PageSize == 0 {
		log.Printf("kind:%v,way:%v,pageNumber:%v.pageSize:%v", Kind, Way, PageNumber, PageSize)
		util.ResponseParaError(c)
		return
	}
	//这里是展示从page到（pageSize+page）的数据
	page := (PageNumber - 1) * PageNumber
	//分类展示商品
	//这里只输入了三个种类的商品的信息，所以只做了三个种类的商品的展示
	if Kind == "all" {

		//如果选择“all”,展示所有商品
		err, products := service.ShowAllProduct(Way, page, PageSize)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
		util.ResponseProduct(c, products)
	} else {
		//根据选择来分类别展示商品
		err, products := service.ShowCategoriesProduct(Kind, Kind, page, PageSize)
		if err != nil {
			util.ResponseInternalError(c)
			return
		}
		util.ResponseProduct(c, products)
	}
}

func Explore(c *gin.Context) {
	//采用MySQL的模糊搜索，利用limit分页展示
	words := c.PostForm("words")                            //用户输入的查询词语
	way := c.PostForm("way")                                //用户选择的排序方式：价格/销量/评价
	pageNumber, _ := strconv.Atoi(c.PostForm("pageNumber")) //页数
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))     //页面容量
	if words == "" || way == "" || pageNumber == 0 || pageSize == 0 {
		util.ResponseParaError(c)
		return
	}
	page := (pageNumber - 1) * pageSize //计算出总体的商品数量规模
	err, p := service.ExploreProducts(words, way, page, pageSize)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseProduct(c, p) //返回商品的所有信息

}

func Detail(c *gin.Context) {
	//商品详情页的展示
	//这里商品详情页都是商家提供的商品的详情展示的图片，因而单独放一张表专门储存
	//每次只展示一件商品的所有详情页
	productID, _ := strconv.Atoi(c.PostForm("productID"))
	if productID == 0 {
		util.ResponseParaError(c)
		return
	}

	err, details := service.SearchDetail(productID)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseDetail(c, details)
}
