package model

type Orders struct {
	Time          int64  `json:"time"`          //时间
	UserID        int    `json:"userID"`        //用户的ID
	Address       string `json:"address"`       //用户地址
	PaymentAmount int    `json:"paymentAmount"` //订单的总金额
	PayMethod     string `json:"PayMethod"`     //订单支付方式
}

type Address struct {
	UserID      int    `json:"userID"`      //用户的ID
	AddressName string `json:"addressName"` //地址
}
