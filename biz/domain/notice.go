package domain

import "context"

type Notice interface {
	ShowNotice(ctx context.Context, page_id int, page_size int) ([]*NoticeInfo, int, error)
	CreateNotice(ctx context.Context, text string) (*NoticeInfo, error)
}

type NoticeInfo struct {
	ID        uint64
	Text      string
	CreatedAt uint64
}
