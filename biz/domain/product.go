package domain

import "context"

type Product interface {
	CreateProduct(ctx context.Context, req CreateProductReq) (*ProductInfo, error)
	ShowProduct(ctx context.Context, id uint64) (*ProductInfo, error)
	ListProduct(ctx context.Context, req ListProductReq) ([]*ProductInfo, error)
	UpdateProduct(ctx context.Context, req UpdateProductReq) error
	DeleteProduct(ctx context.Context, id uint64) error
	CreateProductImg(ctx context.Context, req CreateProductImgReq) (*ProductImg, error)
	ShowProductImg(ctx context.Context, id uint64) ([]*ProductImg, error)
	ShowProductRankings(ctx context.Context) ([]*ProductInfo, error)
}

type CreateProductReq struct {
	Name          string
	CategoryID    uint64
	Title         string
	Info          string
	ImgPath       string
	Price         int64
	DiscountPrice int64
}
type UpdateProductReq struct {
	ID            uint64
	Name          string
	CategoryID    uint64
	Title         string
	Info          string
	ImgPath       string
	Price         int64
	DiscountPrice int64
}
type ListProductReq struct {
	Limit      int32
	Start      int32
	CategoryID uint64
}
type ProductInfo struct {
	Name          string
	CategoryID    uint64
	Title         string
	Info          string
	ImgPath       string
	DiscountPrice int64
	View          uint64
	CreatedAt     uint64
	Price         int64
}

type CreateProductImgReq struct {
	ProductID uint64
	ImgPath   string
}

type ProductImg struct {
	ID        uint64
	ProductID uint64
	ImgPath   string
	CreatedAt uint64
}
