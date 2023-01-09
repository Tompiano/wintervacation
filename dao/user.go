package dao

import (
	"log"
	"wintervacation/model"
)

//注册相关

func SelectUserInformation(UserName string) (u model.User) { //找用户名是否已经存在
	stmt, err := DB.Prepare("select * from user where userName=?")
	if err != nil {
		log.Printf("when search userName if exist,mysql prepare error:%v", err)
		return
	}
	row, err := stmt.Query(UserName)
	if err != nil {
		log.Println(err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&u.PersonInformation, &u.ID, &u.UserName, &u.Phone)
		if err != nil {
			log.Printf("when search userName if exist.mysql scan error:%v", err)
			return
		}
	}
	return
}
func InsertUser(u model.User) (err error) {
	result, err := DB.Exec("insert into user (userName,password,phone,personInformation) value(?,?,?,?)",
		u.UserName, u.Password, u.Phone, u.PersonInformation)
	if err != nil {
		log.Printf("when insert user error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}

//登录相关
