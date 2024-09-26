package domain

import "context"

type Category interface {
	CreateCategory(ctx context.Context, req CreateCategoryReq) error
	ListCategory(ctx context.Context) ([]*CategoryInfo, error)
}

type CreateCategoryReq struct {
	ID           uint64
	CategoryName string
}

type CategoryInfo struct {
	ID           uint64
	CategoryName string
	CreatedAt    uint64
}
