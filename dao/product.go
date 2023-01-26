package dao

import (
	"database/sql"
	"log"
	"wintervacation/model"
)

func InsertProduct(p model.Product) error {
	result, err := DB.Exec("insert into product(kind,productName,tile,info,imagePath,price,discountPrice,onSale,shopID)value(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Printf("when insert prduct informaton error:%v", err)
		return err
	}

	result.LastInsertId()
	result.RowsAffected()
	return err
}

func ListAllProduct() (err error, p model.Product) {
	rows, err := DB.Query("select * from product")
	if err != nil {
		log.Printf("when show all products,query error:%v", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&p.Kind, &p.ProductName, &p.ShopName, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Info, &p.Title, &p.Sales)
		if err != nil {
			log.Printf("when show all products,scan error:%v", err)
			return
		}
	}
	return
}

func SearchCategoriesProduct(kind, way string) (err error, p model.Product) {
	stmt, err := DB.Prepare("select* from product where kind=? order by ?")
	if err != nil {
		log.Printf("when search categoried products,prepare error:%v", err)
		return
	}
	var row *sql.Rows
	if way == "price" {
		//如果选择的排序方式是价格，则降序排列
		row, err = stmt.Query(kind, "price desc")
	} else {
		//如果选择的排序方式是评价或者销量，则升序排列
		row, err = stmt.Query(kind, way)
	}
	if err != nil {
		log.Printf("when fuzzy search products,query error:%v", err)
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
	for row.Next() {
		err = row.Scan(&p.Kind, &p.ProductName, &p.ShopName, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Info, &p.Title, &p.Sales)
		if err != nil {
			log.Printf("when fuzzy search products,scan error:%v", err)
			return
		}
	}
	return
}
