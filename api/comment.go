package api

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wintervacation/model"
	"wintervacation/service"
	"wintervacation/util"
)

func Writer(c *gin.Context) {
	content := c.PostForm("content")
	userID, _ := strconv.Atoi(c.PostForm("userID"))
	productID, _ := strconv.Atoi(c.PostForm("productID"))
	parentID, _ := strconv.Atoi(c.PostForm("parentID"))
	if content == "" || userID == 0 || productID == 0 {
		util.ResponseParaError(c)
		return
	}
	err := service.AddComment(model.Comment{
		ParentID:  parentID,
		UserID:    userID,
		ProductID: productID,
		Content:   content,
	})
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
}

func DeleteComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.PostForm("commentID")) //获取评论的ID
	content := "该评论已删除"
	if commentID == 0 {
		util.ResponseParaError(c)
		return
	}
	err := service.DeleteComment(commentID, content)
	if err != nil {
		util.ResponseParaError(c)
		return
	}
	util.ResponseOK(c)
}

func LookComment(c *gin.Context) {
}
