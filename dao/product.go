package dao

import (
	"database/sql"
	"log"
	"wintervacation/model"
)

func InsertProduct(p model.Product) error {
	result, err := DB.Exec("insert into product(kind,productName,title,info,imagePath,price,discountPrice,Sales,shopID)value(?,?,?,?,?,?,?,?,?)",
		p.Kind, p.ProductName, p.Title, p.Info, p.ImagePath, p.Price, p.DiscountPrice, p.Sales, p.ShopID)
	if err != nil {
		log.Printf("when insert prduct informaton error:%v", err)
		return err
	}

	result.LastInsertId()
	result.RowsAffected()
	return err
}

func ListAllProduct(way string, page, pageSize int) (err error, p model.Product) {
	stmt, err := DB.Prepare("select* from product where kind=? order by ? and limit ?,?")
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
		err = row.Scan(&p.Kind, &p.ProductName, &p.ShopName, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Info, &p.Title, &p.Sales)
		if err != nil {
			log.Printf("when show all products,scan error:%v", err)
			return
		}
	}
	return
}

func SearchCategoriesProduct(kind, way string, page, pageSize int) (err error, p model.Product) {
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
		err = row.Scan(&p.Kind, &p.ProductName, &p.ShopName, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Info, &p.Title, &p.Sales)
		if err != nil {
			log.Printf("when show categoried products,scan error:%v", err)
			return
		}
	}
	return

}

func FuzzySearchProducts(words, way string, page, pageSize int) (err error, p model.Product) {
	stmt, err := DB.Prepare("select* from product where productName like ? order by ?limit ?,?")
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
		err = row.Scan(&p.Kind, &p.ProductName, &p.ShopName, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Info, &p.Title, &p.Sales)
		if err != nil {
			log.Printf("when fuzzy search products,scan error:%v", err)
			return
		}
	}
	return
}

func SelectDetail(productID int, productName string) (err error, d model.ProductDetail) {
	stmt, err := DB.Prepare("select* from detail where productID=? and ProductName=?")
	if err != nil {
		log.Printf("when select detail ,mysql prepare error:%v ", err)
		return
	}
	row, err := stmt.Query(productID, productName)
	if err != nil {
		log.Printf("when select detail,mysql query error:%v ", err)
		return
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&d.ProductID, &d.ProductName, &d.Detail)
		if err != nil {
			log.Printf("when scan error:%v ", err)
			return
		}
	}
	return
}
