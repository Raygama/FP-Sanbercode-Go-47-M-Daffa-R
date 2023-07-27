package controllers

import (
	"net/http"

	"Final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ratingInput struct {
	Score     string `json:"core"`
	Deskripsi string `json:"deskripsi"`
}

// GetAllRating godoc
// @Summary Get all Rating.
// @Description Get a list of Rating.
// @Tags Rating
// @Produce json
// @Success 200 {object} []models.Rating
// @Router /ratings [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var ratings []models.Rating
	db.Find(&ratings)

	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// CreateRating godoc
// @Summary Create New Rating.
// @Description Creating a new Rating.
// @Tags Rating
// @Param Body body ratingInput true "the body to create a new Rating"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Rating
// @Router /ratings [post]
func CreateRating(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input ratingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating := models.Rating{Score: input.Score, Deskripsi: input.Deskripsi}
	db.Create(&rating)

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetRatingById godoc
// @Summary Get Rating.
// @Description Get a Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "rating id"
// @Success 200 {object} models.Rating
// @Router /ratings/{id} [get]
func GetRatingById(c *gin.Context) { // Get model if exist
	var rating models.Rating

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// GetReviewsByRatingId godoc
// @Summary Get Reviews.
// @Description Get all Reviews by RatingId.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Success 200 {object} []models.Review
// @Router /ratings/{id}/reviews [get]
func GetReviewsByRatingId(c *gin.Context) {
	var reviews []models.Review

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("rating_id = ?", c.Param("id")).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})
}

// UpdateRating godoc
// @Summary Update Rating.
// @Description Update Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Param Body body ratingInput true "the body to update Rating"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Rating
// @Router /ratings/{id} [patch]
func UpdateRating(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&Rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input ratingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Rating
	updatedInput.Score = input.Score
	updatedInput.Deskripsi = input.Deskripsi

	db.Model(&Rating).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": Rating})
}

// DeleteRating godoc
// @Summary Delete one Rating.
// @Description Delete a Rating by id.
// @Tags Rating
// @Produce json
// @Param id path string true "Rating id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /ratings/{id} [delete]
func DeleteRating(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var Rating models.Rating
	if err := db.Where("id = ?", c.Param("id")).First(&Rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	db.Delete(&Rating)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
