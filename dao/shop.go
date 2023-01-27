package dao

import (
	"log"
	"wintervacation/model"
)

func InsertAnnouncement(a model.Shop) (err error) {
	result, err := DB.Exec("insert into shop(shopID,shopName,announcement)value(?,?,?)", a.ShopID, a.ShopName, a.Announcement)
	if err != nil {
		log.Printf("when insert announcement,exec error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}
func UpdateAnnouncement(a model.Shop) (err error) {
	_, err = DB.Exec("update shop set shopID=?,shopName=?,announcement=?", a.ShopID, a.ShopName, a.Announcement)
	if err != nil {
		log.Printf("when update announcement error:%v ", err)
		return
	}
	return
}
