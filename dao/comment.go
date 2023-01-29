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

func SelectAllComments(productID, parentID int) (err error, t model.Comment) {
	stmt, err := DB.Prepare("select*from comment where productID=? and parentID=?")
	if err != nil {
		log.Printf("when select all comments,mysql prepare error:%v ", err)
		return
	}
	row, err := stmt.Query(productID, parentID)
	if err != nil {
		log.Printf("when select all comments,mysql query error:%v ", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&t.CommentID, &t.Content, &t.ProductID, t.UserID, &t.ParentID)
		if err != nil {
			log.Printf("when select all comments,scan error:%v ", err)
			return
		}
	}
	return
}
