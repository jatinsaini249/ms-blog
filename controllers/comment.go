package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jatinsaini249/ms-blog/models"
	"github.com/jatinsaini249/ms-blog/services"
)

// ICommentController => interface
type ICommentController interface {
	HandleGetPostComments(gin *gin.Context)
	HandleGetComment(gin *gin.Context)
	HandleSaveComment(gin *gin.Context)
	HandleDeleteComment(gin *gin.Context)
	HandleUpdateComment(gin *gin.Context)
}

// CommentController => Comment Controller
type CommentController struct {
	CommentService services.ICommentService
}

// NewCommentController => Function to initialize comment controller with all its dependencies
func NewCommentController(commentService services.ICommentService) ICommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

// HandleGetPostComments => Method to get all comments on a particular post
func (ctrl *CommentController) HandleGetPostComments(gin *gin.Context) {
	postId := gin.Param("id")
	if postId != "" {
		comments, err := ctrl.CommentService.GetAllCommentsByPostId(postId)
		if err != nil {
			gin.JSON(http.StatusBadRequest, err.Error())
			return
		}
		gin.JSON(http.StatusOK, comments)
		return
	}

	gin.JSON(http.StatusNotFound, fmt.Errorf("Incorrect resource name"))
}

// HandleGetComment => Method handler to get a particualr comment by passing comment id in request
func (ctrl *CommentController) HandleGetComment(gin *gin.Context) {
	commentId := gin.Param("id")
	if commentId != "" {
		comment, err := ctrl.CommentService.GetCommentById(commentId)
		if err != nil {
			gin.JSON(http.StatusBadRequest, err.Error())
			return
		}
		gin.JSON(http.StatusOK, comment)
	}
	gin.JSON(http.StatusNotFound, fmt.Errorf("Incorrect resource name"))
}

// HandleSaveComment => Method handler to save new comment on a particular post by passing postId
func (ctrl *CommentController) HandleSaveComment(gin *gin.Context) {
	var comment *models.Comment

	err := gin.ShouldBindJSON(&comment)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(comment); err != nil {
		gin.JSON(http.StatusBadRequest, gin.Error(fmt.Errorf("Invalid request data")))
		return
	}

	postId := gin.Param("id")
	parentCommentId := gin.Param("parentId")

	commentList, err := ctrl.CommentService.SaveComment(comment, parentCommentId, postId)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	gin.JSON(http.StatusOK, commentList)
}

// HandleUpdateComment => Method handler to update a particular comment by passing comment id
func (ctrl *CommentController) HandleUpdateComment(gin *gin.Context) {
	var comment *models.Comment

	err := gin.ShouldBind(&comment)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(comment); err != nil {
		gin.JSON(http.StatusBadRequest, gin.Error(fmt.Errorf("Invalid request data")))
		return
	}

	id := gin.Param("id")

	res, err := ctrl.CommentService.UpdateCommentById(id, comment)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	gin.JSON(http.StatusOK, res)
}

// HandleDeleteComment => Method to delete a particular comment by passing id in request
func (ctrl *CommentController) HandleDeleteComment(gin *gin.Context) {
	id := gin.Param("id")

	success, err := ctrl.CommentService.DeleteCommentById(id)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if success {
		gin.JSON(http.StatusOK, fmt.Sprintf("Record of Id %v deleted successfully", id))
	}
}
