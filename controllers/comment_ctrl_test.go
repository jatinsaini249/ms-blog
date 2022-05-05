package controllers

import (
	"testing"

	"github.com/jatinsaini249/ms-blog/models"
	"github.com/jatinsaini249/ms-blog/testhelper"
	"github.com/stretchr/testify/assert"
)

func TestCommentController_NewCommentController(testRunner *testing.T) {
	testRunner.Run("NewCommentController should return valid response", func(testRunnerChild *testing.T) {

		commentController := NewCommentController(&testhelper.CommentServiceMock{})
		assert.NotNil(testRunnerChild, commentController, "CommentController can't be NIL object")
		assert.NotNil(testRunnerChild, commentController.CommentService, "CommentService can't be NIL object")
	})
}

func TestCommentController_HandleGetPostComments(testRunner *testing.T) {
	testRunner.Run("HandleGetBlogPosts should return valid response", func(testRunnerChild *testing.T) {
		commentController := NewCommentController(&testhelper.CommentServiceMock{})
		posts, err := commentController.CommentService.GetAllCommentsByPostId("1")
		assert.NotNil(testRunnerChild, posts, "posts can't be empty")
		assert.Nil(testRunnerChild, err, "Error should be nil")
	})
}

func TestCommentController_HandleGetComment(testRunner *testing.T) {
	testRunner.Run("HandleGetComment should return valid response", func(testRunnerChild *testing.T) {
		commentController := NewCommentController(&testhelper.CommentServiceMock{})
		comments, err := commentController.CommentService.GetAllCommentsByParentCommentId("1")
		assert.NotNil(testRunnerChild, comments, "post can't be nil")
		assert.True(testRunnerChild, len(comments) > 0, "length of comments should be greater than 0")
		assert.Nil(testRunnerChild, err, "error should be nil")
	})
}

func TestCommentController_HandleSaveComment(testRunner *testing.T) {
	testRunner.Run("HandleSaveComment should return valid response", func(testRunnerChild *testing.T) {
		commentController := NewCommentController(&testhelper.CommentServiceMock{})
		comment, err := commentController.CommentService.SaveComment(&models.Comment{Author: "Test", Content: "Test"}, "", "1")
		assert.NotNil(testRunnerChild, comment, "post can't be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
	})
}

func TestCommentController_HandleDeleteComment(testRunner *testing.T) {
	testRunner.Run("HandleDeleteComment should return valid response", func(testRunnerChild *testing.T) {
		commentController := NewCommentController(&testhelper.CommentServiceMock{})
		success, err := commentController.CommentService.DeleteCommentById("1234567")
		assert.True(testRunnerChild, success, "post can't be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
	})
}

func TestCommentController_HandleUpdateComment(testRunner *testing.T) {
	testRunner.Run("HandleUpdateComment should return valid response", func(testRunnerChild *testing.T) {
		commentController := NewCommentController(&testhelper.CommentServiceMock{})
		post, err := commentController.CommentService.UpdateCommentById("1", &models.Comment{Author: "Test", Content: "Test"})
		assert.NotNil(testRunnerChild, post, "post can't be nil")
		assert.Nil(testRunnerChild, err, "error should be nil")
	})
}
