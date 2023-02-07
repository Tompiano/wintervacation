package dao

import (
	"database/sql"
	"log"
	"wintervacation/model"
)

func InsertAnnouncement(a model.Shop) (err error) {
	result, err := DB.Exec("insert into shop(shopID,shopName,announcement)value(?,?,?)", a.ShopID, a.ShopName, a.Announcement)
	if err != nil {
		log.Printf("when insert announcement,exec error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}
func UpdateAnnouncement(announcement string, shopID int) (err error) {
	result, err := DB.Exec("update shop set announcement=? where shopID=?", announcement, shopID)
	if err != nil {
		log.Printf("when update announcement error:%v ", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}
func SelectProductByShopID(way, kind string, shopID int) (err error, shopProducts []*model.Product) {
	var p model.Product
	stmt, err := DB.Prepare("select*from product where shopID=? and kind=? order by ?")
	if err != nil {
		log.Printf("when select product by shopID,mysql prepare error:%v", err)
		return
	}
	var row *sql.Rows
	if way == "price" {
		//如果选择的排序方式是价格，则降序排列
		row, err = stmt.Query(shopID, kind, "price desc")
	} else {
		//如果选择的排序方式是评价或者销量，则升序排列
		row, err = stmt.Query(shopID, kind, way)
	}

	if err != nil {
		log.Printf("when search products by shopID,query error:%v", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		log.Printf("close error:%v ", err)
		return
	}
	for row.Next() {
		err = row.Scan(&p.ProductID, &p.Kind, &p.ProductName, &p.Title, &p.Info, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Sales, &p.ShopID, &p.Score, &p.Number)
		if err != nil {
			log.Printf("when search products by shopID,scan error:%v", err)
			return
		}
		temporary := model.Product{
			ProductID:     p.ProductID,
			Kind:          p.Kind,
			ProductName:   p.ProductName,
			Title:         p.Title,
			Info:          p.Info,
			ImagePath:     p.ImagePath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			Sales:         p.Sales,
			ShopID:        p.ShopID,
			Score:         p.Score,
			Number:        p.Number,
		}
		shopProducts = append(shopProducts, &temporary)
	}
	return
}
func SelectAllProductsByShopID(way string, shopID int) (err error, shopProducts []*model.Product) {
	stmt, err := DB.Prepare("select*from product where shopID=? order by ?")
	if err != nil {
		log.Printf("when search all products by shopID,mysql prepare error:%v", err)
		return
	}
	var row *sql.Rows
	if way == "price" {
		//如果选择的排序方式是价格，则降序排列
		row, err = stmt.Query(shopID, "price desc")
	} else {
		//如果选择的排序方式是评价或者销量，则升序排列
		row, err = stmt.Query(shopID, way)
	}

	if err != nil {
		log.Printf("when search all products by shopID,query error:%v", err)
		return
	}
	var p model.Product
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&p.ProductID, &p.Kind, &p.ProductName, &p.Title, &p.Info, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Sales, &p.ShopID, &p.Score, &p.Number)
		if err != nil {
			log.Printf("when search all products by shopID,scan error:%v", err)
			return
		}
		temporary := model.Product{
			ProductID:     p.ProductID,
			Kind:          p.Kind,
			ProductName:   p.ProductName,
			Title:         p.Title,
			Info:          p.Info,
			ImagePath:     p.ImagePath,
			Price:         p.Price,
			DiscountPrice: p.DiscountPrice,
			Sales:         p.Sales,
			ShopID:        p.ShopID,
			Score:         p.Score,
			Number:        p.Number,
		}
		shopProducts = append(shopProducts, &temporary)
	}
	return
}
func InsertProductDetailPhotos(d model.ProductDetail) (err error) {
	_, err = DB.Exec("insert into detail(productID,productName,detailPath)values(?,?,?)", d.ProductID, d.ProductName, d.URL)
	if err != nil {
		log.Printf("when insert product detail,mysql exec error:%v ", err)
		return
	}
	return
}

func UpdateDetailPhotos(d model.ProductDetail) (err error) {
	_, err = DB.Exec("update detail set detailPath where ProductID=? ", d.URL, d.ProductID)
	if err != nil {
		log.Printf("when update product detail,mysql exec error:%v ", err)
		return
	}
	return
}
