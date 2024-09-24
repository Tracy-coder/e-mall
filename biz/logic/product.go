package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/product"
	"github.com/jinzhu/copier"
)

type Product struct {
	Data *data.Data
}

func NewProduct(data *data.Data) domain.Product {
	return &Product{
		Data: data,
	}
}
func (p *Product) CreateProduct(ctx context.Context, req domain.CreateProductReq) (*domain.ProductInfo, error) {
	productDB, err := p.Data.DBClient.Product.Create().
		SetCategoryID(req.CategoryID).
		SetImgPath(req.ImgPath).
		SetName(req.Name).
		SetInfo(req.Info).
		SetTitle(req.Title).
		SetPrice(req.Price).
		SetDiscountPrice(req.DiscountPrice).Save(ctx)
	if err != nil {
		return nil, err
	}
	info := new(domain.ProductInfo)
	err = copier.Copy(info, &productDB)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (p *Product) ShowProduct(ctx context.Context, id uint64) (*domain.ProductInfo, error) {
	productDB, err := p.Data.DBClient.Product.Query().Where(product.ID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	view, err := p.AddView(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(productDB)
	info := new(domain.ProductInfo)
	info.View = view
	info.CreatedAt = uint64(productDB.CreatedAt.Unix())
	err = copier.Copy(info, &productDB)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (p *Product) AddView(ctx context.Context, id uint64) (uint64, error) {
	p.Data.Redis.Incr(ctx, data.ProductViewKey(id)).Result()
	view, _ := p.Data.Redis.Get(ctx, data.ProductViewKey(id)).Result()
	return strconv.ParseUint(view, 10, 64)
}

func (p *Product) ListProduct(ctx context.Context, req domain.ListProductReq) ([]*domain.ProductInfo, error) {
	productsDB, err := p.Data.DBClient.Product.Query().Where(product.CategoryID(req.CategoryID)).Limit(int(req.Limit)).Offset(int(req.Start)).All(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*domain.ProductInfo, len(productsDB))
	for i, v := range productsDB {
		res[i] = &domain.ProductInfo{
			Name:          v.Name,
			CreatedAt:     uint64(v.CreatedAt.Unix()),
			CategoryID:    v.CategoryID,
			Title:         v.Title,
			Info:          v.Info,
			ImgPath:       v.ImgPath,
			DiscountPrice: v.DiscountPrice,
			Price:         v.Price,
		}
		view, err := p.AddView(ctx, v.ID)
		if err != nil {
			continue
		}
		res[i].View = view
	}
	return res, nil
}

func (p *Product) UpdateProduct(ctx context.Context, req domain.UpdateProductReq) error {
	_, err := p.Data.DBClient.Product.Update().
		Where(product.ID(req.ID)).
		SetCategoryID(req.CategoryID).
		SetNillableImgPath(&req.ImgPath).
		SetName(req.Name).
		SetInfo(req.Info).
		SetTitle(req.Title).
		SetPrice(req.Price).
		SetDiscountPrice(req.DiscountPrice).Save(ctx)
	return err
}

func (p *Product) DeleteProduct(ctx context.Context, id uint64) error {
	return p.Data.DBClient.Product.DeleteOneID(id).Exec(ctx)
}
