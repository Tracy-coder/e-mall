package domain

import (
	"context"
	"net/url"
)

type Payment interface {
	Init(ctx context.Context, orderNum uint64, userID uint64) (*url.URL, error)
	Notify(ctx context.Context, v url.Values) error
}
