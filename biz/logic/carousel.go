package logic

import (
	"context"

	"github.com/Tracy-coder/e-mall/biz/domain"
	"github.com/Tracy-coder/e-mall/data"
)

type Carousel struct {
	Data *data.Data
}

func NewCarousel(data *data.Data) domain.Carousel {
	return &Carousel{
		Data: data,
	}
}

func (c *Carousel) CreateCarousel(ctx context.Context, req domain.CreateCarouselReq) error {
	return c.Data.DBClient.Carousel.Create().
		SetImgPath(req.ImgPath).
		SetProductID(req.ProductID).Exec(ctx)
}

func (c *Carousel) ListCarousel(ctx context.Context) ([]*domain.CarouselInfo, error) {
	categories, err := c.Data.DBClient.Carousel.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	info := make([]*domain.CarouselInfo, len(categories))
	for i, v := range categories {
		info[i] = &domain.CarouselInfo{
			ID:        v.ID,
			ImgPath:   v.ImgPath,
			ProductID: v.ProductID,
			CreatedAt: uint64(v.CreatedAt.Unix()),
		}
	}

	return info, nil
}
