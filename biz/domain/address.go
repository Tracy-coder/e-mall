package domain

import "context"

type Address interface {
	CreateAddress(ctx context.Context, req CreateAddressReq) (*AddressInfo, error)
	ShowAddress(ctx context.Context, userID uint64) ([]*AddressInfo, error)
	UpdateAddress(ctx context.Context, req UpdateAddressReq) error
	DeleteAddress(ctx context.Context, id uint64, userID uint64) error
}

type CreateAddressReq struct {
	Name    string
	Phone   string
	Address string
	UserID  uint64
}

type AddressInfo struct {
	ID        uint64
	Name      string
	Phone     string
	Address   string
	CreatedAt uint64
}

type UpdateAddressReq struct {
	Name    string
	Phone   string
	Address string
	ID      uint64
	UserID  uint64
}
