package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	
	Name        string  `json:"name" gorm:"not null;size:100" binding:"required,min=1,max=100"`
	Description string  `json:"description" gorm:"size:500"`
	Price       float64 `json:"price" gorm:"not null" binding:"required,gt=0"`
	Stock       int     `json:"stock" gorm:"default:0" binding:"min=0"`
	Category    string  `json:"category" gorm:"size:50"`
	Status      int     `json:"status" gorm:"default:1"` // 1:上架 0:下架
}

// TableName 指定表名
func (Product) TableName() string {
	return "products"
}

// ProductCreateRequest 创建产品请求
type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required,min=1,max=100"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"min=0"`
	Category    string  `json:"category"`
}

// ProductUpdateRequest 更新产品请求
type ProductUpdateRequest struct {
	Name        string   `json:"name" binding:"omitempty,min=1,max=100"`
	Description string   `json:"description"`
	Price       *float64 `json:"price" binding:"omitempty,gt=0"`
	Stock       *int     `json:"stock" binding:"omitempty,min=0"`
	Category    string   `json:"category"`
	Status      *int     `json:"status" binding:"omitempty,oneof=0 1"`
}

// ToResponse 转换为响应格式
func (p *Product) ToResponse() *ProductResponse {
	return &ProductResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		Category:    p.Category,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

// ProductResponse 产品响应
type ProductResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Category    string    `json:"category"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}