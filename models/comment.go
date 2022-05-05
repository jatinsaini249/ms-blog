package models

type Comment struct {
	Id       string  `json:"id"`
	Author   string  `json:"author" validate:"required"`
	Content  string  `json:"content"  validate:"required"`
	ParentId *string `json:"parentId"`
	PostId   string  `json:"postId"`
}
