package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"Final/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoryInput struct {
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
}

// GetAllCategory godoc
// @Summary Get all Category.
// @Description Get a list of Category.
// @Tags Category
// @Produce json
// @Param sortById query string false "Sort by Id (asc or desc)"
// @Success 200 {object} []models.Category
// @Router /categories [get]
func GetAllCategory(c *gin.Context) {
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

	var categories []models.Category
	query.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// CreateCategory godoc
// @Summary Create New Category.
// @Description Creating a new Category.
// @Tags Category
// @Param Body body categoryInput true "the body to create a new Category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Category
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{Nama: input.Nama, Deskripsi: input.Deskripsi}
	db.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetCategoryById godoc
// @Summary Get Category.
// @Description Get a Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "category id"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func GetCategoryById(c *gin.Context) {
	var category models.Category

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GetGamesByCategoryId godoc
// @Summary Get Games.
// @Description Get all Games by CategoryId.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Param title query string false "Title filter (case insensitive search)"
// @Param minYear query integer false "Minimum year filter"
// @Param maxYear query integer false "Maximum year filter"
// @Param sortByTitle query string false "Sort by title (asc or desc)"
// @Success 200 {object} []models.Game
// @Router /categories/{id}/games [get]
func GetGamesByCategoryId(c *gin.Context) {
	var games []models.Game
	categoryID := c.Param("id")

	db := c.MustGet("db").(*gorm.DB)
	query := db.Where("category_id = ?", categoryID)

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

	query.Find(&games)

	c.JSON(http.StatusOK, gin.H{"data": games})
}

// UpdateCategory godoc
// @Summary Update Category.
// @Description Update Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Param Body body categoryInput true "the body to update age rating category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} models.Category
// @Router /categories/{id} [patch]
func UpdateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input categoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Category
	updatedInput.Nama = input.Nama
	updatedInput.Deskripsi = input.Deskripsi

	db.Model(&category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DeleteCategory godoc
// @Summary Delete one Category.
// @Description Delete a Category by id.
// @Tags Category
// @Produce json
// @Param id path string true "Category id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var category models.Category
	if err := db.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

	db.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
