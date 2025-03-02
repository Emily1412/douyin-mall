package service

import (
	"context"
	cartpb "cart-service/proto"
	"order-service/internal/model"
	"order-service/internal/repository"
	pb "order-service/proto"
	"fmt"
	"strconv"

	"gorm.io/gorm"
)

// OrderService 实现 OrderService 的业务逻辑
type OrderService struct {
	pb.UnimplementedOrderServiceServer
	Repo *repository.OrderRepository
}

func NewOrderService(db *gorm.DB) *OrderService {
	repo := repository.NewOrderRepository(db)
	return &OrderService{Repo: repo}
}

// PlaceOrder 创建订单
func (s *OrderService) PlaceOrder(ctx context.Context, req *pb.PlaceOrderReq) (*pb.PlaceOrderResp, error) {
	// 将地址转换为字符串格式，这里简单拼接
	address := fmt.Sprintf("%s, %s, %s, %s, %d",
		req.Address.StreetAddress,
		req.Address.City,
		req.Address.State,
		req.Address.Country,
		req.Address.ZipCode,
	)
	order := &model.Order{
		UserID:       req.UserId,
		UserCurrency: req.UserCurrency,
		Address:      address,
		Email:        req.Email,
		Status:       "pending",
	}
	// 构建订单项
	for _, oi := range req.OrderItems {
		orderItem := model.OrderItem{
			ProductID: oi.Item.ProductId,
			Quantity:  oi.Item.Quantity,
			Cost:      float64(oi.Cost),
		}
		order.OrderItems = append(order.OrderItems, orderItem)
	}
	if err := s.Repo.PlaceOrder(order); err != nil {
		return nil, err
	}
	return &pb.PlaceOrderResp{
		Order: &pb.OrderResult{
			OrderId: strconv.FormatUint(uint64(order.ID), 10),
		},
	}, nil
}

// ListOrder 获取订单列表
func (s *OrderService) ListOrder(ctx context.Context, req *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	orders, err := s.Repo.ListOrders(req.UserId)
	if err != nil {
		return nil, err
	}
	var protoOrders []*pb.Order
	for _, o := range orders {
		protoOrder := &pb.Order{
			OrderId:      strconv.FormatUint(uint64(o.ID), 10),
			UserId:       o.UserID,
			UserCurrency: o.UserCurrency,
			Address: &pb.Address{
				StreetAddress: "", // You need to parse the address string and fill these fields accordingly
				City:          "",
				State:         "",
				Country:       "",
				ZipCode:       0,
			}, // 简单返回存储的地址字符串
			Email:     o.Email,
			CreatedAt: int32(o.CreatedAt.Unix()),
		}
		// 转换订单项
		for _, item := range o.OrderItems {
			protoItem := &pb.OrderItem{
				Item: &cartpb.CartItem{
					ProductId: item.ProductID,
					Quantity:  item.Quantity,
				},
				Cost: float32(item.Cost),
			}
			protoOrder.OrderItems = append(protoOrder.OrderItems, protoItem)
		}
		protoOrders = append(protoOrders, protoOrder)
	}
	return &pb.ListOrderResp{
		Orders: protoOrders,
	}, nil
}

// MarkOrderPaid 标记订单为已支付
func (s *OrderService) MarkOrderPaid(ctx context.Context, req *pb.MarkOrderPaidReq) (*pb.MarkOrderPaidResp, error) {
	if err := s.Repo.MarkOrderPaid(req.OrderId, req.UserId); err != nil {
		return nil, err
	}
	return &pb.MarkOrderPaidResp{}, nil
}
