package dao

import (
	"log"
	"wintervacation/model"
)

func SelectOrdersPrepare(userID int) (err error, m model.Address, addr map[int]string, length int) {

	stmt, err := DB.Prepare("select*from address where userID=?")
	if err != nil {
		log.Printf("when select,prepare error:%v ", err)
		return
	}
	defer stmt.Close()
	row, err := stmt.Query(userID)
	if err != nil {
		log.Printf("when query error:%v ", err)
		return
	}
	columns, err := row.Columns()
	if err != nil {
		log.Printf("when columns error:%v ", err)
		return
	}
	length = len(columns)
	addr = make(map[int]string, length)
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for i := 0; row.Next(); i++ {
		err = row.Scan(&m.UserID, &m.AddressName)
		addr[i] = m.AddressName
	}
	return
}
