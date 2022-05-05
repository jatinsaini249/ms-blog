package testhelper

import (
	"github.com/jatinsaini249/ms-blog/models"
)

// CommentServiceMock => Comment Service Mock
type CommentServiceMock struct {
	Comments []*models.Comment
}

var commentTestData = []*models.Comment{
	&models.Comment{
		Id:      "1",
		Author:  "Test Author",
		Content: "Test Content",
	},
}

func (commentService *CommentServiceMock) GetAllCommentsByPostId(id string) ([]*models.Comment, error) {
	return []*models.Comment{
		&models.Comment{},
		&models.Comment{},
	}, nil
}

func (commentService *CommentServiceMock) GetAllCommentsByParentCommentId(id string) ([]*models.Comment, error) {
	return []*models.Comment{
		&models.Comment{},
		&models.Comment{},
	}, nil
}

func (commentService *CommentServiceMock) GetCommentById(id string) ([]*models.Comment, error) {
	return []*models.Comment{
		&models.Comment{},
		&models.Comment{},
	}, nil
}

func (commentService *CommentServiceMock) SaveComment(comment *models.Comment, parentId string, postId string) (*models.Comment, error) {
	return &models.Comment{}, nil
}

func (commentService *CommentServiceMock) UpdateCommentById(id string, blog *models.Comment) (*models.Comment, error) {
	return &models.Comment{}, nil
}

func (commentService *CommentServiceMock) DeleteCommentById(id string) (bool, error) {
	return true, nil
}
