package models

// Comment => struct
type Comment struct {
	Id       string  `json:"id"`
	Author   string  `json:"author" validate:"required"`
	Content  string  `json:"content"  validate:"required"`
	ParentId *string `json:"parentId,omitempty"`
	PostId   string  `json:"postId"`
}
