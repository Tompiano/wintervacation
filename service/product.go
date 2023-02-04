package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func CreateProduct(p model.Product) (err error) {
	err = dao.InsertProduct(p)
	return
}
func ShowAllProduct(way string, page, pageSize int) (err error, products []*model.Product) {
	return dao.ListAllProduct(way, page, pageSize)
}
func ShowCategoriesProduct(kind, way string, page, pageSize int) (err error, products []*model.Product) {
	return dao.SearchCategoriesProduct(kind, way, page, pageSize)
}
func ExploreProducts(words, way string, page, pageSize int) (err error, p model.Product) {
	return dao.FuzzySearchProducts(words, way, page, pageSize)
}
func SearchDetail(productID int) (err error, details []*model.ProductDetail) {
	return dao.SelectDetail(productID)
}
