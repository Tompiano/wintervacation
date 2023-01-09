package api

import "github.com/gin-gonic/gin"

func Entrance() {
	r := gin.Default()
	r.Use(handleWare())
	u := r.Group("/user")
	{
		u.POST("/register", Register)
	}
	r.Run()
}
