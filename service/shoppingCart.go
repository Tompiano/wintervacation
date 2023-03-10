package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func SearchProductsInCart(productID int) (err error, s model.ShoppingCart) {
	return dao.SelectProductID(productID)
}

func ChangeAmount(amount int, productID int) (err error) {
	return dao.UpdateAmount(amount, productID)
}

func AddProductsInCart(s model.ShoppingCart) (err error) {
	return dao.InsertProductInCart(s)
}
func DeleteAllProducts() (err error) {
	return dao.DeleteAllProductsInCart()
}
func DeleteSomeProducts(productID int) (err error) {
	return dao.DeleteSomeProductsInCart(productID)
}
func SearchProductIfEnough(productID, number int) (err error, judge bool, p model.Product) {
	return dao.SelectProductsIfEnough(productID, number)
}
func ChangeCheck(userID, productID, check int) (err error) {
	return dao.UpdateCheck(userID, productID, check)
}
func ListCheck(userID int) (err error, cart []*model.ShoppingCart) {
	return dao.SelectCheck(userID)
}
func ChangeProductsNumber(productID, number int) (err error) {
	return dao.UpdateProductsNumber(productID, number)
}
