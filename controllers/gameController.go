package controllers

import (
	"net/http"

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

// GetAllGames godoc
// @Summary Get all Game.
// @Description Get a list of Game.
// @Tags Game
// @Produce json
// @Success 200 {object} []models.Game
// @Router /games [get]
func GetAllGames(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var games []models.Game
	db.Find(&games)

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
func GetGamesById(c *gin.Context) { // Get model if exist
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
// @Success 200 {object} []models.Review
// @Router /games/{id}/reviews [get]
func GetReviewsByGameId(c *gin.Context) {
	var reviews []models.Review

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("game_id = ?", c.Param("id")).Find(&reviews).Error; err != nil {
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

	// Validate input
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
