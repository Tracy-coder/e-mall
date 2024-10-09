package domain

import "context"

type Order interface {
	CreateOrder(ctx context.Context, req CreateOrderReq) error
	ShowOrder(ctx context.Context, orderNum uint64) (*OrderInfo, error)
	ListOrder(ctx context.Context, req ListOrderReq) ([]*OrderInfo, int, error)
	ShowCount(ctx context.Context, userID uint64) (*ShowCountResp, error)
}

type CreateOrderReq struct {
	ProductID uint64
	Num       int32
	AddressID uint64
	UserID    uint64
}

type OrderInfo struct {
	ID           uint64
	OrderNum     int32
	CreatedAt    uint64
	UpdatedAt    uint64
	UserID       uint64
	ProductID    uint64
	Num          int32
	AddressName  string
	AddressPhone string
	Address      string
	Type         uint64
	Name         string
	ImgPath      string
	Price        int64
}

type ListOrderReq struct {
	PageID   int32
	PageSize int32
	Type     uint64
	UserID   uint64
}

type ShowCountResp struct {
	FavoriteTotal int32
	NotPayTotal   int32
	PayTotal      int32
}
