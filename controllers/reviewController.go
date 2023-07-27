package controllers

import (
	"net/http"
	"time"

	"Final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type reviewInput struct {
	GameID   int    `json:"game_id"`
	RatingID int    `json:"rating_id"`
	UserID   int    `json:"user_id"`
	Title    string `json:"title" gorm:"type:varchar(255)"`
	Content  string `json:"content" gorm:"type:text"`
}

// GetAllReviews godoc
// @Summary Get all reviews.
// @Description Get a list of reviews.
// @Tags Review
// @Produce json
// @Success 200 {object} []models.Review
// @Router /reviews [get]
func GetAllReview(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var reviews []models.Review
	db.Find(&reviews)

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// CreateReview godoc
// @Summary Create New Review.
// @Description Creating a new Review.
// @Tags Review
// @Param Body body reviewInput true "the body to create a new review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Review
// @Router /reviews [post]
func CreateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input reviewInput
	var game models.Game
	var rating models.Rating
	var user models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found!"})
		return
	}

	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	if err := db.Where("id = ?", input.RatingID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RatingID not found!"})
		return
	}

	review := models.Review{GameID: input.GameID, RatingID: input.RatingID, Title: input.Title, Content: input.Content}
	db.Create(&review)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// GetReviewById godoc
// @Summary Get Review.
// @Description Get a Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "review id"
// @Success 200 {object} models.Review
// @Router /reviews/{id} [get]
func GetReviewById(c *gin.Context) { // Get model if exist
	var review models.Review

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// UpdateReview godoc
// @Summary Update Review.
// @Description Update Review by id.
// @Tags Review
// @Produce json
// @Param id path string true "review id"
// @Param Body body reviewInput true "the body to update a review"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Review
// @Router /reviews/{id} [patch]
func UpdateReview(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var review models.Review
	var rating models.Rating
	var game models.Game
	var user models.User
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input reviewInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.RatingID).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "RatingID not found!"})
		return
	}

	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	if err := db.Where("id = ?", input.GameID).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "GameID not found!"})
		return
	}

	var updatedInput models.Review
	updatedInput.GameID = input.GameID
	updatedInput.RatingID = input.RatingID
	updatedInput.Title = input.Title
	updatedInput.Content = input.Content
	updatedInput.UpdatedAt = time.Now()

	db.Model(&review).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": review})
}

// DeleteReview godoc
// @Summary Delete one Review.
// @Description Delete a review by id.
// @Tags Review
// @Produce json
// @Param id path string true "review id"
// @Success 200 {object} map[string]boolean
// @Router /reviews/{id} [delete]
func DeleteReview(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var review models.Review
	if err := db.Where("id = ?", c.Param("id")).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
