package domain

import "context"

type Cart interface {
	CreateCart(ctx context.Context, req CreateCartReq) (*CartInfo, error)
	ShowCart(ctx context.Context, userID uint64) ([]*CartInfo, error)
	UpdateCart(ctx context.Context, req UpdateCartReq) error
	DeleteCart(ctx context.Context, id uint64, userID uint64) error
}

type CreateCartReq struct {
	UserID    uint64
	ProductID uint64
}

type CartInfo struct {
	ID        uint64
	UserID    uint64
	ProductID uint64
	Num       int32
	MaxNum    int32
	Check     bool
	CreatedAt uint64
}

type UpdateCartReq struct {
	Num       int32
	ProductID uint64
	UserID    uint64
}
