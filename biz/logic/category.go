package logic

import (
	"context"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
)

type Category struct {
	Data *data.Data
}

func NewCategory(data *data.Data) domain.Category {
	return &Category{
		Data: data,
	}
}

func (c *Category) CreateCategory(ctx context.Context, req domain.CreateCategoryReq) error {
	return c.Data.DBClient.Category.Create().
		SetID(req.ID).
		SetCategoryName(req.CategoryName).Exec(ctx)
}

func (c *Category) ListCategory(ctx context.Context) ([]*domain.CategoryInfo, error) {
	categories, err := c.Data.DBClient.Category.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	info := make([]*domain.CategoryInfo, len(categories))
	for i, v := range categories {
		info[i] = &domain.CategoryInfo{
			ID:           v.ID,
			CategoryName: v.CategoryName,
			CreatedAt:    uint64(v.CreatedAt.Unix()),
		}
	}

	return info, nil
}
