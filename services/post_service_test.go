package services

import (
	"testing"

	"github.com/jatinsaini249/ms-blog/models"
	"github.com/stretchr/testify/assert"
)

func TestPostService_SaveBlogPosts(testRunner *testing.T) {

	testRunner.Run("SaveBlogPosts should return valid response", func(testRunnerChild *testing.T) {
		postService := NewPostService()

		post, err := postService.SaveBlogPosts(&models.Post{Title: "TestTitle", Author: "TestAuthor", Content: "TestContent"})
		assert.NotNil(testRunnerChild, post, "post should not be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
		assert.Equal(testRunnerChild, post.Title, "TestTitle", "title property should be same that we pass in request")
	})
}

func TestPostService_GetAllBlogsList(testRunner *testing.T) {
	testRunner.Run("GetAllBlogList should return valid response", func(testRunnerChild *testing.T) {
		postService := NewPostService()
		post := &models.Post{Title: "TestTitle", Author: "TestAuthor", Content: "TestContent"}
		postService.Blogs = append(postService.Blogs, post)
		posts, err := postService.GetAllBlogsList()

		assert.NotNil(testRunnerChild, posts, "posts should not be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
		assert.True(testRunnerChild, len(posts) > 0, "Length of posts should be greater than 0")
	})
}

func TestPostService_GetBlogById(testRunner *testing.T) {
	testRunner.Run("GetBlogById should return valid response", func(testRunnerChild *testing.T) {
		postService := NewPostService()
		post := &models.Post{Id: "1", Title: "TestTitle", Author: "TestAuthor", Content: "TestContent"}
		postService.Blogs = append(postService.Blogs, post)
		res, err := postService.GetBlogById("1")

		assert.NotNil(testRunnerChild, res, "post should not be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
		assert.Equal(testRunnerChild, res.Id, "1", "Id should be equal to 1")
	})
}

func TestPostService_UpdateBlogPost(testRunner *testing.T) {
	testRunner.Run("UpdateBlogPost should return valid response", func(testRunnerChild *testing.T) {
		postService := NewPostService()
		post := &models.Post{Title: "TestTitle", Author: "TestAuthor", Content: "TestContent"}
		postService.Blogs = append(postService.Blogs, post)
		res, err := postService.UpdateBlogPost("1", post)

		assert.NotNil(testRunnerChild, res, "post should not be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
		assert.Equal(testRunnerChild, res.Title, "TestTitle", "Title should be equal to TestTitle")
	})
}

func TestPostService_DeleteBlogPost(testRunner *testing.T) {
	testRunner.Run("UpdateBlogPost should return valid response", func(testRunnerChild *testing.T) {
		postService := NewPostService()
		post := &models.Post{Title: "TestTitle", Author: "TestAuthor", Content: "TestContent"}
		postService.Blogs = append(postService.Blogs, post)
		success, err := postService.DeleteBlogPost("1")

		assert.NotNil(testRunnerChild, success, "success should not be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
		assert.True(testRunnerChild, success, "Success should be truthy value")
	})
}
