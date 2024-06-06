package services_test

import (
	"rest_go/app/models"
	"rest_go/app/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShowAllPost(t *testing.T) {
	mockQuery := new(MockQueryHelperInterface)

	postService := services.PostService{Query: mockQuery}
	data, err, meta := postService.ShowAllPost()

	if err != nil {
		t.Errorf("Error was not expected, got %v", err)
	}

	assert.Equal(t, 3, meta.TotalItems)
	assert.Equal(t, data, posts)
}

func TestCreatePost(t *testing.T) {
	mockQuery := new(MockQueryHelperInterface)

	postService := services.PostService{Query: mockQuery}
	data, err := postService.CreatePost(post)

	if err != nil {
		t.Errorf("Error was not expected, got %v", err)
	}

	assert.Equal(t, post, data)
}

func TestGetPostByID(t *testing.T) {
	mockQuery := new(MockQueryHelperInterface)

	postService := services.PostService{Query: mockQuery}
	data, err := postService.GetPostByID(1)

	if err != nil {
		t.Errorf("Error was not expected, got %v", err)
	}

	assert.Equal(t, post, data)
}

func TestUpdatePost(t *testing.T) {
	mockQuery := new(MockQueryHelperInterface)

	postService := services.PostService{Query: mockQuery}

	data, err := postService.UpdatePost(1, post)

	if err != nil {
		t.Errorf("Error was not expected, got %v", err)
	}

	assert.Equal(t, post, data)
}

func TestDeletePost(t *testing.T) {
	mockQuery := new(MockQueryHelperInterface)

	postService := services.PostService{Query: mockQuery}

	err := postService.DeletePost(1)

	if err != nil {
		t.Errorf("Error was not expected, got %v", err)
	}

	assert.Nil(t, err)
}


type MockQueryHelperInterface struct {
	mock.Mock
}

var posts = []models.Post{
	{Title: "Title 1", Content: "Content 1"},
	{Title: "Title 2", Content: "Content 2"},
	{Title: "Title 3", Content: "Content 3"},
}

var post = models.Post{Title: "Title 1", Content: "Content 1"}

func (m *MockQueryHelperInterface) Count() (int, error) {
	return 3, nil
}


func (m *MockQueryHelperInterface) Create(data models.Post) (*models.Post, error) {
	return &post, nil
}


func (m *MockQueryHelperInterface) DeleteByID(ID uint) error {
	return nil
}


func (m *MockQueryHelperInterface) DeleteWhere(query interface{}, args ...interface{}) error {
	panic("unimplemented")
}

func (m *MockQueryHelperInterface) FindAll() ([]models.Post, error) {
	return posts, nil
}


func (m *MockQueryHelperInterface) FindAllWithPagination(page int, size int) ([]models.Post, error) {
	panic("unimplemented")
}


func (m *MockQueryHelperInterface) FindByID(ID uint) (*models.Post, error) {
	return &post, nil
}


func (m *MockQueryHelperInterface) FindManyByColumn(column string, value interface{}) ([]models.Post, error) {
	panic("unimplemented")
}


func (m *MockQueryHelperInterface) FindOneByColumn(column string, value interface{}) (*models.Post, error) {
	panic("unimplemented")
}


func (m *MockQueryHelperInterface) FirstWhere(query interface{}, args ...interface{}) (*models.Post, error) {
	panic("unimplemented")
}

func (m *MockQueryHelperInterface) Update(data models.Post) (*models.Post, error) {
	return &post, nil
}

func (m *MockQueryHelperInterface) Where(query interface{}, args ...interface{}) ([]models.Post, error) {
	panic("unimplemented")
}
