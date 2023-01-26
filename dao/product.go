package dao

import (
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
