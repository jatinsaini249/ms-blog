package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jatinsaini249/ms-blog/models"
)

var blogs []*models.Post = []*models.Post{
	&models.Post{
		Id:        "1",
		Title:     "Random Title",
		Author:    "Benjamin Stroustrup",
		Content:   "This is a random Post",
		TimeStamp: time.Now(),
	},
}

type IPostService interface {
	SaveBlogPosts(blogs []*models.Post) ([]*models.Post, error)
	GetAllBlogsList() ([]*models.Post, error)
	GetBlogById(id string) (*models.Post, error)
	FindBlogById(id string) (*models.Post, int, error)
	UpdateBlogPost(id string, blog *models.Post) (*models.Post, error)
	DeleteBlogPost(id string) (bool, error)
}

type PostService struct {
	Blogs []*models.Post
}

func NewPostService() IPostService {
	return &PostService{
		Blogs: blogs,
	}
}

func (postService *PostService) GetAllBlogsList() ([]*models.Post, error) {
	if len(postService.Blogs) < 1 {
		return nil, errors.New("Empty Blog List")
	}
	return postService.Blogs, nil
}

func (postService *PostService) GetBlogById(id string) (*models.Post, error) {
	blog, _, err := postService.FindBlogById(id)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (postService *PostService) SaveBlogPosts(blogs []*models.Post) ([]*models.Post, error) {
	blogList := make([]*models.Post, 0)
	for _, post := range blogs {
		post.Id = CreateWithoutHashGUID()
		blogList = append(blogList, post)
	}

	postService.Blogs = append(postService.Blogs, blogList...)
	return blogList, nil
}

func (postService *PostService) UpdateBlogPost(id string, blog *models.Post) (*models.Post, error) {
	blog, index, err := postService.FindBlogById(id)
	if err != nil {
		return nil, err
	}
	postService.Blogs[index] = patchBlogPost(postService.Blogs[index], blog)
	return postService.Blogs[index], nil
}

func (postService *PostService) DeleteBlogPost(id string) (bool, error) {
	_, index, err := postService.FindBlogById(id)
	if err != nil {
		return false, err
	}
	postService.Blogs = append(postService.Blogs[:index], postService.Blogs[index+1:]...)
	return true, nil
}

func (PostService *PostService) FindBlogById(id string) (*models.Post, int, error) {
	for index, post := range PostService.Blogs {
		if strings.EqualFold(post.Id, id) {
			return post, index, nil
		}
	}
	return nil, -1, fmt.Errorf("No record found for Id : %v", id)
}

func CreateWithoutHashGUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

func patchBlogPost(blogPost *models.Post, blog *models.Post) *models.Post {
	blogPost.Author = blog.Author
	blogPost.Content = blog.Content
	blogPost.TimeStamp = blog.TimeStamp
	blogPost.Title = blog.Title
	return blogPost
}
