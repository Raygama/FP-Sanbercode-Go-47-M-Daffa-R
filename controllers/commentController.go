package controllers

import (
	"net/http"
	"time"

	"Final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentInput struct {
	ReviewID int    `json:"review_id"`
	UserID   int    `json:"user_id"`
	Content  string `json:"content" gorm:"type:text"`
	Likes    int    `json:"likes"`
}

// GetAllComments godoc
// @Summary Get all Comments.
// @Description Get a list of Comments.
// @Tags Comment
// @Produce json
// @Success 200 {object} []models.Comment
// @Router /comments [get]
func GetAllComment(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var comments []models.Comment
	db.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// CreateComment godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags Comment
// @Param Body body commentInput true "the body to create a new Comment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comments [post]
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input commentInput
	var review models.Review
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ReviewID).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReviewID not found!"})
		return
	}

	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	comment := models.Comment{ReviewID: input.ReviewID, UserID: input.UserID, Content: input.Content, Likes: input.Likes}
	db.Create(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// GetCommentById godoc
// @Summary Get Comment.
// @Description Get a Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
func GetCommentById(c *gin.Context) { // Get model if exist
	var comment models.Comment

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// GetCommentsByUserId godoc
// @Summary Get Comments.
// @Description Get all Comments by UserId.
// @Tags User
// @Produce json
// @Param id path string true "User id"
// @Success 200 {object} []models.Comment
// @Router /users/{id}/comments [get]
func GetCommentsByUserId(c *gin.Context) {
	var comments []models.Comment

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("user_id = ?", c.Param("id")).Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

// UpdateComment godoc
// @Summary Update Comment.
// @Description Update Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "Comment id"
// @Param Body body commentInput true "the body to update a Comment"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Comment
// @Router /comments/{id} [patch]
func UpdateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist

	var comment models.Comment
	var review models.Review
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ReviewID).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReviewID not found!"})
		return
	}

	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	var updatedInput models.Comment
	updatedInput.ReviewID = input.ReviewID
	updatedInput.UserID = input.UserID
	updatedInput.Likes = input.Likes
	updatedInput.Content = input.Content
	updatedInput.UpdatedAt = time.Now()

	db.Model(&comment).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteComment godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} map[string]boolean
// @Router /comments/{id} [delete]
func DeleteComment(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
