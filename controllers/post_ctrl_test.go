package controllers

import (
	"testing"

	"github.com/jatinsaini249/ms-blog/models"
	"github.com/jatinsaini249/ms-blog/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestPostController_NewPostController(testRunner *testing.T) {
	testRunner.Run("NewPostController should return valid response", func(testRunnerChild *testing.T) {

		postController := NewPostController(&testhelper.PostServiceMock{})
		assert.NotNil(testRunnerChild, postController, "PostController can't be NIL object")
		assert.NotNil(testRunnerChild, postController.BlogPostService, "BlogPostService can't be NIL object")
	})
}

func TestPostController_HandleGetBlogPosts(testRunner *testing.T) {
	testRunner.Run("HandleGetBlogPosts should return valid response", func(testRunnerChild *testing.T) {
		postController := NewPostController(&testhelper.PostServiceMock{})
		posts, err := postController.BlogPostService.GetAllBlogsList()
		assert.NotEmpty(testRunnerChild, posts, "posts can't be empty")
		assert.Nil(testRunnerChild, err, "Error should be nil")
	})
}

func TestPostController_HandleSaveBlogs(testRunner *testing.T) {
	testRunner.Run("HandleSaveBlogs should return valid response", func(testRunnerChild *testing.T) {
		postController := NewPostController(&testhelper.PostServiceMock{})
		post, err := postController.BlogPostService.SaveBlogPosts(&models.Post{})
		assert.NotNil(testRunnerChild, post, "post can't be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
	})
}

func TestPostController_HandleUpdateBlogPost(testRunner *testing.T) {
	testRunner.Run("HandleUpdateBlogPost should return valid response", func(testRunnerChild *testing.T) {
		postController := NewPostController(&testhelper.PostServiceMock{})
		post, err := postController.BlogPostService.UpdateBlogPost("1234567", &models.Post{})
		assert.NotNil(testRunnerChild, post, "post can't be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
	})
}

func TestPostController_HandleDeleteBlogPost(testRunner *testing.T) {
	testRunner.Run("HandleDeleteBlogPost should return valid response", func(testRunnerChild *testing.T) {
		postController := NewPostController(&testhelper.PostServiceMock{})
		success, err := postController.BlogPostService.DeleteBlogPost("1234567")
		assert.True(testRunnerChild, success, "post can't be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
	})
}
