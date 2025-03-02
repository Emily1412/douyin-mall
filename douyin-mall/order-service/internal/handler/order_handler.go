package handler

import (
	"context"
	"order-service/internal/service"
	"order-service/proto"
	"fmt"
	"log"
)

// OrderHandler 处理 OrderService 的请求
type OrderHandler struct {
	proto.UnimplementedOrderServiceServer
	OrderService *service.OrderService
}

// NewOrderHandler 创建一个新的 OrderHandler
func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

// PlaceOrder 创建订单
func (h *OrderHandler) PlaceOrder(ctx context.Context, req *proto.PlaceOrderReq) (*proto.PlaceOrderResp, error) {
	// 调用 OrderService 中的 PlaceOrder 方法
	resp, err := h.OrderService.PlaceOrder(ctx, req)
	if err != nil {
		log.Printf("failed to place order: %v", err)
		return nil, fmt.Errorf("could not place order: %v", err)
	}

	return resp, nil
}

// ListOrder 获取订单列表
func (h *OrderHandler) ListOrder(ctx context.Context, req *proto.ListOrderReq) (*proto.ListOrderResp, error) {
	// 调用 OrderService 中的 ListOrder 方法
	resp, err := h.OrderService.ListOrder(ctx, req)
	if err != nil {
		log.Printf("failed to list orders: %v", err)
		return nil, fmt.Errorf("could not list orders: %v", err)
	}

	return resp, nil
}

// MarkOrderPaid 标记订单为已支付
func (h *OrderHandler) MarkOrderPaid(ctx context.Context, req *proto.MarkOrderPaidReq) (*proto.MarkOrderPaidResp, error) {
	// 调用 OrderService 中的 MarkOrderPaid 方法
	resp, err := h.OrderService.MarkOrderPaid(ctx, req)
	if err != nil {
		log.Printf("failed to mark order as paid: %v", err)
		return nil, fmt.Errorf("could not mark order as paid: %v", err)
	}

	return resp, nil
}
