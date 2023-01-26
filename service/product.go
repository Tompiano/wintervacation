package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func CreateProduct(p model.Product) (err error) {
	err = dao.InsertProduct(p)
	return
}
func ShowAllProduct() (err error, p model.Product) {
	return dao.ListAllProduct()
}
func ShowCategoriesProduct(kind, way string) (err error, p model.Product) {
	return dao.SearchCategoriesProduct(kind, way)
}
func ExploreProducts(words, way string, page, pageSize int) (err error, p model.Product) {
	return dao.FuzzySearchProducts(words, way, page, pageSize)
}
