package dao

import (
	"log"
	"wintervacation/model"
)

func InsertComment(m model.Comment) (err error) {
	result, err := DB.Exec("insert into comment(userID,productID,parentID,content)value(?,?,?,?)",
		m.UserID, m.ProductID, m.ParentID, m.Content)
	if err != nil {
		log.Printf("when insert comment,mysql exec error:%v ", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}

func UpdateComment(commentID int, content string) (err error) {
	_, err = DB.Exec("update comment set content=? where commentID=? ", content, commentID)
	if err != nil {
		log.Printf("when delete comment,exec error:%v ", err)
		return
	}
	return
}
