package model

type ShoppingCart struct {
	CartID    int `json:"cartID"`
	UserID    int `json:"userID"`
	ProductID int `json:"productID"`
	Amount    int `json:"amount"` //数量
	Check     int `json:"check"`  //是否勾选
}
