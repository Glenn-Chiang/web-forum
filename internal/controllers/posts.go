package controllers

import (
	"cvwo-backend/internal/models"
	"cvwo-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	service services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{service}
}

// GET /posts or /posts?topic_id=1
func (controller *PostController) GetAll(ctx *gin.Context) {
	topicIdParam := ctx.Query("topic_id")

	var posts []models.Post
	var err error

	if topicIdParam != "" { // If topicId is specified
		// Parse topic_id from request query param
		topicID, convErr := strconv.Atoi(topicIdParam)
		if convErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid topic_id"})
			return
		}
		posts, err = controller.service.GetByTopic(uint(topicID))
	} else { // If no topicId is specified
		posts, err = controller.service.GetAll()
	}
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, posts)
}

// GET /posts/:id
func (controller *PostController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
	}

	post, err := controller.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, post)
}

// POST /posts
func (controller *PostController) Create(ctx *gin.Context) {
	// Validate request body
	var requestBody models.CreatePostRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map fields from request body to Post model
	post := models.Post{
		Title: requestBody.Title,
		Content: requestBody.Content,
		AuthorID: requestBody.AuthorID,
	}

	newPost, err := controller.service.Create(&post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.IndentedJSON(http.StatusCreated, newPost)
}

// DELETE /posts/:id
func (controller *PostController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})		
	}

	if controller.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
