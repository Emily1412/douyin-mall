package handler

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"product-service/api"
	"product-service/internal/service"
)

// ProductHandler 定义商品处理结构体
type ProductHandler struct {
	api.UnimplementedProductCatalogServiceServer
	service *service.ProductService
}

// NewProductHandler 创建一个新的商品处理实例
func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// ListProducts 处理获取商品列表请求
func (h *ProductHandler) ListProducts(ctx context.Context, req *api.ListProductsReq) (*api.ListProductsResp, error) {
	return h.service.ListProducts(ctx, req)
}

// GetProduct 处理通过 ID 获取单个商品请求
func (h *ProductHandler) GetProduct(ctx context.Context, req *api.GetProductReq) (*api.GetProductResp, error) {
	return h.service.GetProduct(ctx, req)
}

// SearchProducts 处理根据查询关键字搜索商品请求
func (h *ProductHandler) SearchProducts(ctx context.Context, req *api.SearchProductsReq) (*api.SearchProductsResp, error) {
	return h.service.SearchProducts(ctx, req)
}

// CreateProducts 处理创建商品请求
func (h *ProductHandler) CreateProducts(ctx context.Context, req *api.Product) (*emptypb.Empty, error) {
	return h.service.CreateProducts(ctx, req)
}

// DeleteProducts 处理删除商品请求
func (h *ProductHandler) DeleteProducts(ctx context.Context, req *api.DeleteProductsInfo) (*emptypb.Empty, error) {
	return h.service.DeleteProducts(ctx, req)
}

// UpdateProducts 处理更新商品信息请求
func (h *ProductHandler) UpdateProducts(ctx context.Context, req *api.Product) (*emptypb.Empty, error) {
	return h.service.UpdateProducts(ctx, req)
}
