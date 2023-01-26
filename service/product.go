package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func CreateProduct(p model.Product) (err error) {
	err = dao.InsertProduct(p)
	return
}
