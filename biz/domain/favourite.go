package domain

import "context"

type Favourite interface {
	CreateFavourite(ctx context.Context, userID uint64, productID uint64) (*FavouriteInfo, error)
	ShowFavourite(ctx context.Context, userID uint64, page_id int, page_size int) ([]*FavouriteInfo, int, error)
	DeleteFavourite(ctx context.Context, userID uint64, productID uint64) error
}

type FavouriteInfo struct {
	ID        uint64
	ProductID uint64
	UserID    uint64
	CreatedAt uint64
}
