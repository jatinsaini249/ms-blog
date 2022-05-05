package testhelper

import (
	"time"

	"github.com/jatinsaini249/ms-blog/models"
)

// PostServiceMock => PostService Mock
type PostServiceMock struct {
	Blogs []*models.Post
}

var postTestData = []*models.Post{
	&models.Post{
		Id:        "1",
		Title:     "Test Title",
		Author:    "Test Author",
		Content:   "Test Content",
		TimeStamp: time.Now(),
	},
}

func (postService *PostServiceMock) GetAllBlogsList() ([]*models.Post, error) {
	return postTestData, nil
}

func (postService *PostServiceMock) GetBlogById(id string) (*models.Post, error) {
	return &models.Post{
		Id:        id,
		Title:     "Test Title",
		Author:    "Test Author",
		Content:   "Test Content",
		TimeStamp: time.Now(),
	}, nil
}

func (postService *PostServiceMock) FindBlogById(id string) (*models.Post, int, error) {
	return &models.Post{
		Id:        id,
		Title:     "Test Title",
		Author:    "Test Author",
		Content:   "Test Content",
		TimeStamp: time.Now(),
	}, 1, nil
}

func (postService *PostServiceMock) SaveBlogPosts(blogs []*models.Post) ([]*models.Post, error) {
	return []*models.Post{}, nil
}

func (postService *PostServiceMock) UpdateBlogPost(id string, blog *models.Post) (*models.Post, error) {
	return &models.Post{}, nil
}

func (postService *PostServiceMock) DeleteBlogPost(id string) (bool, error) {
	return true, nil
}
