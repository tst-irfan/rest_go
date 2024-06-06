
package services

import (
	"rest_go/app/models"
	"rest_go/app/types"
	"sync"
	"rest_go/db"
)

type PostService struct {
	Query db.QueryHelperInterface[models.Post]
}

type PostServiceInterface interface {
	ShowAllPost() ([]models.Post, error, types.MetaData)
	CreatePost(models.Post) (models.Post, error)
	GetPostByID(uint) (models.Post, error)
	UpdatePost(uint, models.Post) (models.Post, error)
	DeletePost(uint) error
}

func NewPostService() *PostService {
	return &PostService{Query: &models.PostQuery}
}

func (s *PostService) ShowAllPost() ([]models.Post, error, types.MetaData) {
	var wg sync.WaitGroup

	var posts []models.Post
	var totalItems int
	var err error

	wg.Add(2)
	go func() {
		defer wg.Done()
		posts, err = s.Query.FindAll()
	}()
	go func() {
		defer wg.Done()
		totalItems, err = s.Query.Count()
	}()
	
	wg.Wait()

	if err != nil {
		return []models.Post{}, err, types.MetaData{}
	}

	metaData := types.MetaData{
		TotalItems: totalItems,
	}

	return posts, nil, metaData
}

func (s *PostService) CreatePost(post models.Post) (models.Post, error) {
	createdPost, err := s.Query.Create(post)
	if err != nil {
		return models.Post{}, err
	}

	return *createdPost, nil
}

func (s *PostService) GetPostByID(ID uint) (models.Post, error) {
	post, err := s.Query.FindByID(ID)
	if err != nil {
		return models.Post{}, err
	}

	return *post, nil
}

func (s *PostService) UpdatePost(ID uint, postParams models.Post) (models.Post, error) {
	post, err := s.Query.FindByID(ID)
	if err != nil {
		return models.Post{}, err
	}

	postParams.ID = post.ID

	updatedPost, err := s.Query.Update(postParams)
	if err != nil {
		return models.Post{}, err
	}

	return *updatedPost, nil
}

func (s *PostService) DeletePost(ID uint) error {
	err := s.Query.DeleteByID(ID)
	if err != nil {
		return err
	}

	return nil
}
