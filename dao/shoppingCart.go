package dao

import (
	"log"
	"wintervacation/model"
)

func SelectProductID(productID int) (err error, s model.ShoppingCart) {
	stmt, err := DB.Prepare("select productID from cart where productID=?")
	if err != nil {
		log.Printf("when selct productID in cart,mysql prepare err:%v", err)
		return
	}
	row, err := stmt.Query(productID)
	if err != nil {
		log.Printf("when query productID in cart,err:%v", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&s.ProductID)
		if err != nil {
			log.Printf("when scan productID in cart,err:%v", err)
			return
		}
	}
	return
}

func UpdateAmount(amount int, productID int) (err error) {
	_, err = DB.Exec("update cart set amount=? where productID=? ", amount, productID)
	_, err = DB.Exec("update cart set check=? where productID=?", 1, productID)
	if err != nil {
		log.Printf("when update amount,exec error:%v", err)
		return
	}
	return
}

func InsertProductInCart(s model.ShoppingCart) (err error) {
	result, err := DB.Exec("insert into cart(userID,productID,amount,check)values(?,?,?,?)", s.UserID, s.ProductID, s.Amount, s.Check)
	if err != nil {
		log.Printf("when insert into cart,exec error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return

}

func DeleteAllProductsInCart() (err error) {
	_, err = DB.Exec("delete from cart")
	if err != nil {
		log.Printf("when delete all products,error :%v", err)
		return
	}
	return
}
func DeleteSomeProductsInCart(productID int) (err error) {
	result, err := DB.Exec("delete from cart where productID=?", productID)
	if err != nil {
		log.Printf("when delete some products,error:%v", err)
		return
	}
	result.LastInsertId()
	result.RowsAffected()
	return
}
func SelectProductsIfEnough(productID, number int) (err error, judge bool, p model.Product) {

	stmt, err := DB.Prepare("select*from cart where productID=?")
	if err != nil {
		log.Printf("when prepare error:%v ", err)
		return
	}
	row, err := stmt.Query(productID)
	if err != nil {
		log.Printf("when query error:%v ", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&p.ProductID, &p.Kind, &p.ProductName, &p.Title, &p.Info, &p.ImagePath, &p.Price, &p.DiscountPrice, &p.Sales, &p.ShopID, &p.Score, &p.Number)
		if err != nil {
			log.Printf("when scan error:%v ", err)
			return
		}
		if p.Number < number {
			judge = false
		} else {
			judge = true
		}
	}
	return
}
func UpdateCheck(userID, productID, check int) (err error) {
	_, err = DB.Exec("update cart set check=? where userID=? and productID=?", check, userID, productID)
	if err != nil {
		log.Printf("when update cart check error :%v ", err)
		return
	}
	return
}
func SelectCheck(userID int) (err error, cart []*model.ShoppingCart) {
	var s model.ShoppingCart
	stmt, err := DB.Prepare("select*from cart where userID=?")
	if err != nil {
		log.Printf("when prepare error:%v ", err)
		return
	}
	row, err := stmt.Query(userID)
	if err != nil {
		log.Printf("when query error:%v ", err)
		return
	}
	defer row.Close()
	if err = row.Err(); err != nil {
		return
	}
	for row.Next() {
		err = row.Scan(&s.CartID, &s.ProductID, &s.UserID, &s.Amount, &s.Check)
		if err != nil {
			log.Printf("when scan error:%v ", err)
			return
		}
		temporary := model.ShoppingCart{
			CartID:    s.CartID,
			ProductID: s.ProductID,
			UserID:    s.UserID,
			Amount:    s.Amount,
			Check:     s.Check,
		}
		cart = append(cart, &temporary)
	}
	return
}
func UpdateProductsNumber(productID, number int) (err error) {
	_, err = DB.Exec("update product set amount=? where productID=? ", number, productID)
	if err != nil {
		log.Printf("when update products number error:%v ", err)
		return
	}
	return
}
