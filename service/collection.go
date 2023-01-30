package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func JoinCollection(e model.Collection) (err error) {
	return dao.InsetCollection(e)
}
func DeleteCollectionByID(productID int) (err error) {
	return dao.DeleteCollections(productID)
}
func LookCollections(userID int) (Err error, e model.Collection) {
	return dao.SelectCollection(userID)
}
