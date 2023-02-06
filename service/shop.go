package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func WriteAnnouncement(a model.Shop) (err error) {
	return dao.InsertAnnouncement(a)
}
func ChangeAnnouncement(announcement string, shopID int) (err error) {
	return dao.UpdateAnnouncement(announcement, shopID)
}
func ShowProductByShopID(way, kind string, shopID int) (err error, shopProducts []*model.Product) {
	return dao.SelectProductByShopID(way, kind, shopID)
}
func SearchAllProductsByShopID(way string, shopID int) (err error, shopProducts []*model.Product) {
	return dao.SelectAllProductsByShopID(way, shopID)
}
func CreateProductDetail(d model.ProductDetail) (err error) {
	return dao.InsertProductDetailPhotos(d)
}
func ChangeProductDetail(d model.ProductDetail) (err error) {
	return dao.UpdateDetailPhotos(d)
}
