package logic

import (
	"context"
	"errors"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/address"
	"github.com/jinzhu/copier"
)

type Address struct {
	Data *data.Data
}

func NewAddress(data *data.Data) domain.Address {
	return &Address{
		Data: data,
	}
}

func (a *Address) CreateAddress(ctx context.Context, req domain.CreateAddressReq) (*domain.AddressInfo, error) {
	addressDB, err := a.Data.DBClient.Address.Create().
		SetName(req.Name).
		SetPhone(req.Phone).
		SetAddress(req.Address).
		SetUserID(req.UserID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	info := new(domain.AddressInfo)
	err = copier.Copy(info, &addressDB)
	if err != nil {
		return nil, err
	}
	info.CreatedAt = uint64(addressDB.CreatedAt.Unix())
	return info, nil
}

func (a *Address) ShowAddress(ctx context.Context, userID uint64) ([]*domain.AddressInfo, error) {
	addressDB, err := a.Data.DBClient.Address.Query().Where(address.UserID(userID)).All(ctx)
	if err != nil {
		return nil, err
	}
	info := make([]*domain.AddressInfo, len(addressDB))
	for i, v := range addressDB {
		info[i] = &domain.AddressInfo{
			ID:        v.ID,
			Name:      v.Name,
			Phone:     v.Phone,
			Address:   v.Address,
			CreatedAt: uint64(v.CreatedAt.Unix()),
		}
	}

	return info, nil
}

func (a *Address) UpdateAddress(ctx context.Context, req domain.UpdateAddressReq) error {
	res, err := a.Data.DBClient.Address.Update().
		Where(address.ID(req.ID), address.UserID(req.UserID)).
		SetName(req.Name).
		SetAddress(req.Address).
		SetPhone(req.Phone).Save(ctx)
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("update failed")
	}
	return nil
}
func (a *Address) DeleteAddress(ctx context.Context, id uint64, userID uint64) error {
	res, err := a.Data.DBClient.Address.Delete().Where(address.ID(id), address.UserID(userID)).Exec(ctx)
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("delete failed")
	}
	return nil
}
