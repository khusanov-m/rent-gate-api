package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khusanov-m/rent-gate-api/models"
	"github.com/khusanov-m/rent-gate-api/utils"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(DB *gorm.DB) PostController {
	return PostController{DB}
}

// [...] Create Post Handler
func (pc *PostController) CreatePost(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreatePostInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPost := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    currentUser.ID,
	}

	result := pc.DB.Create(&newPost)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Post with that title already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPost})
}

// [...] Update Post Handler
func (pc *PostController) UpdatePostInput(ctx *gin.Context) {
	postId := ctx.Param("postId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePostInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var post models.Post
	result := pc.DB.First(&post, "uuid = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	if post.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to update this post"})
		return
	}

	var updatedPost models.Post
	pc.DB.First(&updatedPost, "uuid = ?", postId)

	now := time.Now()
	postToUpdate := models.Post{
		Title:     payload.Title,
		Content:   payload.Content,
		Image:     payload.Image,
		CreatedAt: updatedPost.CreatedAt,
		UpdatedAt: now,
		UserID:    currentUser.ID,
	}

	pc.DB.Model(&updatedPost).Updates(postToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPost})
}

// [...] Get Single Post Handler
func (pc *PostController) FindPostById(ctx *gin.Context) {
	postId := ctx.Param("postId")

	var post models.Post
	// .Preload("User")
	result := pc.DB.First(&post, "uuid = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	postResponse := models.PostResponse{
		UUID:    post.UUID,
		Title:   post.Title,
		Content: post.Content,
		Image:   post.Image,
		// User:    post.User,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": postResponse})
}

// [...] Get All Posts Handler
func (pc *PostController) FindPosts(ctx *gin.Context) {
	pagination, err := utils.NewPaginationFromQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var posts []models.Post
	results := pc.DB.Limit(pagination.Limit).Offset(pagination.Offset).Find(&posts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	var postsResponse []models.PostResponse = make([]models.PostResponse, len(posts))
	for i, post := range posts {
		postsResponse[i] = models.PostResponse{
			UUID:    post.UUID,
			Title:   post.Title,
			Content: post.Content,
			Image:   post.Image,
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(posts), "data": postsResponse})
}

// [...] Delete Post Handler
func (pc *PostController) DeletePost(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	postId := ctx.Param("postId")

	var post models.Post
	result := pc.DB.First(&post, "uuid = ?", postId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No post with that title exists"})
		return
	}

	if post.UserID != currentUser.ID {
		ctx.JSON(http.StatusForbidden, gin.H{"status": "fail", "message": "You are not authorized to delete this post"})
		return
	}

	pc.DB.Delete(&models.Post{}, "uuid = ?", postId)

	ctx.JSON(http.StatusNoContent, nil)
}
