package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"rest_go/app/controllers"
	"rest_go/app/models"
	"rest_go/app/types"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) ShowAllPost() ([]models.Post, error, types.MetaData) {
	args := m.Called()
	return args.Get(0).([]models.Post), args.Error(1), args.Get(2).(types.MetaData)
}

func (m *MockPostService) GetPostByID(ID uint) (models.Post, error) {
	args := m.Called(ID)
	return args.Get(0).(models.Post), args.Error(1)
}

func (m *MockPostService) CreatePost(post models.Post) (models.Post, error) {
	args := m.Called(post)
	return args.Get(0).(models.Post), args.Error(1)
}

func (m *MockPostService) UpdatePost(ID uint, post models.Post) (models.Post, error) {
	args := m.Called(ID, post)
	return args.Get(0).(models.Post), args.Error(1)
}

func (m *MockPostService) DeletePost(ID uint) error {
	args := m.Called(ID)
	return args.Error(0)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func setupController() (*controllers.PostController, *MockPostService) {
	mockService := new(MockPostService)
	controller := &controllers.PostController{
		Service: mockService,
	}
	return controller, mockService
}

func TestGetAllPosts(t *testing.T) {
	r := setupRouter()
	controller, mockService := setupController()

	mockPosts := []models.Post{}
	mockMeta := types.MetaData{TotalItems: 0}
	mockService.On("ShowAllPost").Return(mockPosts, nil, mockMeta)

	r.GET("/posts", controller.GetAllPosts)

	req, _ := http.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetPost(t *testing.T) {
	r := setupRouter()
	controller, mockService := setupController()

	mockPost := models.Post{}
	mockService.On("GetPostByID", uint(1)).Return(mockPost, nil)

	r.GET("/posts/:id", controller.GetPost)

	req, _ := http.NewRequest("GET", "/posts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestCreatePost(t *testing.T) {
	r := setupRouter()
	controller, mockService := setupController()

	mockPost := models.Post{Title: "Test Title", Content: "Test Content"}
	mockService.On("CreatePost", mock.AnythingOfType("models.Post")).Return(mockPost, nil)

	r.POST("/posts", controller.CreatePost)

	jsonStr := `{"title":"Test Title","content":"Test Content"}`
	req, _ := http.NewRequest("POST", "/posts", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	r := setupRouter()
	controller, mockService := setupController()

	mockPost := models.Post{Title: "Updated Title", Content: "Updated Content"}
	mockService.On("UpdatePost", uint(1), mock.AnythingOfType("models.Post")).Return(mockPost, nil)

	r.PUT("/posts/:id", controller.UpdatePost)

	jsonStr := `{"title":"Updated Title","content":"Updated Content"}`
	req, _ := http.NewRequest("PUT", "/posts/1", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeletePost(t *testing.T) {
	r := setupRouter()
	controller, mockService := setupController()

	mockService.On("DeletePost", uint(1)).Return(nil)

	r.DELETE("/posts/:id", controller.DeletePost)

	req, _ := http.NewRequest("DELETE", "/posts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockService.AssertExpectations(t)
}
