package controller

import (
	"crud-go/src/entity"
	"crud-go/src/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostController struct {
	PostUsecase usecase.PostUsecase
}

func (ph *PostController) GetPosts(c *gin.Context) {
	posts, err := ph.PostUsecase.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (ph *PostController) GetPostById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	post, err := ph.PostUsecase.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (ph *PostController) CreatePost(c *gin.Context) {
	var request entity.CreatePostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Body request"})
		return
	}

	err := ph.PostUsecase.CreatePost(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (ph *PostController) DeletePost(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	deleteError := ph.PostUsecase.DeletePost(id)
	if deleteError != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": deleteError.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (ph *PostController) UpdatePost(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}
	var request entity.CreatePostRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Body request"})
		return
	}

	post, updateError := ph.PostUsecase.UpdatePost(id, &request)
	if updateError != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": updateError.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}
