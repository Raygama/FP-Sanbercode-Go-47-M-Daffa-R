package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"Final/controllers"
	"Final/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/reviews", controllers.GetAllReview)
	r.GET("/reviews/:id", controllers.GetReviewById)

	reviewsMiddlewareRoute := r.Group("/reviews")
	reviewsMiddlewareRoute.Use(middlewares.UserCheckMiddleware())
	reviewsMiddlewareRoute.POST("/", controllers.CreateReview)
	reviewsMiddlewareRoute.PATCH("/:id", controllers.UpdateReview)
	reviewsMiddlewareRoute.DELETE("/:id", controllers.DeleteReview)

	r.GET("/ratings", controllers.GetAllRating)
	r.GET("/ratings/:id", controllers.GetRatingById)
	r.GET("/ratings/:id/review", controllers.GetReviewsByRatingId)

	ratingsMiddlewareRoute := r.Group("/ratings")
	ratingsMiddlewareRoute.Use(middlewares.AdminCheckMiddleware())
	ratingsMiddlewareRoute.POST("/", controllers.CreateRating)
	ratingsMiddlewareRoute.PATCH("/:id", controllers.UpdateRating)
	ratingsMiddlewareRoute.DELETE("/:id", controllers.DeleteRating)

	r.GET("/categories", controllers.GetAllCategory)
	r.GET("/categories/:id", controllers.GetCategoryById)
	r.GET("/categories/:id/games", controllers.GetGamesByCategoryId)

	categoriesMiddlewareRoute := r.Group("/categories")
	categoriesMiddlewareRoute.Use(middlewares.AdminCheckMiddleware())
	categoriesMiddlewareRoute.POST("/", controllers.CreateCategory)
	categoriesMiddlewareRoute.PATCH("/:id", controllers.UpdateCategory)
	categoriesMiddlewareRoute.DELETE("/:id", controllers.DeleteCategory)

	r.GET("/games", controllers.GetAllGames)
	r.GET("/games/:id", controllers.GetGamesById)
	r.GET("/games/:id/reviews", controllers.GetReviewsByGameId)

	gamesMiddlewareRoute := r.Group("/games")
	gamesMiddlewareRoute.Use(middlewares.AdminCheckMiddleware())
	gamesMiddlewareRoute.POST("/", controllers.CreateGames)
	gamesMiddlewareRoute.PATCH("/:id", controllers.UpdateGames)
	gamesMiddlewareRoute.DELETE("/:id", controllers.DeleteGames)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
