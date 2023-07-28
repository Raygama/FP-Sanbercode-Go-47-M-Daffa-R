package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"Final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type gameInput struct {
	CategoryID    int    `json:"category_id"`
	Nama          string `json:"nama"`
	Deskripsi     string `json:"deskripsi"`
	Developer     string `json:"developer" gorm:"type:varchar(255)"`
	YearPublished string `json:"year_published" gorm:"type:varchar(10)"`
}

type filterInput struct {
	Title   string `form:"title"`
	MinYear int    `form:"minYear"`
	MaxYear int    `form:"maxYear"`
}

// GetAllGames godoc
// @Summary Get all Game.
// @Description Get a list of Game.
// @Tags Game
// @Produce json
// @Param title query string false "Title filter (case insensitive search)"
// @Param minYear query integer false "Minimum year filter"
// @Param maxYear query integer false "Maximum year filter"
// @Param sortByTitle query string false "Sort by title (asc or desc)"
// @Success 200 {object} []models.Game
// @Router /games [get]
func GetAllGames(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	query := db
	title := c.Query("title")
	if title != "" {
		query = query.Where("LOWER(nama) LIKE ?", "%"+strings.ToLower(title)+"%")
	}

	minYearStr := c.Query("minYear")
	if minYearStr != "" {
		minYear, err := strconv.Atoi(minYearStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minYear parameter"})
			return
		}
		query = query.Where("year_published >= ?", minYear)
	}

	maxYearStr := c.Query("maxYear")
	if maxYearStr != "" {
		maxYear, err := strconv.Atoi(maxYearStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maxYear parameter"})
			return
		}
		query = query.Where("year_published <= ?", maxYear)
	}

	sortByTitle := c.DefaultQuery("sortByTitle", "")
	if sortByTitle != "" {
		if strings.ToLower(sortByTitle) == "asc" {
			query = query.Order("nama asc")
		} else if strings.ToLower(sortByTitle) == "desc" {
			query = query.Order("nama desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByTitle parameter"})
			return
		}
	}

	var games []models.Game
	query.Find(&games)

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// CreateGame godoc
// @Summary Create New Game.
// @Description Creating a new Game.
// @Tags Game
// @Param Body body gameInput true "the body to create a new Category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Game
// @Router /games [post]
func CreateGames(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input gameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := models.Game{CategoryID: input.CategoryID, Nama: input.Nama, Deskripsi: input.Deskripsi, Developer: input.Developer, YearPublished: input.YearPublished}
	db.Create(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// GetGameById godoc
// @Summary Get Game.
// @Description Get a Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "game id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Game
// @Router /games/{id} [get]
func GetGamesById(c *gin.Context) {
	var game models.Game

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// GetReviewsByGameId godoc
// @Summary Get Reviews.
// @Description Get all Reviews by GameId.
// @Tags Game
// @Produce json
// @Param id path string true "Game id"
// @Param title query string false "Title filter (case insensitive search)"
// @Param sortByTitle query string false "Sort by title (asc or desc)"
// @Param sortByRatingID query string false "Sort by RatingID (asc or desc)"
// @Param sortByCreatedAt query string false "Sort by created_at (asc or desc)"
// @Success 200 {object} []models.Review
// @Router /games/{id}/reviews [get]
func GetReviewsByGameId(c *gin.Context) {
	gameID := c.Param("id")

	db := c.MustGet("db").(*gorm.DB)

	title := c.Query("title")
	sortByTitle := c.Query("sortByTitle")
	sortByRatingID := c.Query("sortByRatingID")
	sortByCreatedAt := c.Query("sortByCreatedAt")

	var reviews []models.Review
	query := db.Where("game_id = ?", gameID)

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

// UpdateGame godoc
// @Summary Update Game.
// @Description Update Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "Game id"
// @Param Body body gameInput true "the body to update age rating game"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Game
// @Router /games/{id} [patch]
func UpdateGames(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input gameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Game
	updatedInput.CategoryID = input.CategoryID
	updatedInput.Nama = input.Nama
	updatedInput.Deskripsi = input.Deskripsi
	updatedInput.Developer = input.Developer
	updatedInput.YearPublished = input.YearPublished

	db.Model(&game).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// DeleteGame godoc
// @Summary Delete one Game.
// @Description Delete a Game by id.
// @Tags Game
// @Produce json
// @Param id path string true "Game id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /games/{id} [delete]
func DeleteGames(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var game models.Game
	if err := db.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	db.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
