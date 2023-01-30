package model

type Collection struct {
	UserID        int    `json:"userID"`
	ProductID     int    `json:"productID"`
	Kind          string `json:"kind"`
	ProductName   string `json:"productName"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImagePath     string `json:"imagePath"`
	Price         int    `json:"price"`
	DiscountPrice int    `json:"discountPrice"`
	Sales         int    `json:"sales"`
	Score         int    `json:"score"`
	ShopID        int    `json:"shopID"`
	ShopName      string `json:"shopName"`
}
