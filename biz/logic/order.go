package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/address"
	"github.com/Tracy-coder/e-mall/data/ent/favourite"
	"github.com/Tracy-coder/e-mall/data/ent/order"
	"github.com/Tracy-coder/e-mall/data/ent/product"
	productimg "github.com/Tracy-coder/e-mall/data/ent/productimg"
	"github.com/redis/go-redis/v9"
	"golang.org/x/exp/rand"
)

type Order struct {
	Data *data.Data
}

func NewOrder(data *data.Data) domain.Order {
	return &Order{
		Data: data,
	}
}

func (a *Order) CreateOrder(ctx context.Context, req domain.CreateOrderReq) error {
	addr, err := a.Data.DBClient.Address.Query().Where(address.ID(req.AddressID)).First(ctx)
	if err != nil {
		return err
	}
	productDB, err := a.Data.DBClient.Product.Query().Where(product.ID(req.ProductID)).First(ctx)
	if err != nil {
		return err
	}
	number := fmt.Sprintf("%09v", rand.New(rand.NewSource(uint64(time.Now().UnixNano()))).Int31n(1000000000))
	productNum := strconv.Itoa(int(req.ProductID))
	userNum := strconv.Itoa(int(req.UserID))
	number = number + productNum + userNum
	orderNum, err := strconv.ParseUint(number, 10, 64)
	price := int64(0)
	if productDB.DiscountPrice != 0 {
		price = int64(req.Num) * int64(productDB.DiscountPrice)
	} else {
		price = int64(req.Num) * int64(productDB.Price)
	}
	_, err = a.Data.DBClient.Order.Create().
		SetOrderNum(orderNum).
		SetAddress(addr.Address).
		SetAddressName(addr.Name).
		SetAddressPhone(addr.Phone).
		SetProductID(req.ProductID).
		SetNum(req.Num).
		SetUserID(req.UserID).
		SetType(1).
		SetPrice(price).
		Save(ctx)
	if err != nil {
		return err
	}
	data := redis.Z{Score: float64(time.Now().Unix()) + 15*time.Minute.Seconds(), Member: orderNum}
	return a.Data.Redis.ZAdd(ctx, "order", data).Err()
}

func (a *Order) ShowOrder(ctx context.Context, orderNum uint64) (*domain.OrderInfo, error) {
	orderDB, err := a.Data.DBClient.Order.Query().Where(order.OrderNum(orderNum)).First(ctx)
	if err != nil {
		return nil, err
	}
	productDB, err := a.Data.DBClient.Product.Query().Where(product.ID(orderDB.ProductID)).First(ctx)
	if err != nil {
		return nil, err
	}
	productImgDB, err := a.Data.DBClient.ProductImg.Query().Where(productimg.ProductID(orderDB.ProductID)).First(ctx)
	if err != nil {
		return nil, err
	}
	info := domain.OrderInfo{
		ID:           orderDB.ID,
		OrderNum:     orderDB.Num,
		CreatedAt:    uint64(orderDB.CreatedAt.Unix()),
		UpdatedAt:    uint64(orderDB.UpdatedAt.Unix()),
		UserID:       orderDB.UserID,
		ProductID:    orderDB.ProductID,
		Num:          orderDB.Num,
		Address:      orderDB.Address,
		AddressName:  orderDB.AddressName,
		AddressPhone: orderDB.AddressPhone,
		Name:         productDB.Name,
		ImgPath:      productImgDB.ImgPath,
		Price:        orderDB.Price,
	}

	return &info, nil
}

func (a *Order) ListOrder(ctx context.Context, req domain.ListOrderReq) ([]*domain.OrderInfo, int, error) {
	total := a.Data.DBClient.Order.Query().Where(order.UserID(req.UserID)).CountX(ctx)
	ordersDB, err := a.Data.DBClient.Order.Query().Where(order.UserID(req.UserID)).Offset(int(req.PageSize) * (int(req.PageID - 1))).Limit(int(req.PageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	info := make([]*domain.OrderInfo, len(ordersDB))
	for i, orderDB := range ordersDB {
		productDB, err := a.Data.DBClient.Product.Query().Where(product.ID(orderDB.ProductID)).First(ctx)
		if err != nil {
			return nil, 0, err
		}
		productImgDB, err := a.Data.DBClient.ProductImg.Query().Where(productimg.ProductID(orderDB.ProductID)).First(ctx)
		if err != nil {
			return nil, 0, err
		}
		info[i] = &domain.OrderInfo{
			ID:           orderDB.ID,
			OrderNum:     orderDB.Num,
			CreatedAt:    uint64(orderDB.CreatedAt.Unix()),
			UpdatedAt:    uint64(orderDB.UpdatedAt.Unix()),
			UserID:       orderDB.UserID,
			ProductID:    orderDB.ProductID,
			Num:          orderDB.Num,
			Address:      orderDB.Address,
			AddressName:  orderDB.AddressName,
			AddressPhone: orderDB.AddressPhone,
			Name:         productDB.Name,
			ImgPath:      productImgDB.ImgPath,
			Price:        orderDB.Price,
		}
	}
	return info, total, nil
}

func (a *Order) ShowCount(ctx context.Context, userID uint64) (*domain.ShowCountResp, error) {
	favouriteCount, err := a.Data.DBClient.Favourite.Query().Where(favourite.UserID(userID)).Count(ctx)
	if err != nil {
		return nil, err
	}
	notPayCount, err := a.Data.DBClient.Order.Query().Where(order.UserID(userID), order.Type(1)).Count(ctx)
	if err != nil {
		return nil, err
	}
	payCount, err := a.Data.DBClient.Order.Query().Where(order.UserID(userID), order.Type(2)).Count(ctx)
	if err != nil {
		return nil, err
	}
	return &domain.ShowCountResp{
		FavoriteTotal: int32(favouriteCount),
		NotPayTotal:   int32(notPayCount),
		PayTotal:      int32(payCount),
	}, nil
}
