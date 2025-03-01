package repository

import (
	"douyin-mall/order-service/internal/model"
	"strconv"

	"gorm.io/gorm"
)

// OrderRepository 封装订单数据操作
type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// PlaceOrder 在数据库中创建订单记录
func (r *OrderRepository) PlaceOrder(order *model.Order) error {
	return r.DB.Create(order).Error
}

// ListOrders 获取指定用户的所有订单
func (r *OrderRepository) ListOrders(userID uint32) ([]model.Order, error) {
	var orders []model.Order
	err := r.DB.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

// MarkOrderPaid 更新订单状态为 "paid"
func (r *OrderRepository) MarkOrderPaid(orderID string, userID uint32) error {
	id, err := strconv.ParseUint(orderID, 10, 64)
	if err != nil {
		return err
	}
	return r.DB.Model(&model.Order{}).
		Where("id = ? AND user_id = ?", uint(id), userID).
		Update("status", "paid").Error
}
