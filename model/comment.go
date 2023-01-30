package model

type Comment struct {
	CommentID int        `json:"commentID"`
	ParentID  int        `json:"parentID"`
	UserID    int        `json:"userID"`
	ProductID int        `json:"productID"`
	Content   string     `json:"content"`
	Children  []*Comment `json:"children"`
}
