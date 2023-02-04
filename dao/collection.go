package dao

import (
	"log"
	"wintervacation/model"
)

func InsetCollection(e model.Collection) (err error) {
	result, err := DB.Exec("insert into collection(userID,productID)values(?,?)", &e.UserID, &e.ProductID)
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

func SelectCollection(userID int) (Err error, allCollections []*model.Collection) {
	stmt, err := DB.Prepare("select*from collection where userID=?")
	if err != nil {
		log.Printf("when selct collection prepare error:%v ", err)
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
		var e model.Collection
		err = row.Scan(&e.CollectionID, &e.UserID, &e.ProductID)
		if err != nil {
			log.Printf("when selct collection scan error:%v ", err)
			return
		}
		everyCollection := model.Collection{
			CollectionID: e.CollectionID,
			UserID:       e.UserID,
			ProductID:    e.ProductID,
		}
		allCollections = append(allCollections, &everyCollection)

	}
	return
}
func SelectProductsInCollection(productID int) (err error, products []*model.Product) {
	stmt, err := DB.Prepare("select*from product where productID=?")
	if err != nil {
		log.Printf("when select products prepare error:%v ", err)
		return
	}
	defer stmt.Close()
	row, err := stmt.Query(productID)
	if err != nil {
		log.Printf("when select products query error:%v ", err)
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
			log.Printf("when select product scan error:%v ", err)
			return
		}
		everyProduct := model.Product{
			ProductID:     productID,
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
		products = append(products, &everyProduct)
	}
	return
}
