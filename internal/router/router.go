package router

import (
	"net/http"
	"topService/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler *handler.UserHandler, productHandler *handler.ProductHandler, movieHandler *handler.MovieHandler) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"message": "topService is running",
		})
	})
	
	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 用户相关路由
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
		
		// 产品相关路由
		products := v1.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
		
		// 电影相关路由
		movies := v1.Group("/movies")
		{
			movies.POST("", movieHandler.CreateMovie)
			movies.GET("", movieHandler.GetMovies)
			movies.GET("/stats", movieHandler.GetMovieStats)
			movies.GET("/top-rated", movieHandler.GetTopRatedMovies)
			movies.GET("/by-genre", movieHandler.GetMoviesByGenre)
			movies.GET("/:id", movieHandler.GetMovie)
			movies.PUT("/:id", movieHandler.UpdateMovie)
			movies.DELETE("/:id", movieHandler.DeleteMovie)
		}
	}
}