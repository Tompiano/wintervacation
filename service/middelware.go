package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func SearchUserID(userID string) (u model.User, err error) {
	u, err = dao.SelectUserID(userID)
	return
}
