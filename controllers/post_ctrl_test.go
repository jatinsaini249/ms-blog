package controllers

import (
	"testing"

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
