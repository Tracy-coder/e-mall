package logic

import (
	"context"
	"errors"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/favourite"
	"github.com/jinzhu/copier"
)

type Favourite struct {
	Data *data.Data
}

func NewFavourite(data *data.Data) domain.Favourite {
	return &Favourite{
		Data: data,
	}
}

func (f *Favourite) CreateFavourite(ctx context.Context, userID uint64, productID uint64) (*domain.FavouriteInfo, error) {
	favouriteDB, err := f.Data.DBClient.Favourite.Create().
		SetUserID(userID).SetProductID(productID).Save(ctx)
	if err != nil {
		return nil, err
	}
	info := new(domain.FavouriteInfo)
	err = copier.Copy(info, favouriteDB)
	if err != nil {
		return nil, err
	}
	return info, err
}

func (f *Favourite) ShowFavourite(ctx context.Context, userID uint64, page_id int, page_size int) ([]*domain.FavouriteInfo, int, error) {
	total := f.Data.DBClient.Favourite.Query().CountX(ctx)
	FavouritesDB, err := f.Data.DBClient.Favourite.Query().Where(favourite.UserID(userID)).Offset(int(page_size) * (int(page_id - 1))).Limit(int(page_size)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	info := make([]*domain.FavouriteInfo, len(FavouritesDB))
	for i, v := range FavouritesDB {
		info[i] = &domain.FavouriteInfo{
			ID:        v.ID,
			ProductID: v.ProductID,
			UserID:    v.UserID,
			CreatedAt: uint64(v.CreatedAt.Unix()),
		}
	}

	return info, total, nil
}

func (f *Favourite) DeleteFavourite(ctx context.Context, userID uint64, productID uint64) error {
	res, err := f.Data.DBClient.Favourite.Delete().Where(favourite.UserID(userID), favourite.ProductID(productID)).Exec(ctx)
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("delete failed")
	}
	return nil
}
