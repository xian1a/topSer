package service

import (
	"errors"
	"topService/internal/model"

	"gorm.io/gorm"
)

type MovieService struct {
	db *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{db: db}
}

// CreateMovie 创建电影
func (s *MovieService) CreateMovie(req *model.MovieCreateRequest) (*model.Movie, error) {
	movie := &model.Movie{
		Title:       req.Title,
		Cover:       req.Cover,
		Genre:       req.Genre,
		Director:    req.Director,
		M3u8:        req.M3u8,
		Actors:      req.Actors,
		ReleaseDate: req.ReleaseDate,
		Duration:    req.Duration,
		Language:    req.Language,
		Country:     req.Country,
		Rating:      req.Rating,
		Description: req.Description,
	}
	
	if err := s.db.Create(movie).Error; err != nil {
		return nil, err
	}
	
	return movie, nil
}

// GetMovieByID 根据ID获取电影
func (s *MovieService) GetMovieByID(id uint) (*model.Movie, error) {
	var movie model.Movie
	if err := s.db.First(&movie, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("电影不存在")
		}
		return nil, err
	}
	
	return &movie, nil
}

// GetMovies 获取电影列表
func (s *MovieService) GetMovies(page, pageSize int, keyword, genre string) ([]*model.Movie, int64, error) {
	var movies []*model.Movie
	var total int64
	
	query := s.db.Model(&model.Movie{})
	
	// 搜索条件
	if keyword != "" {
		query = query.Where("title LIKE ? OR director LIKE ? OR actors LIKE ? OR description LIKE ?", 
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}
	
	if genre != "" {
		query = query.Where("genre = ?", genre)
	}
	
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&movies).Error; err != nil {
		return nil, 0, err
	}
	
	return movies, total, nil
}

// UpdateMovie 更新电影
func (s *MovieService) UpdateMovie(id uint, req *model.MovieUpdateRequest) (*model.Movie, error) {
	var movie model.Movie
	if err := s.db.First(&movie, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("电影不存在")
		}
		return nil, err
	}
	
	// 更新字段
	if req.Title != "" {
		movie.Title = req.Title
	}
	if req.Cover != "" {
		movie.Cover = req.Cover
	}
	if req.Genre != "" {
		movie.Genre = req.Genre
	}
	if req.Director != "" {
		movie.Director = req.Director
	}
	if req.M3u8 != "" {
		movie.M3u8 = req.M3u8
	}
	if req.Actors != "" {
		movie.Actors = req.Actors
	}
	if req.ReleaseDate != nil {
		movie.ReleaseDate = req.ReleaseDate
	}
	if req.Duration != nil {
		movie.Duration = *req.Duration
	}
	if req.Language != "" {
		movie.Language = req.Language
	}
	if req.Country != "" {
		movie.Country = req.Country
	}
	if req.Rating != nil {
		movie.Rating = *req.Rating
	}
	if req.Description != "" {
		movie.Description = req.Description
	}
	
	if err := s.db.Save(&movie).Error; err != nil {
		return nil, err
	}
	
	return &movie, nil
}

// DeleteMovie 删除电影
func (s *MovieService) DeleteMovie(id uint) error {
	result := s.db.Delete(&model.Movie{}, id)
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return errors.New("电影不存在")
	}
	
	return nil
}

// GetMoviesByGenre 根据类型获取电影
func (s *MovieService) GetMoviesByGenre(genre string, limit int) ([]*model.Movie, error) {
	var movies []*model.Movie
	query := s.db.Model(&model.Movie{})
	
	if genre != "" {
		query = query.Where("genre = ?", genre)
	}
	
	if limit > 0 {
		query = query.Limit(limit)
	}
	
	if err := query.Order("rating DESC, created_at DESC").Find(&movies).Error; err != nil {
		return nil, err
	}
	
	return movies, nil
}

// GetTopRatedMovies 获取高评分电影
func (s *MovieService) GetTopRatedMovies(limit int) ([]*model.Movie, error) {
	var movies []*model.Movie
	query := s.db.Model(&model.Movie{}).Where("rating >= ?", 8.0)
	
	if limit > 0 {
		query = query.Limit(limit)
	}
	
	if err := query.Order("rating DESC, created_at DESC").Find(&movies).Error; err != nil {
		return nil, err
	}
	
	return movies, nil
}

// GetMovieStats 获取电影统计信息
func (s *MovieService) GetMovieStats() (map[string]interface{}, error) {
	var total int64
	var avgRating float64
	
	// 总电影数
	if err := s.db.Model(&model.Movie{}).Count(&total).Error; err != nil {
		return nil, err
	}
	
	// 平均评分
	if err := s.db.Model(&model.Movie{}).Select("AVG(rating)").Scan(&avgRating).Error; err != nil {
		return nil, err
	}
	
	// 各类型电影数量
	var genreStats []struct {
		Genre string `json:"genre"`
		Count int64  `json:"count"`
	}
	if err := s.db.Model(&model.Movie{}).Select("genre, COUNT(*) as count").
		Where("genre != ''").Group("genre").Scan(&genreStats).Error; err != nil {
		return nil, err
	}
	
	return map[string]interface{}{
		"total":       total,
		"avg_rating":  avgRating,
		"genre_stats": genreStats,
	}, nil
}