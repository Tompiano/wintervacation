package dao

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"time"
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

//添加地址相关的-----------------------------------------------------------------------------------------------------------

func InsertAddress(userID int, address string) (err error) {
	result, err := DB.Exec("insert into address(userID,address)value(?,?)", userID, address)
	if err != nil {
		log.Printf("when insert into address error:%v ", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}

//生成token和refresh_token相关--------------------------------------------------------------------------------------------

func TokenAndRefresh(UserName string, c *gin.Context) (tokenString, refreshString string, err error) {
	claims := model.MyStandardClaims{
		UserName: UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           //在此之前不可用
			ExpiresAt: time.Now().Unix() + 60*60*2, //将jwt设置为2小时过期
			Issuer:    "Tom-Jerry",                 //发行人：Tom-Jerry
		},
	}
	claims2 := model.MyStandardClaims{
		UserName: UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),              //在此之前不可用
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //将jwt设置为7天过期
			Issuer:    "Tom-Jerry",                    //发行人：Tom-Jerry
		},
	}
	mySignedKey := []byte("token") //自定义签名
	//用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claims2)
	// 使用指定的mySignedKey签名并获得完整的编码后的字符串token
	tokenString, err = token.SignedString(mySignedKey)
	if err != nil {
		log.Printf("when signed tokenString error:%v ", err)
		return
	}
	refreshString, err = refresh.SignedString(mySignedKey)
	if err != nil {
		log.Printf("when signed refreshString error:%v ", err)
		return
	}
	return

}
