package models

// Post  => struct
type Post struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title" validate:"required"`
	Author    string `json:"author" validate:"required"`
	Content   string `json:"content"  validate:"required"`
	TimeStamp string `json:"timeStamp,omitempty"`
}
