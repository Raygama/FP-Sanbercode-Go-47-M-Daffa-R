package controllers

import (
	"net/http"
	"strings"

	"Final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ratingInput struct {
	Score     string `json:"score"`
	Deskripsi string `json:"deskripsi"`
}

// GetAllRating godoc
// @Summary Get all Rating.
// @Description Get a list of Rating.
// @Tags Rating
// @Produce json
// @Param sortById query string false "Sort by Id (asc or desc)"
// @Success 200 {object} []models.Rating
// @Router /ratings [get]
func GetAllRating(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	query := db

	sortById := c.DefaultQuery("sortById", "")
	if sortById != "" {
		if strings.ToLower(sortById) == "asc" {
			query = query.Order("id asc")
		} else if strings.ToLower(sortById) == "desc" {
			query = query.Order("id desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortById parameter"})
			return
		}
	}

	var ratings []models.Rating
	query.Find(&ratings)

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
func GetRatingById(c *gin.Context) {
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
// @Param title query string false "Title filter (case insensitive search)"
// @Param sortByTitle query string false "Sort by title (asc or desc)"
// @Param sortByRatingID query string false "Sort by RatingID (asc or desc)"
// @Param sortByCreatedAt query string false "Sort by created_at (asc or desc)"
// @Success 200 {object} []models.Review
// @Router /ratings/{id}/reviews [get]
func GetReviewsByRatingId(c *gin.Context) {
	ratingID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	title := c.Query("title")
	sortByTitle := c.Query("sortByTitle")
	sortByRatingID := c.Query("sortByRatingID")
	sortByCreatedAt := c.Query("sortByCreatedAt")

	var reviews []models.Review
	query := db.Where("rating_id = ?", ratingID)

	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}

	if sortByTitle != "" {
		if strings.ToLower(sortByTitle) == "asc" {
			query = query.Order("title asc")
		} else if strings.ToLower(sortByTitle) == "desc" {
			query = query.Order("title desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByTitle parameter"})
			return
		}
	}

	if sortByRatingID != "" {
		if strings.ToLower(sortByRatingID) == "asc" {
			query = query.Order("rating_id asc")
		} else if strings.ToLower(sortByRatingID) == "desc" {
			query = query.Order("rating_id desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByRatingID parameter"})
			return
		}
	}

	if sortByCreatedAt != "" {
		if strings.ToLower(sortByCreatedAt) == "asc" {
			query = query.Order("created_at asc")
		} else if strings.ToLower(sortByCreatedAt) == "desc" {
			query = query.Order("created_at desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByCreatedAt parameter"})
			return
		}
	}

	if err := query.Find(&reviews).Error; err != nil {
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
