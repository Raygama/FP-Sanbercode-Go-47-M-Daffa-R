package controllers

import (
	"net/http"
	"strings"
	"time"

	"Final/models"
	"Final/utils/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentInput struct {
	ReviewID int    `json:"review_id"`
	Content  string `json:"content" gorm:"type:text"`
	Likes    int    `json:"likes"`
}

// GetAllComments godoc
// @Summary Get all Comments.
// @Description Get a list of Comments.
// @Tags Comment
// @Produce json
// @Param sortByLikes query string false "Sort by likes (asc or desc)"
// @Param sortByCreatedAt query string false "Sort by created_at (asc or desc)"
// @Success 200 {object} []models.Comment
// @Router /comments [get]
func GetAllComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	sortByLikes := c.Query("sortByLikes")
	sortByCreatedAt := c.Query("sortByCreatedAt")

	var comments []models.Comment
	query := db

	if sortByLikes != "" {
		if strings.ToLower(sortByLikes) == "asc" {
			query = query.Order("likes asc")
		} else if strings.ToLower(sortByLikes) == "desc" {
			query = query.Order("likes desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByLikes parameter"})
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

	if err := query.Find(&comments).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
		return
	}

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

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.ReviewID).First(&review).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ReviewID not found!"})
		return
	}

	UserID, err := token.ExtractUserID(c)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}

	comment := models.Comment{ReviewID: input.ReviewID, UserID: UserID, Content: input.Content, Likes: input.Likes}
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
func GetCommentById(c *gin.Context) {
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
// @Param sortByLikes query string false "Sort by likes (asc or desc)"
// @Param sortByCreatedAt query string false "Sort by created_at (asc or desc)"
// @Success 200 {object} []models.Comment
// @Router /users/{id}/comments [get]
func GetCommentsByUserId(c *gin.Context) {
	userID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	sortByLikes := c.Query("sortByLikes")
	sortByCreatedAt := c.Query("sortByCreatedAt")

	var comments []models.Comment
	query := db.Where("user_id = ?", userID)

	if sortByLikes != "" {
		if strings.ToLower(sortByLikes) == "asc" {
			query = query.Order("likes asc")
		} else if strings.ToLower(sortByLikes) == "desc" {
			query = query.Order("likes desc")
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortByLikes parameter"})
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

	if err := query.Find(&comments).Error; err != nil {
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

	var comment models.Comment
	var review models.Review
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

	var updatedInput models.Comment
	updatedInput.ReviewID = input.ReviewID
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
	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
