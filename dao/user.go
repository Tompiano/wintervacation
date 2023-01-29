package dao

import (
	"log"
	"wintervacation/model"
)

//注册相关----------------------------------------------------------------------------------------------------------

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
		err = row.Scan(&u.ID, &u.UserName, &u.Phone, &u.Password)
		if err != nil {
			log.Printf("when search userName if exist.mysql scan error:%v", err)
			return
		}
	}
	return
}
func InsertUser(u model.User) (err error) {
	result, err := DB.Exec("insert into user (userName,password,phone) value(?,?,?)",
		u.UserName, u.Password, u.Phone)
	if err != nil {
		log.Printf("when insert user error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}

//登录相关----------------------------------------------------------------------------------------------------------

func SelectUserNameIfFirPassword(userName string) (u model.User) {
	stmt, err := DB.Prepare("select * from user where userName=?")
	if err != nil {
		log.Printf("when selct userName in login,mysql prepare err:%v", err)
		return
	}
	row, err := stmt.Query(userName)
	if err != nil {
		log.Printf("when query username in login,err:%v", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&u.UserName, &u.Password, &u.Phone, &u.ID)
		if err != nil {
			log.Printf("when scan in login,err:%v", err)
			return
		}
	}
	return
}

// 忘记密码相关-------------------------------------------------------------------------------------------------------

func SelectPhoneIfExist(phone string) (u model.User, err error) {
	stmt, err := DB.Prepare("select * from user where phone=?")
	if err != nil {
		log.Printf("when select phone if exist err:%v", err)
		return
	}
	row, err := stmt.Query(phone)
	if err != nil {
		log.Printf("when query phone,err:%v", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&u.UserName, &u.Password, &u.Phone, &u.ID)
		if err != nil {
			log.Printf("when scan in phone,err:%v", err)
			return
		}
	}
	return
}

func UpdatePassword(password, userName string) (err error) {
	_, err = DB.Exec("update user SET password=? where userName=?", password, userName)
	if err != nil {
		log.Printf("when update password,err:%v", err)
		return
	}

	return
}

//添加用户信息相关---------------------------------------------------------------------------------------------------------

func InsertPersonInformation(p model.PersonInformation) (err error) {
	result, err := DB.Exec("insert into personInformation(userName,nickName,gender,phone,email,birthday)value(?,?,?,?,?,?)",
		p.UserName, p.NickName, p.Gender, p.Phone, p.Email, p.Birthday)
	if err != nil {
		log.Printf("when insert into personinformation error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}

//添加头像相关的-----------------------------------------------------------------------------------------------------------

func InsertPersonAvatar(a model.Avatar) (err error) {
	result, err := DB.Exec("insert into avatar(userID,avatarName,avatarPath) value(?,?,?)",
		a.UserID, a.AvatarName, a.AvatarPath)
	if err != nil {
		log.Printf("when insert into person avatar error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}
