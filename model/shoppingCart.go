package model

type ShoppingCart struct {
	UserID   int     `json:"userID"`
	UserName string  `json:"userName"`
	Product  Product //Product结构体内嵌
	Amount   int     `json:"amount"`  //数量
	Account  int     `json:"account"` //总金额
}
