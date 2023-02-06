package model

type ShoppingCart struct {
	UserID    int `json:"userID"`
	ProductID int `json:"productID"`
	Amount    int `json:"amount"` //数量
	Price     int `json:"price"`  //单价
}
