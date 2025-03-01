package repository

import (
	"douyin-mall/cart-service/internal/model"

	"gorm.io/gorm"
)

// CartRepository 封装购物车数据操作
type CartRepository struct {
	DB *gorm.DB
}

// NewCartRepository 创建一个新的 CartRepository 实例
func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{DB: db}
}

// AddItem 将商品添加到购物车
func (r *CartRepository) AddItem(userID uint32, productID uint32, quantity int32) error {
	// 查找用户购物车

	var cart model.Cart
	err := r.DB.First(&cart, "user_id = ?", userID).Error
	if err != nil {
		// 如果错误是记录不存在，则创建一个新的购物车
		if err == gorm.ErrRecordNotFound {
			cart = model.Cart{UserID: userID}
			if err := r.DB.Create(&cart).Error; err != nil {
				return err
			}
		} else {
			// 其他错误直接返回
			return err
		}
	}

	// 查找购物车中的商品
	var cartItem model.CartItem
	err = r.DB.First(&cartItem, "cart_id = ? AND product_id = ?", cart.ID, productID).Error
	if err == nil {
		// 商品已存在，更新数量
		cartItem.Quantity += quantity
		r.DB.Save(&cartItem)
	} else {
		// 商品不存在，新增商品
		newItem := model.CartItem{
			CartID:    cart.ID,
			ProductID: productID,
			Quantity:  quantity,
		}
		r.DB.Create(&newItem)
	}

	return nil
}

// GetCart 获取用户的购物车信息
func (r *CartRepository) GetCart(userID uint32) (*model.Cart, error) {
	var cart model.Cart
	if err := r.DB.Preload("Items").First(&cart, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &cart, nil
}

// EmptyCart 清空购物车
func (r *CartRepository) EmptyCart(userID uint32) error {
	var cart model.Cart
	if err := r.DB.First(&cart, "user_id = ?", userID).Error; err != nil {
		return err
	}
	// 删除购物车中的所有商品
	r.DB.Delete(&model.CartItem{}, "cart_id = ?", cart.ID)
	return nil
}
