package model

import (
	"time"
)

type Movie struct {
	ID          uint      `json:"id" gorm:"primarykey;comment:电影ID"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:create_at;comment:创建时间"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:update_at;comment:更新时间"`
	
	Title       string     `json:"title" gorm:"not null;size:255;comment:电影名称" binding:"required,min=1,max=255"`
	Cover       string     `json:"cover" gorm:"size:255;comment:封面"`
	Genre       string     `json:"genre" gorm:"size:100;comment:电影类型"`
	Director    string     `json:"director" gorm:"size:100;comment:导演"`
	M3u8        string     `json:"m3u8" gorm:"size:500"`
	Actors      string     `json:"actors" gorm:"size:500;comment:主演"`
	ReleaseDate *time.Time `json:"release_date" gorm:"type:date;comment:上映日期"`
	Duration    int        `json:"duration" gorm:"comment:片长（分钟）" binding:"min=0"`
	Language    string     `json:"language" gorm:"size:50;comment:语言"`
	Country     string     `json:"country" gorm:"size:100;comment:国家/地区"`
	Rating      float32    `json:"rating" gorm:"type:decimal(2,1);comment:评分 (0.0 - 10.0)" binding:"min=0,max=10"`
	Description string     `json:"description" gorm:"type:text;comment:剧情简介"`
}

// TableName 指定表名
func (Movie) TableName() string {
	return "movies"
}

// MovieCreateRequest 创建电影请求
type MovieCreateRequest struct {
	Title       string     `json:"title" binding:"required,min=1,max=255"`
	Cover       string     `json:"cover"`
	Genre       string     `json:"genre"`
	Director    string     `json:"director"`
	M3u8        string     `json:"m3u8"`
	Actors      string     `json:"actors"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    int        `json:"duration" binding:"min=0"`
	Language    string     `json:"language"`
	Country     string     `json:"country"`
	Rating      float32    `json:"rating" binding:"min=0,max=10"`
	Description string     `json:"description"`
}

// MovieUpdateRequest 更新电影请求
type MovieUpdateRequest struct {
	Title       string     `json:"title" binding:"omitempty,min=1,max=255"`
	Cover       string     `json:"cover"`
	Genre       string     `json:"genre"`
	Director    string     `json:"director"`
	M3u8        string     `json:"m3u8"`
	Actors      string     `json:"actors"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    *int       `json:"duration" binding:"omitempty,min=0"`
	Language    string     `json:"language"`
	Country     string     `json:"country"`
	Rating      *float32   `json:"rating" binding:"omitempty,min=0,max=10"`
	Description string     `json:"description"`
}

// MovieResponse 电影响应
type MovieResponse struct {
    ID          uint       `json:"id"`
    Title       string     `json:"title"`
    Poster      string     `json:"poster"`
    Genre       string     `json:"genre"`
    Director    string     `json:"director"`
    VideoUrl    string     `json:"videoUrl"`
    Actors      string     `json:"actors"`
    ReleaseDate *time.Time `json:"release_date"`
    Duration    int        `json:"duration"`
    Language    string     `json:"language"`
    Country     string     `json:"country"`
    Rating      float32    `json:"rating"`
    Description string     `json:"description"`
    CreatedAt   time.Time  `json:"created_at"`
    UpdatedAt   time.Time  `json:"updated_at"`
}


// ToResponse 转换为响应格式
func (m *Movie) ToResponse() *MovieResponse {
    return &MovieResponse{
        ID:          m.ID,
        Title:       m.Title,
        Poster:      m.Cover,
        Genre:       m.Genre,
        Director:    m.Director,
        VideoUrl:    m.M3u8,
        Actors:      m.Actors,
        ReleaseDate: m.ReleaseDate,
        Duration:    m.Duration,
        Language:    m.Language,
        Country:     m.Country,
        Rating:      m.Rating,
        Description: m.Description,
        CreatedAt:   m.CreatedAt,
        UpdatedAt:   m.UpdatedAt,
    }
}
