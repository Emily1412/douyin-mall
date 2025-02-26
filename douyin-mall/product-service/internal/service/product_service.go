package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"product-service/api"
	"product-service/internal/model"
	"product-service/internal/repository"
)

// ProductService 定义商品服务结构体
type ProductService struct {
	repo *repository.ProductRepository
}

// NewProductService 创建一个新的商品服务实例
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// ListProducts 获取商品列表
func (s *ProductService) ListProducts(ctx context.Context, req *api.ListProductsReq) (*api.ListProductsResp, error) {
	products, err := s.repo.ListProducts(ctx, req.GetPage(), req.GetPageSize(), req.GetCategoryName())
	if err != nil {
		return nil, err
	}

	respProducts := make([]*api.Product, len(products))
	for i, p := range products {
		respProducts[i] = &api.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  p.Categories,
		}
	}

	return &api.ListProductsResp{
		Products: respProducts,
	}, nil
}

// GetProduct 通过 ID 获取单个商品
func (s *ProductService) GetProduct(ctx context.Context, req *api.GetProductReq) (*api.GetProductResp, error) {
	product, err := s.repo.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	respProduct := &api.Product{
		Id:          uint32(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Picture:     product.Picture,
		Price:       product.Price,
		Categories:  product.Categories,
	}

	return &api.GetProductResp{
		Product: respProduct,
	}, nil
}

// SearchProducts 根据查询关键字搜索商品
func (s *ProductService) SearchProducts(ctx context.Context, req *api.SearchProductsReq) (*api.SearchProductsResp, error) {
	products, err := s.repo.SearchProducts(ctx, req.GetQuery())
	if err != nil {
		return nil, err
	}

	respProducts := make([]*api.Product, len(products))
	for i, p := range products {
		respProducts[i] = &api.Product{
			Id:          uint32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Picture:     p.Picture,
			Price:       p.Price,
			Categories:  p.Categories,
		}
	}

	return &api.SearchProductsResp{
		Results: respProducts,
	}, nil
}

// CreateProducts 创建商品
func (s *ProductService) CreateProducts(ctx context.Context, req *api.Product) (*emptypb.Empty, error) {
	product := model.Product{
		ID:          int32(req.GetId()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Picture:     req.GetPicture(),
		Price:       req.GetPrice(),
		Categories:  req.GetCategories(),
	}

	err := s.repo.CreateProducts(ctx, product)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// DeleteProducts 删除商品
func (s *ProductService) DeleteProducts(ctx context.Context, req *api.DeleteProductsInfo) (*emptypb.Empty, error) {
	err := s.repo.DeleteProducts(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// UpdateProducts 更新商品信息
func (s *ProductService) UpdateProducts(ctx context.Context, req *api.Product) (*emptypb.Empty, error) {
	product := model.Product{
		ID:          int32(req.GetId()),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Picture:     req.GetPicture(),
		Price:       req.GetPrice(),
		Categories:  req.GetCategories(),
	}

	err := s.repo.UpdateProducts(ctx, product)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
