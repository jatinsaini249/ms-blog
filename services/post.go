package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jatinsaini249/ms-blog/models"
)

// In memory storage for blogs
var blogs []*models.Post = []*models.Post{
	&models.Post{
		Id:        "1",
		Title:     "Random Title",
		Author:    "Benjamin Stroustrup",
		Content:   "This is a random Post",
		TimeStamp: time.Now().Format("2006-01-02"),
	},
}

// IPostService => Interface
type IPostService interface {
	SaveBlogPosts(blogs *models.Post) (*models.Post, error)
	GetAllBlogsList() ([]*models.Post, error)
	GetBlogById(id string) (*models.Post, error)
	FindBlogById(id string) (*models.Post, int, error)
	UpdateBlogPost(id string, blog *models.Post) (*models.Post, error)
	DeleteBlogPost(id string) (bool, error)
}

// PostService => struct
type PostService struct {
	Blogs []*models.Post
}

// NewPostService => Function to initialize Post service with all its dependencies
func NewPostService() IPostService {
	return &PostService{
		Blogs: blogs,
	}
}

// GetAllBlogsList => Method to return all blog posts
func (postService *PostService) GetAllBlogsList() ([]*models.Post, error) {
	if len(postService.Blogs) < 1 {
		return nil, errors.New("Empty Blog List")
	}
	return postService.Blogs, nil
}

// GetBlogById => Method to return particular blog post by passing Id in request
func (postService *PostService) GetBlogById(id string) (*models.Post, error) {
	blog, _, err := postService.FindBlogById(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

// SaveBlogPosts => Method to insert new blog post
func (postService *PostService) SaveBlogPosts(blog *models.Post) (*models.Post, error) {
	if blog != nil {
		blog.Id = CreateWithoutHashGUID()
		blog.TimeStamp = time.Now().Format("2006-01-02")
	}

	postService.Blogs = append(postService.Blogs, blog)
	return blog, nil
}

// UpdateBlogPost => Method to update existing blog post by passing Id in request
func (postService *PostService) UpdateBlogPost(id string, blog *models.Post) (*models.Post, error) {
	_, index, err := postService.FindBlogById(id)
	if err != nil {
		return nil, err
	}
	postService.Blogs[index] = patchBlogPost(postService.Blogs[index], blog)
	return postService.Blogs[index], nil
}

// DeleteBlogPost => Method to remove a particular blog post by passing Id in request
func (postService *PostService) DeleteBlogPost(id string) (bool, error) {
	_, index, err := postService.FindBlogById(id)
	if err != nil {
		return false, err
	}
	postService.Blogs = append(postService.Blogs[:index], postService.Blogs[index+1:]...)
	return true, nil
}

// FindBlogById => Method to find a particular blog post by passing Id in request
func (PostService *PostService) FindBlogById(id string) (*models.Post, int, error) {
	for index, post := range PostService.Blogs {
		if strings.EqualFold(post.Id, id) {
			return post, index, nil
		}
	}
	return nil, -1, fmt.Errorf("No record found for Id : %v", id)
}

// CreateWithoutHashGUID => To create new guid without hashes to set it as Id
func CreateWithoutHashGUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

// patchBlogPost => To update particular blog post with values passing in request
func patchBlogPost(blogPost *models.Post, blog *models.Post) *models.Post {
	blogPost.Author = blog.Author
	blogPost.Content = blog.Content
	blogPost.Title = blog.Title
	return blogPost
}
