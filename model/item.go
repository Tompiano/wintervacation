package model

type Product struct {
	ProductID     int    `json:"productID"` //标识着唯一的商品
	Name          string `json:"name"`
	CategoryID    string `json:"categoryID"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImagePath     string `json:"imagePath"`
	Price         int    `json:"price"`
	DiscountPrice int    `json:"discountPrice"`
	OnSale        int    `json:"onSale"`
	ShopID        int    `json:"shopID"`   //店家的id,标识着唯一的店家
	ShopName      string `json:"shopName"` //店家的名字
}

type ProductImage struct {
	ProductID int    `json:"productID"` //商品的唯一标识符
	ImagePath string `json:"imagePath"`
}

type ProductDetail struct {
	Product   Product //Product结构体内嵌
	ProductID int     `json:"productID"` //商品的唯一标识符
	ImagePath string  `json:"imagePath"`
}
type ShoppingCart struct {
	UserID    int     `json:"UserID"`
	Product   Product //Product结构体内嵌
	ProductID int     `json:"productID"` //商品的唯一标识符
	ShopID    int     `json:"shopID"`    //店家的唯一标识符
	Account   int     `json:"account"`   //总金额
}

type Order struct {
	User      User
	UserID    int `json:"UserID"`
	Product   Product
	ProductID int `json:"productID"`
	ShopID    int `json:"shopID"`
}
type Shop struct {
	ShopID       int    `json:"shopID"`
	ShopName     string `json:"shopName"`
	Announcement string `json:"announcement"`
}
type Collection struct {
	UserID    int    `json:"userID"`
	UserName  string `json:"userName"`
	Product   Product
	ProductID int  `json:"productID"`
	Shop      Shop //shop结构体内嵌
	ShopID    int  `json:"shopID"`
}
type Address struct {
	User     User
	UserID   int    `json:"userID"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
}
type Carousel struct {
	ImagePath string `json:"imagePath"`
	Product   Product
	ProductID int `json:"productID"`
}
