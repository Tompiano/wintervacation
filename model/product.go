package model

type Product struct {
	ProductID     int    `json:"productID"` //标识着唯一的商品
	Kind          string `json:"kind"`      //商品种类
	ProductName   string `json:"productName"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImagePath     string `json:"imagePath"`
	Price         int    `json:"price"`
	DiscountPrice int    `json:"discountPrice"`
	Sales         int    `json:"Sales"`
	Score         int    `json:"score"`
	ShopID        int    `json:"shopID"`   //店家的id,标识着唯一的店家
	ShopName      string `json:"shopName"` //店家的名字
}

type ShowProduct struct {
	Kind       string `json:"kind"`       //商品种类
	Way        string `json:"way"`        //商品排序方式
	PageNumber int    `json:"pageNumber"` //页数
	PageSize   int    `json:"pageSize"`   //每页容量
}
