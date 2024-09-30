package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
	"github.com/Tracy-coder/e-mall/data/ent/product"
	"github.com/Tracy-coder/e-mall/data/ent/productimg"
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
	err = p.AddView(ctx, id)
	if err != nil {
		return nil, err
	}
	view, err := p.GetView(ctx, id)
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

func (p *Product) AddView(ctx context.Context, id uint64) error {
	_, err := p.Data.Redis.Incr(ctx, data.ProductViewKey(id)).Result()
	if err != nil {
		return err
	}
	_, err = p.Data.Redis.ZIncrBy(ctx, data.RankKey, 1, strconv.Itoa(int(id))).Result()
	return err
}
func (p *Product) GetView(ctx context.Context, id uint64) (uint64, error) {
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
			DiscountPrice: v.DiscountPrice,
			Price:         v.Price,
		}
		view, err := p.GetView(ctx, v.ID)
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

func (p *Product) CreateProductImg(ctx context.Context, req domain.CreateProductImgReq) (*domain.ProductImg, error) {
	productImgDB, err := p.Data.DBClient.ProductImg.Create().
		SetImgPath(req.ImgPath).SetProductID(req.ProductID).Save(ctx)
	if err != nil {
		return nil, err
	}
	productImg := new(domain.ProductImg)
	err = copier.Copy(productImg, productImgDB)
	if err != nil {
		return nil, err
	}
	productImg.CreatedAt = uint64(productImgDB.CreatedAt.Unix())
	return productImg, nil
}

func (p *Product) ShowProductImg(ctx context.Context, id uint64) ([]*domain.ProductImg, error) {
	productImgDB, err := p.Data.DBClient.ProductImg.Query().Where(productimg.ProductID(id)).All(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(productImgDB)
	info := make([]*domain.ProductImg, len(productImgDB))

	err = copier.Copy(&info, &productImgDB)
	if err != nil {
		return nil, err
	}
	for i, v := range productImgDB {
		info[i].CreatedAt = uint64(v.CreatedAt.Unix())
	}
	return info, nil
}

func (p *Product) ShowProductRankings(ctx context.Context) ([]*domain.ProductInfo, error) {
	pros, err := p.Data.Redis.ZRevRange(ctx, data.RankKey, 0, 9).Result()
	if err != nil {
		return nil, err
	}
	ids := make([]uint64, len(pros))
	for i, v := range pros {
		ids[i], _ = strconv.ParseUint(v, 10, 64)
	}
	productsDB, err := p.Data.DBClient.Product.Query().Where(product.IDIn(ids...)).All(ctx)
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
			DiscountPrice: v.DiscountPrice,
			Price:         v.Price,
		}
		view, err := p.GetView(ctx, v.ID)
		if err != nil {
			continue
		}
		res[i].View = view
	}
	return res, nil
}
