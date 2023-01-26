package model

type Product struct {
	ProductID     string `json:"productID"` //标识着唯一的商品
	Kind          string `json:"kind"`      //商品种类
	ProductName   string `json:"productName"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImagePath     string `json:"imagePath"`
	Price         int    `json:"price"`
	DiscountPrice int    `json:"discountPrice"`
	Sales         int    `json:"Sales"`
	ShopID        int    `json:"shopID"`   //店家的id,标识着唯一的店家
	ShopName      string `json:"shopName"` //店家的名字
}

type ProductDetail struct {
	Product    Product //Product结构体内嵌
	PageNumber int     `json:"pageNumber"`
}

type ShowProduct struct {
	Kind        string `json:"kind"` //商品种类
	ProductName string `json:"productName"`
}