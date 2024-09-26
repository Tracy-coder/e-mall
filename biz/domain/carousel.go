package domain

import "context"

type Carousel interface {
	ListCarousel(ctx context.Context) ([]*CarouselInfo, error)
	CreateCarousel(ctx context.Context, req CreateCarouselReq) error
}

type CreateCarouselReq struct {
	ProductID uint64
	ImgPath   string
}

type CarouselInfo struct {
	ID        uint64
	ImgPath   string
	ProductID uint64
	CreatedAt uint64
}
