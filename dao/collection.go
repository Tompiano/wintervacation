package dao

import (
	"log"
	"wintervacation/model"
)

func InsetCollection(e model.Collection) (err error) {
	result, err := DB.Exec("insert into collection(userID,productID,kind,productName,title,info,imagePath,price,discountPrice,Sales,shopID,score)",
		&e.UserID, &e.ProductID, &e.Kind, &e.ProductName, &e.Title, e.Info, e.ImagePath, &e.Price, &e.DiscountPrice, &e.Sales, &e.ShopID, &e.ShopID, &e.Score)
	if err != nil {
		log.Printf("when insert collection error:%v ", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}

func DeleteCollections(productID int) (err error) {
	_, err = DB.Exec("delete from collection where productID=?", productID)
	if err != nil {
		log.Printf("when delete collection error:%v ", err)
		return
	}
	return
}

func SelectCollection(userID int) (Err error, e model.Collection) {
	stmt, err := DB.Prepare("select*from collection where userID=?")
	if err != nil {
		log.Printf("when prepare error:%v ", err)
		return
	}
	defer stmt.Close()
	row, err := stmt.Query(userID)
	if err != nil {
		log.Printf("when query error:%v ", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&e.UserID, &e.Info, &e.Kind, &e.Sales, &e.Price, &e.ProductID, &e.ShopID, &e.ImagePath, &e.Title, &e.DiscountPrice, &e.ShopName, &e.ProductName)
		if err != nil {
			log.Printf("when scan error:%v ", err)
			return
		}
	}
	return
}
