package model

type Orders struct {
	UserID        int     `json:"userID"`        //用户的ID
	PaymentAmount int     `json:"paymentAmount"` //订单的总金额
	PayMethod     string  `json:"PayMethod"`     //订单支付方式
	Product       Product //商品信息
}

type Address struct {
	UserID      int    `json:"userID"`      //用户的ID
	AddressName string `json:"addressName"` //地址
}
