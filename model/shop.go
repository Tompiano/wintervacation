package model

type Shop struct {
	ShopID       int    `json:"shopID"`
	ShopName     string `json:"shopName"`
	Announcement string `json:"announcement"`
}
