package service

import (
	"context"

	"cart-service/internal/repository"
	"cart-service/proto"

	"gorm.io/gorm"
)

// CartService 购物车服务
type CartService struct {
	proto.UnimplementedCartServiceServer
	Repo *repository.CartRepository
}

// NewCartService 创建一个新的 CartService 实例
func NewCartService(db *gorm.DB) *CartService {
	repo := repository.NewCartRepository(db)
	return &CartService{Repo: repo}
}

// AddItem 添加商品到购物车
func (s *CartService) AddItem(ctx context.Context, req *proto.AddItemReq) (*proto.AddItemResp, error) {
	err := s.Repo.AddItem(req.UserId, req.Item.ProductId, req.Item.Quantity)
	if err != nil {
		return nil, err
	}
	return &proto.AddItemResp{}, nil
}

// GetCart 获取购物车
func (s *CartService) GetCart(ctx context.Context, req *proto.GetCartReq) (*proto.GetCartResp, error) {
	cart, err := s.Repo.GetCart(req.UserId)
	if err != nil {
		return nil, err
	}

	var items []*proto.CartItem
	for _, item := range cart.Items {
		items = append(items, &proto.CartItem{
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	return &proto.GetCartResp{
		Cart: &proto.Cart{
			UserId: cart.UserID,
			Items:  items,
		},
	}, nil
}

// EmptyCart 清空购物车
func (s *CartService) EmptyCart(ctx context.Context, req *proto.EmptyCartReq) (*proto.EmptyCartResp, error) {
	err := s.Repo.EmptyCart(req.UserId)
	if err != nil {
		return nil, err
	}
	return &proto.EmptyCartResp{}, nil
}
