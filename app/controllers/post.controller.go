
package controllers

import (
	"net/http"
	"rest_go/app/helpers"
	"rest_go/app/services"
	"rest_go/app/types"
	"rest_go/app/models"
	"rest_go/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Service services.PostServiceInterface
}

func NewPostController() PostController {
	return PostController{Service: services.NewPostService()}
}

// Show All Posts godoc
// @Summary Show all posts
// @Description get all posts
// @Tags Post
// @Accept  json
// @Produce  json
// @Success 200 {object} types.SuccessWithMeta[[]models.Post]
// @Failure 400 {object} types.Error
// @Router /posts [get]
// @Security Bearer
func (con *PostController) GetAllPosts(c *gin.Context) {
	var posts []models.Post
	var err error
	var metadata types.MetaData

	posts, err, metadata = con.Service.ShowAllPost()

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccessWithMeta(c, "Post found", posts, http.StatusOK, metadata)
}

// Get Post godoc
// @Summary Get a post
// @Description get a post by id
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} types.Success[models.Post]
// @Failure 400 {object} types.Error
// @Router /posts/{id} [get]
// @Security Bearer
func (con *PostController) GetPost(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := con.Service.GetPostByID(uint(ID))

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Post found", post, http.StatusOK)
}

// Create Post godoc
// @Summary Create a post
// @Description create a post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param input body types.PostRequest true "User input"
// @Success 201 {object} types.Success[models.Post]
// @Failure 400 {object} types.Error
// @Router /posts [post]
// @Security Bearer
func (con *PostController) CreatePost(c *gin.Context) {
	var input types.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}
	postParams, err := utils.TypeConverter[models.Post](&input)

	post, err := con.Service.CreatePost(*postParams)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Post created", post, http.StatusCreated)
}

// Update Post godoc
// @Summary Update a post
// @Description update a post by id
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Param input body types.PostRequest true "User input"
// @Success 200 {object} types.Success[models.Post]
// @Failure 400 {object} types.Error
// @Router /posts/{id} [put]
// @Security Bearer
func (con *PostController) UpdatePost(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	var input types.PostRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	postParams, err := utils.TypeConverter[models.Post](&input)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := con.Service.UpdatePost(uint(ID), *postParams)

	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Post updated", post, http.StatusOK)
}

// Delete Post godoc
// @Summary Delete a post
// @Description delete a post by id
// @Tags Post
// @Accept  json
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {string} types.Success
// @Failure 400 {object} types.Error
// @Router /posts/{id} [delete]
// @Security Bearer
func (con *PostController) DeletePost(c *gin.Context) {
	IDStr := c.Param("id")
	ID, err := strconv.ParseUint(IDStr, 10, 64)
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	err = con.Service.DeletePost(uint(ID))
	if err != nil {
		helpers.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	helpers.ResponseSuccess(c, "Post deleted", "", http.StatusOK)
}
