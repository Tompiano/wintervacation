package dao

import (
	"database/sql"
	"log"
	"wintervacation/model"
)

func InsertProduct(p model.Product) error {
	result, err := DB.Exec("insert into product(kind,productName,title,info,imagePath,price,discountPrice,Sales,shopID,score,number)value(?,?,?,?,?,?,?,?,?,?,?)",
		p.Kind, p.ProductName, p.Title, p.Info, p.ImagePath, p.Price, p.DiscountPrice, p.Sales, p.ShopID, p.Score, p.Number)
	if err != nil {
		log.Printf("when insert prduct informaton error:%v", err)
		return err
	}

	result.LastInsertId()
	result.RowsAffected()
	return err
}

func ListAllProduct(way string, page, pageSize int) (err error, products []*model.Product) {
	var p model.Product
	stmt, err := DB.Prepare("select* from product order by ? limit ?,?")
	if err != nil {
		log.Printf("when search all products,prepare error:%v", err)
		return
	}
	var row *sql.Rows
	if way == "price" {
		//如果选择的排序方式是价格，则降序排列
		row, err = stmt.Query("price desc", page, pageSize)
	} else {
		//如果选择的排序方式是评价或者销量，则升序排列
		row, err = stmt.Query(way, page, pageSize)
	}
	if err != nil {
		log.Printf("when search all products,query error:%v", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&p.ProductID, &p.Kind, &p.ProductName, &p.Title, &p.Info, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Sales, &p.ShopID, &p.Score, &p.Number)
		if err != nil {
			log.Printf("when show all products,scan error:%v", err)
			return
		}
		//返回所有的商品的信息
		productDetails := model.Product{
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

		products = append(products, &productDetails)
	}
	return
}

func SearchCategoriesProduct(kind, way string, page, pageSize int) (err error, products []*model.Product) {
	stmt, err := DB.Prepare("select* from product where kind=? order by ?  limit ?,?")
	if err != nil {
		log.Printf("when search categoried products,prepare error:%v", err)
		return
	}
	var row *sql.Rows
	if way == "price" {
		//如果选择的排序方式是价格，则降序排列
		row, err = stmt.Query(kind, "price desc", page, pageSize)
	} else {
		//如果选择的排序方式是评价或者销量，则升序排列
		row, err = stmt.Query(kind, way, page, pageSize)
	}
	if err != nil {
		log.Printf("when search categoried products,query error:%v", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		var p model.Product
		err = row.Scan(&p.ProductID, &p.Kind, &p.ProductName, &p.Title, &p.Info, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Sales, &p.ShopID, &p.Score, &p.Number)
		if err != nil {
			log.Printf("when show categoried products,scan error:%v", err)
			return
		}
		productDetails := model.Product{
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
		products = append(products, &productDetails)
	}
	return

}

func FuzzySearchProducts(words, way string, page, pageSize int) (err error, p model.Product) {
	stmt, err := DB.Prepare("select* from product where productName like ? order by ? limit ?,?")
	if err != nil {
		log.Printf("when fuzzy search products,prepare error:%v", err)
		return
	}
	var row *sql.Rows
	if way == "price" {
		//如果选择的排序方式是价格，则降序排列
		row, err = stmt.Query("%"+words+"%", "price desc", page, pageSize)
	} else {
		//如果选择的排序方式是评价或者销量，则升序排列
		row, err = stmt.Query("%"+words+"%", way, page, pageSize)
	}
	if err != nil {
		log.Printf("when fuzzy search products,query error:%v", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&p.ProductID, &p.Kind, &p.ProductName, &p.Title, &p.Info, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Sales, &p.ShopID, &p.Score, &p.Number)
		if err != nil {
			log.Printf("when fuzzy search products,scan error:%v", err)
			return
		}
	}
	return
}

func SelectDetail(productID int) (err error, details []*model.ProductDetail) {
	var d model.ProductDetail
	stmt, err := DB.Prepare("select* from detail where productID=?")
	if err != nil {
		log.Printf("when select detail ,mysql prepare error:%v ", err)
		return
	}
	row, err := stmt.Query(productID)
	if err != nil {
		log.Printf("when select detail,mysql query error:%v ", err)
		return
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&d.DetailID, &d.ProductID, &d.ProductName, &d.URL)
		if err != nil {
			log.Printf("when scan error:%v ", err)
			return
		}
		detail := model.ProductDetail{
			DetailID:    d.DetailID,
			ProductID:   productID,
			ProductName: d.ProductName,
			URL:         "http://logalhost:8080" + d.URL,
		}
		details = append(details, &detail)

	}
	return
}
