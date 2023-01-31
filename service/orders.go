package service

import (
	"wintervacation/dao"
	"wintervacation/model"
)

func PrepareOrders(userID int) (err error, m model.Address, addr map[int]string, length int) {
	return dao.SelectOrdersPrepare(userID)
}
