package handler

import (
	"context"

	"cart-service/internal/service"
	"cart-service/proto"
)

// CartHandler 处理与购物车相关的请求

type CartHandler struct {
	service *service.CartService

	proto.UnimplementedCartServiceServer
}

// NewCartHandler 创建一个新的 CartHandler 实例
func NewCartHandler(cartService *service.CartService) *CartHandler {
	return &CartHandler{service: cartService}
}

// AddItem 处理添加商品到购物车的请求
func (h *CartHandler) AddItem(ctx context.Context, req *proto.AddItemReq) (*proto.AddItemResp, error) {
	// 调用 service 层的 AddItem 方法
	return h.service.AddItem(ctx, req)
}

// GetCart 处理获取购物车信息的请求
func (h *CartHandler) GetCart(ctx context.Context, req *proto.GetCartReq) (*proto.GetCartResp, error) {
	// 调用 service 层的 GetCart 方法
	return h.service.GetCart(ctx, req)
}

// EmptyCart 处理清空购物车的请求
func (h *CartHandler) EmptyCart(ctx context.Context, req *proto.EmptyCartReq) (*proto.EmptyCartResp, error) {
	// 调用 service 层的 EmptyCart 方法
	return h.service.EmptyCart(ctx, req)
}
