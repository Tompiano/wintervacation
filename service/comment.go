package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func AddComment(m model.Comment) (err error) {
	return dao.InsertComment(m)
}
func DeleteComment(commentID int, content string) (err error) {
	return dao.UpdateComment(commentID, content)
}
func SearchComments(productID, parentID int, t *model.Comment) (err error, Children []*model.Comment) {
	return dao.SelectAllComments(productID, parentID, t)
}
