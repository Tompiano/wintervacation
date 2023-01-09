package main

import (
	"wintervacation/api"
	"wintervacation/dao"
)

func main() {
	dao.InitDB()
	api.Entrance()
}
