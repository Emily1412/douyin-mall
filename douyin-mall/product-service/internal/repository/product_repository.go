package repository

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"product-service/internal/model"
)

// ProductRepository 定义商品仓库结构体
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository 创建一个新的商品仓库实例
func NewProductRepository(dsn string) (*ProductRepository, error) {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// 自动迁移模型
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return &ProductRepository{
		db: db,
	}, nil
}

// ListProducts 获取商品列表
func (r *ProductRepository) ListProducts(ctx context.Context, page int32, pageSize int64, categoryName string) ([]model.Product, error) {
	var products []model.Product
	query := r.db.WithContext(ctx)

	// 如果指定了分类名称，添加过滤条件
	if categoryName != "" {
		query = query.Where("FIND_IN_SET(?, categories) > 0", categoryName)
	}

	// 分页查询
	err := query.Offset(int((page - 1) * int32(pageSize))).Limit(int(pageSize)).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to list products: %w", err)
	}

	return products, nil
}

// GetProduct 通过 ID 获取单个商品
func (r *ProductRepository) GetProduct(ctx context.Context, id uint32) (model.Product, error) {
	var product model.Product
	err := r.db.WithContext(ctx).First(&product, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Product{}, nil
		}
		return model.Product{}, fmt.Errorf("failed to get product: %w", err)
	}

	return product, nil
}

// SearchProducts 根据查询关键字搜索商品
func (r *ProductRepository) SearchProducts(ctx context.Context, query string) ([]model.Product, error) {
	var products []model.Product
	// 构建包含 categories 字段的查询条件
	queryCondition := "name LIKE ? OR description LIKE ? OR categories LIKE ?"
	err := r.db.WithContext(ctx).Where(queryCondition, "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to search products: %w", err)
	}

	return products, nil
}

// CreateProducts 创建商品
func (r *ProductRepository) CreateProducts(ctx context.Context, product model.Product) error {
	err := r.db.WithContext(ctx).Create(&product).Error
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

// DeleteProducts 删除商品
func (r *ProductRepository) DeleteProducts(ctx context.Context, id int32) error {
	err := r.db.WithContext(ctx).Delete(&model.Product{}, id).Error
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}

// UpdateProducts 更新商品信息
func (r *ProductRepository) UpdateProducts(ctx context.Context, product model.Product) error {
	err := r.db.WithContext(ctx).Save(&product).Error
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}
