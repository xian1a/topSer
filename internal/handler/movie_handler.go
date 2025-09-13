package handler

import (
	"net/http"
	"strconv"
	"topService/internal/model"
	"topService/internal/service"

	"github.com/gin-gonic/gin"
)

type MovieHandler struct {
	movieService *service.MovieService
}

func NewMovieHandler(movieService *service.MovieService) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

// CreateMovie 创建电影
func (h *MovieHandler) CreateMovie(c *gin.Context) {
	var req model.MovieCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}
	
	movie, err := h.movieService.CreateMovie(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "创建电影失败",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"message": "电影创建成功",
		"data":    movie.ToResponse(),
	})
}

// GetMovie 获取单个电影
func (h *MovieHandler) GetMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的电影ID",
		})
		return
	}
	
	movie, err := h.movieService.GetMovieByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": movie.ToResponse(),
	})
}

// GetMovies 获取电影列表
func (h *MovieHandler) GetMovies(c *gin.Context) {
	// 获取分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	keyword := c.Query("search")
	genre := c.Query("genre")
	
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	
	movies, total, err := h.movieService.GetMovies(page, pageSize, keyword, genre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取电影列表失败",
			"details": err.Error(),
		})
		return
	}
	
	// 转换为响应格式
	movieResponses := make([]*model.MovieResponse, len(movies))
	for i, movie := range movies {
		movieResponses[i] = movie.ToResponse()
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"list":      movieResponses,
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

// UpdateMovie 更新电影
func (h *MovieHandler) UpdateMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的电影ID",
		})
		return
	}
	
	var req model.MovieUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}
	
	movie, err := h.movieService.UpdateMovie(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "更新电影失败",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "电影更新成功",
		"data":    movie.ToResponse(),
	})
}

// DeleteMovie 删除电影
func (h *MovieHandler) DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "无效的电影ID",
		})
		return
	}
	
	if err := h.movieService.DeleteMovie(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "电影删除成功",
	})
}

// GetMoviesByGenre 根据类型获取电影
func (h *MovieHandler) GetMoviesByGenre(c *gin.Context) {
	genre := c.Query("genre")
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	
	movies, err := h.movieService.GetMoviesByGenre(genre, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取电影列表失败",
			"details": err.Error(),
		})
		return
	}
	
	// 转换为响应格式
	movieResponses := make([]*model.MovieResponse, len(movies))
	for i, movie := range movies {
		movieResponses[i] = movie.ToResponse()
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": movieResponses,
	})
}

// GetTopRatedMovies 获取高评分电影
func (h *MovieHandler) GetTopRatedMovies(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, _ := strconv.Atoi(limitStr)
	
	movies, err := h.movieService.GetTopRatedMovies(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取高评分电影失败",
			"details": err.Error(),
		})
		return
	}
	
	// 转换为响应格式
	movieResponses := make([]*model.MovieResponse, len(movies))
	for i, movie := range movies {
		movieResponses[i] = movie.ToResponse()
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": movieResponses,
	})
}

// GetMovieStats 获取电影统计信息
func (h *MovieHandler) GetMovieStats(c *gin.Context) {
	stats, err := h.movieService.GetMovieStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "获取统计信息失败",
			"details": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": stats,
	})
}