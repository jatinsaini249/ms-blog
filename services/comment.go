package services

import (
	"fmt"
	"strings"

	"github.com/jatinsaini249/ms-blog/models"
)

// ICommentService => interface
type ICommentService interface {
	SaveComment(comment *models.Comment, parentId string, postId string) (*models.Comment, error)
	GetAllCommentsByPostId(id string) ([]*models.Comment, error)
	GetAllCommentsByParentCommentId(id string) ([]*models.Comment, error)
	GetCommentById(id string) ([]*models.Comment, error)
	UpdateCommentById(id string, blog *models.Comment) (*models.Comment, error)
	DeleteCommentById(id string) (bool, error)
}

// commentList => In memory list of comments
var commentList []*models.Comment = []*models.Comment{
	&models.Comment{
		Id:      "1",
		Author:  "Rob Pike",
		Content: "This is a bug",
		PostId:  "1",
	},
}

// CommentService  => struct
type CommentService struct {
	Comments []*models.Comment
}

// NewCommentService => Initializes comment service with all its dependencies
func NewCommentService() ICommentService {
	return &CommentService{
		Comments: commentList,
	}
}

// GetAllCommentsByPostId => Method to get list of comments by passing post id
func (commentService *CommentService) GetAllCommentsByPostId(id string) ([]*models.Comment, error) {
	commentList := make([]*models.Comment, 0)

	for _, comment := range commentService.Comments {
		if strings.EqualFold(comment.PostId, id) {
			commentList = append(commentList, comment)
		}
	}

	return commentList, nil
}

// GetAllCommentsByParentCommentId => Method to get child comments by passing comment Id in request
func (commentService *CommentService) GetAllCommentsByParentCommentId(id string) ([]*models.Comment, error) {
	commentList := make([]*models.Comment, 0)

	for _, comment := range commentService.Comments {
		if comment.ParentId != nil && strings.EqualFold(*comment.ParentId, id) {
			commentList = append(commentList, comment)
		}
	}

	return commentList, nil
}

// GetCommentById => Method to get comment by Id
func (commentService *CommentService) GetCommentById(id string) ([]*models.Comment, error) {
	commentList := make([]*models.Comment, 0)
	mapList := make(map[string]*models.Comment)
	parentIds := make([]string, 0)
	for _, comment := range commentService.Comments {
		if strings.EqualFold(comment.Id, id) {
			commentList = append(commentList, comment)
			parentIds = append(parentIds, id)
		}
		if comment.ParentId != nil {
			mapList[*comment.ParentId] = comment
		}
	}

	for len(parentIds) > 0 {
		parentCommentIds := make([]string, 0)
		for index, id := range parentIds {
			val, ok := mapList[id]
			if ok {
				commentList = append(commentList, val)
				if val.ParentId != nil {
					parentCommentIds = append(parentIds, *val.ParentId)
				}
			}
			parentCommentIds = append(parentIds[:index], parentIds[(index+1):]...)
		}
		parentIds = parentCommentIds
	}

	if len(commentList) > 0 {
		return commentList, nil
	}

	return nil, fmt.Errorf("No record found for this comment Id : %v", id)
}

// SaveComment => Method to save comment by passing postId and parent Id if it is a child comment
func (commentService *CommentService) SaveComment(comment *models.Comment, parentId string, postId string) (*models.Comment, error) {
	comment.ParentId = &parentId
	comment.PostId = postId
	comment.Id = CreateWithoutHashGUID()

	commentService.Comments = append(commentService.Comments, comment)
	return comment, nil

}

// UpdateCommentById => Method to update comment by comment Id
func (commentService *CommentService) UpdateCommentById(id string, comment *models.Comment) (*models.Comment, error) {
	for index, val := range commentService.Comments {
		if strings.EqualFold(val.Id, id) {
			commentService.Comments[index].Author = comment.Author
			commentService.Comments[index].Content = comment.Content
			return commentService.Comments[index], nil
		}
	}
	return nil, fmt.Errorf("No record found for comment Id : %v", id)
}

//DeleteCommentById => Method to delete comment by comment Id
func (commentService *CommentService) DeleteCommentById(id string) (bool, error) {
	for index, comment := range commentService.Comments {
		if strings.EqualFold(comment.Id, id) {
			commentService.Comments = append(commentService.Comments[:index], commentService.Comments[(index+1):]...)
			return true, nil
		}
	}
	return false, fmt.Errorf("No record found for comment Id : %v", id)
}
