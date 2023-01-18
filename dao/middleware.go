package dao

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"wintervacation/model"
)

//解析token

func Authentication(mySignedKey string, tokenString string, c *gin.Context) (*jwt.Token, model.MyStandardClaims, error) {
	claims := model.MyStandardClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &model.MyStandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySignedKey, nil
	})
	if err != nil {
		log.Printf("when parse token error:%v", err)
		return token, claims, err
	}
	return token, claims, err
}

func SelectUserID(userID string) (u model.User, err error) {
	stmt, err := DB.Prepare("select * from user where ID=?")
	if err != nil {
		log.Printf("when select userId error:%v", err)
		return
	}
	row, err := stmt.Query(userID)
	if err != nil {
		log.Printf("when query userId error:%v", err)
		return
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&u.ID, &u.UserName, &u.Phone, &u.Password)
		if err != nil {
			log.Printf("when search userID if exist.mysql scan error:%v", err)
			return
		}
	}
	return

}
