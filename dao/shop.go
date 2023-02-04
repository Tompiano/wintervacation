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
func UpdateAnnouncement(a model.Shop) (err error) {
	_, err = DB.Exec("update shop set shopID=?,shopName=?,announcement=?", a.ShopID, a.ShopName, a.Announcement)
	if err != nil {
		log.Printf("when update announcement error:%v ", err)
		return
	}
	return
}
func SelectProductByShopID(way, kind string, shopID int) (err error, p model.Product) {
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
		return
	}
	for row.Next() {
		err = row.Scan(&p.ProductID, &p.Kind, &p.ProductName, &p.Title, &p.Info, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Sales, &p.ShopID, &p.Score, &p.Number)
		if err != nil {
			log.Printf("when search products by shopID,scan error:%v", err)
			return
		}
	}
	return
}
func SelectAllProductsByShopID(way string, shopID int) (err error, p model.Product) {
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
	}
	return
}
func InsertProductDetailPhotos(d model.ProductDetail) (err error) {
	result, err := DB.Exec("insert into detail(productID,productName,detail)values(?,?,?,?)", &d.ProductID, &d.ProductName, &d.DetailPath)
	if err != nil {
		log.Printf("when insert product detail,mysql exec error:%v ", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}

func UpdateDetailPhotos(d model.ProductDetail) (err error) {
	_, err = DB.Exec("update detail set detail where ProductID=? and ProductName=?", &d.DetailPath, &d.ProductID, &d.ProductName)
	if err != nil {
		log.Printf("when update product detail,mysql exec error:%v ", err)
		return
	}
	return
}
