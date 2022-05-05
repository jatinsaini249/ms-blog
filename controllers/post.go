package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jatinsaini249/ms-blog/models"
	"github.com/jatinsaini249/ms-blog/services"
)

type IPostController interface {
	HandleGetBlogPosts(gin *gin.Context)
	HandleSaveBlogs(gin *gin.Context)
	HandleDeleteBlogPost(gin *gin.Context)
	HandleUpdateBlogPost(gin *gin.Context)
}

type PostController struct {
	BlogPostService services.IPostService
}

func NewPostController(postService services.IPostService) *PostController {
	return &PostController{
		BlogPostService: postService,
	}
}

func (ctrl *PostController) HandleGetBlogPosts(gin *gin.Context) {
	blogsList, err := ctrl.BlogPostService.GetAllBlogsList()
	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}
	gin.JSON(http.StatusOK, blogsList)
}

func (ctrl *PostController) HandleSaveBlogs(gin *gin.Context) {
	var blogPost []*models.Post

	err := gin.ShouldBindJSON(&blogPost)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	validate := validator.New()
	if err := validate.Struct(blogPost); err != nil {
		gin.JSON(http.StatusBadRequest, gin.Error(fmt.Errorf("Invalid request data")))
		return
	}

	blogs, err := ctrl.BlogPostService.SaveBlogPosts(blogPost)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}

	gin.JSON(http.StatusOK, blogs)
}

func (ctrl *PostController) HandleUpdateBlogPost(gin *gin.Context) {
	var blogPost *models.Post

	err := gin.ShouldBind(&blogPost)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(blogPost); err != nil {
		gin.JSON(http.StatusBadRequest, gin.Error(fmt.Errorf("Invalid request data")))
		return
	}

	id := gin.Param("id")

	blog, err := ctrl.BlogPostService.UpdateBlogPost(id, blogPost)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	gin.JSON(http.StatusOK, blog)
}

func (ctrl *PostController) HandleDeleteBlogPost(gin *gin.Context) {
	id := gin.Param("id")

	success, err := ctrl.BlogPostService.DeleteBlogPost(id)
	if err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if success {
		gin.JSON(http.StatusOK, fmt.Sprintf("Record of Id %v deleted successfully", id))
	}
}
