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
		util.ResponseInternalError(c)
		return
	}
	util.ResponseOK(c)
}

func LookComment(c *gin.Context) {

	productID, _ := strconv.Atoi(c.PostForm("productID"))
	if productID == 0 {
		util.ResponseParaError(c)
		return
	}

	//多级评论：遍历最上层的评论，依次找到每个上层评论的下级评论
	var t model.Comment

	//得到该product的评论
	err, allComments := service.SearchComments(productID, &t)
	if err != nil {
		util.ResponseInternalError(c)
		return
	}
	//将扁平数据转化为树型数据
	var tree []*model.Comment
	IdMapTree := make(map[int]*model.Comment) //建立map表，用于收集所有的节点

	for _, item := range allComments {

		//收集第一层评论
		if item.ParentID == 0 {
			tree = append(tree, item)
		} else {
			//收集子评论
			IdMapTree[item.ParentID].Children = append(IdMapTree[item.ParentID].Children, item)
		}
		//把节点加入map表中
		IdMapTree[item.CommentID] = item

	}

	util.ResponseComments(c, IdMapTree)

}
