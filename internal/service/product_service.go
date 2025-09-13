package service

import (
	"errors"
	"topService/internal/model"

	"gorm.io/gorm"
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

// CreateProduct 创建产品
func (s *ProductService) CreateProduct(req *model.ProductCreateRequest) (*model.Product, error) {
	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		Status:      1,
	}
	
	if err := s.db.Create(product).Error; err != nil {
		return nil, err
	}
	
	return product, nil
}

// GetProductByID 根据ID获取产品
func (s *ProductService) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := s.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("产品不存在")
		}
		return nil, err
	}
	
	return &product, nil
}

// GetProducts 获取产品列表
func (s *ProductService) GetProducts(page, pageSize int, keyword, category string) ([]*model.Product, int64, error) {
	var products []*model.Product
	var total int64
	
	query := s.db.Model(&model.Product{})
	
	// 搜索条件
	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	
	if category != "" {
		query = query.Where("category = ?", category)
	}
	
	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&products).Error; err != nil {
		return nil, 0, err
	}
	
	return products, total, nil
}

// UpdateProduct 更新产品
func (s *ProductService) UpdateProduct(id uint, req *model.ProductUpdateRequest) (*model.Product, error) {
	var product model.Product
	if err := s.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("产品不存在")
		}
		return nil, err
	}
	
	// 更新字段
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Stock != nil {
		product.Stock = *req.Stock
	}
	if req.Category != "" {
		product.Category = req.Category
	}
	if req.Status != nil {
		product.Status = *req.Status
	}
	
	if err := s.db.Save(&product).Error; err != nil {
		return nil, err
	}
	
	return &product, nil
}

// DeleteProduct 删除产品
func (s *ProductService) DeleteProduct(id uint) error {
	result := s.db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return errors.New("产品不存在")
	}
	
	return nil
}