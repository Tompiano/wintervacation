package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func WriteAnnouncement(a model.Shop) (err error) {
	return dao.InsertAnnouncement(a)
}
func ChangeAnnouncement(a model.Shop) (err error) {
	return dao.UpdateAnnouncement(a)
}
