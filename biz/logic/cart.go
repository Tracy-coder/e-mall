package logic

import (
	"context"
	"errors"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/cart"
	"github.com/Tracy-coder/e-mall/data/ent/product"
	"github.com/jinzhu/copier"
)

type Cart struct {
	Data *data.Data
}

func NewCart(data *data.Data) domain.Cart {
	return &Cart{
		Data: data,
	}
}

func (a *Cart) CreateCart(ctx context.Context, req domain.CreateCartReq) (*domain.CartInfo, error) {
	_, err := a.Data.DBClient.Product.Query().Where(product.ID(req.ProductID)).First(ctx)
	if err != nil {
		return nil, err
	}
	cartDB, err := a.Data.DBClient.Cart.Query().Where(cart.UserID(req.UserID), cart.ProductID(req.ProductID)).First(ctx)
	if cartDB == nil {
		cartDB, err = a.Data.DBClient.Cart.Create().
			SetCheck(false).
			SetNum(1).
			SetProductID(req.ProductID).
			SetUserID(req.UserID).
			SetMaxNum(10).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		cartDB.Num += 1
		if cartDB.Num < cartDB.MaxNum {
			a.Data.DBClient.Cart.Update().
				Where(cart.UserID(req.UserID), cart.ProductID(req.ProductID)).
				SetNum(cartDB.Num).
				Save(ctx)
		} else {
			return nil, errors.New("The maximum number is exceeded")
		}
	}
	info := new(domain.CartInfo)
	err = copier.Copy(info, &cartDB)
	if err != nil {
		return nil, err
	}
	info.CreatedAt = uint64(cartDB.CreatedAt.Unix())
	return info, nil
}

func (a *Cart) ShowCart(ctx context.Context, userID uint64) ([]*domain.CartInfo, error) {
	cartDB, err := a.Data.DBClient.Cart.Query().Where(cart.UserID(userID)).All(ctx)
	if err != nil {
		return nil, err
	}
	info := make([]*domain.CartInfo, len(cartDB))
	for i, v := range cartDB {
		info[i] = &domain.CartInfo{
			ID:        v.ID,
			UserID:    v.UserID,
			ProductID: v.ProductID,
			Num:       v.Num,
			MaxNum:    v.MaxNum,
			Check:     v.Check,
			CreatedAt: uint64(v.CreatedAt.Unix()),
		}
	}

	return info, nil
}

func (a *Cart) UpdateCart(ctx context.Context, req domain.UpdateCartReq) error {
	cartDB, err := a.Data.DBClient.Cart.Query().Where(cart.UserID(req.UserID), cart.ProductID(req.ProductID)).First(ctx)
	if req.Num > cartDB.MaxNum {
		return errors.New("The maximum number is exceeded")
	}
	res, err := a.Data.DBClient.Cart.Update().
		Where(cart.UserID(req.UserID), cart.ProductID(req.ProductID)).
		SetNum(req.Num).
		Save(ctx)
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("update failed")
	}
	return nil
}
func (a *Cart) DeleteCart(ctx context.Context, productID uint64, userID uint64) error {
	res, err := a.Data.DBClient.Cart.Delete().Where(cart.ProductID(productID), cart.UserID(userID)).Exec(ctx)
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("delete failed")
	}
	return nil
}
